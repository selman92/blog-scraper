package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/selman92/blog-scraper/internal/models"
)

type SQLiteStorage struct {
	db *sql.DB
}

func NewSQLiteStorage(dbPath string) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Create table if not exists
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS blog_sites (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			url TEXT NOT NULL,
			title_selector TEXT NOT NULL,
			time_selector TEXT NOT NULL
		)
	`)
	if err != nil {
		return nil, err
	}

	return &SQLiteStorage{db: db}, nil
}

func (s *SQLiteStorage) AddBlogSite(site models.BlogSite) error {
	_, err := s.db.Exec("INSERT INTO blog_sites (url, title_selector, time_selector) VALUES (?, ?, ?)",
		site.URL, site.TitleSelector, site.TimeSelector)
	return err
}

func (s *SQLiteStorage) RemoveBlogSite(id int) error {
	_, err := s.db.Exec("DELETE FROM blog_sites WHERE id = ?", id)
	return err
}

func (s *SQLiteStorage) ListBlogSites() ([]models.BlogSite, error) {
	rows, err := s.db.Query("SELECT id, url, title_selector, time_selector FROM blog_sites")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sites []models.BlogSite
	for rows.Next() {
		var site models.BlogSite
		if err := rows.Scan(&site.ID, &site.URL, &site.TitleSelector, &site.TimeSelector); err != nil {
			return nil, err
		}
		sites = append(sites, site)
	}

	return sites, nil
}
