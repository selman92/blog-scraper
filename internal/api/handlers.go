package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/selman92/blog-scraper/internal/models"
	"github.com/selman92/blog-scraper/internal/storage"
)

func SetupRoutes(r *gin.Engine, store storage.Storage) {
	r.POST("/blog-sites", addBlogSite(store))
	r.DELETE("/blog-sites/:id", removeBlogSite(store))
	r.GET("/blog-sites", listBlogSites(store))
}

func addBlogSite(store storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		var site models.BlogSite
		if err := c.ShouldBindJSON(&site); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := store.AddBlogSite(site); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add blog site"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Blog site added successfully"})
	}
}

func removeBlogSite(store storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var idInt int
		if _, err := fmt.Sscanf(id, "%d", &idInt); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		if err := store.RemoveBlogSite(idInt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove blog site"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Blog site removed successfully"})
	}
}

func listBlogSites(store storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		sites, err := store.ListBlogSites()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list blog sites"})
			return
		}

		c.JSON(http.StatusOK, sites)
	}
}
