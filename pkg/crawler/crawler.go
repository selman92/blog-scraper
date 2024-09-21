package crawler

import (
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/selman92/blog-scraper/internal/models"
	"github.com/selman92/blog-scraper/internal/storage"
)

type Crawler struct {
	storage storage.Storage
	mutex   sync.Mutex
}

func NewCrawler(storage storage.Storage) *Crawler {
	return &Crawler{
		storage: storage,
	}
}

func (c *Crawler) Start() {
	for {
		sites, err := c.storage.GetBlogSites()
		if err != nil {
			log.Printf("Error fetching blog sites: %v", err)
			time.Sleep(time.Minute)
			continue
		}

		var wg sync.WaitGroup
		for _, site := range sites {
			wg.Add(1)
			go func(site models.BlogSite) {
				defer wg.Done()
				c.crawlSite(site)
			}(site)
		}
		wg.Wait()

		time.Sleep(time.Second * 5)
	}
}

func (c *Crawler) crawlSite(site models.BlogSite) {
	uri, err := url.Parse(site.URL)

	if err != nil {
		fmt.Println("Error parsing site URL: " + site.URL)
		return
	}

	collector := colly.NewCollector(
		colly.AllowedDomains(uri.Host),
		colly.MaxDepth(2),
	)

	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	collector.OnHTML("body", func(e *colly.HTMLElement) {

		titleElements := e.DOM.Find(site.TitleSelector)
		if titleElements.Length() != 1 {
			log.Printf("Skipping page %s: Title selector found %d times (expected 1)", e.Request.URL, titleElements.Length())
			return
		}

		title := e.ChildText(site.TitleSelector)
		timeStr := e.ChildText(site.TimeSelector)
		postTime, err := time.Parse(site.TimeLayout, timeStr)
		if err != nil {
			log.Printf("Error parsing time for %s: %v", e.Request.URL, err)
			return
		}

		post := models.BlogPost{
			BlogID:   site.ID,
			URL:      e.Request.URL.String(),
			Title:    title,
			PostTime: postTime,
		}

		c.mutex.Lock()
		err = c.storage.AddBlogPost(post)
		c.mutex.Unlock()

		if err != nil {
			log.Printf("Error saving blog post: %v", err)
		}
	})

	err = collector.Visit(site.URL)
	if err != nil {
		log.Printf("Error visiting %s: %v", site.URL, err)
	}
}
