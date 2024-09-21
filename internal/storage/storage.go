package storage

import "github.com/selman92/blog-scraper/internal/models"

type Storage interface {
	AddBlogSite(site models.BlogSite) error
	RemoveBlogSite(id int) error
	AddBlogPost(post models.BlogPost) error
	GetBlogSites() ([]models.BlogSite, error)
	GetBlogPosts() ([]models.BlogPost, error)
}
