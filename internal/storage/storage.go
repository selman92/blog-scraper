package storage

import "github.com/selman92/blog-scraper/internal/models"

type Storage interface {
	AddBlogSite(site models.BlogSite) error
	RemoveBlogSite(id int) error
	ListBlogSites() ([]models.BlogSite, error)
	AddBlogPost(post models.BlogPost) error
	GetBlogSites() ([]models.BlogSite, error)
}
