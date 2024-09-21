package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/selman92/blog-scraper/internal/api"
	"github.com/selman92/blog-scraper/internal/storage"
	"github.com/selman92/blog-scraper/pkg/crawler"
)

func main() {
	store, err := storage.NewSQLiteStorage("blog_sites.db")
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}

	crawler := crawler.NewCrawler(store)
	go crawler.Start()

	r := gin.Default()
	api.SetupRoutes(r, store)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
