package storage

import (
	"database/sql"
	"time"

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
			time_selector TEXT NOT NULL,
			time_layout TEXT NOT NULL
		)
	`)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS blog_posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		blog_id INTEGER,
		url TEXT NOT NULL,
		title TEXT NOT NULL,
		post_time DATETIME,
		created_at DATETIME,
		FOREIGN KEY(blog_id) REFERENCES blog_sites(id)
	)
`)
	if err != nil {
		return nil, err
	}

	return &SQLiteStorage{db: db}, nil
}

func (s *SQLiteStorage) AddBlogSite(site models.BlogSite) error {
	_, err := s.db.Exec("INSERT INTO blog_sites (url, title_selector, time_selector, time_layout) VALUES (?, ?, ?, ?)",
		site.URL, site.TitleSelector, site.TimeSelector, site.TimeLayout)
	return err
}

func (s *SQLiteStorage) RemoveBlogSite(id int) error {
	_, err := s.db.Exec("DELETE FROM blog_sites WHERE id = ?", id)
	return err
}

func (s *SQLiteStorage) ListBlogSites() ([]models.BlogSite, error) {
	rows, err := s.db.Query("SELECT id, url, title_selector, time_selector, time_layout FROM blog_sites")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sites []models.BlogSite
	for rows.Next() {
		var site models.BlogSite
		if err := rows.Scan(&site.ID, &site.URL, &site.TitleSelector, &site.TimeSelector, &site.TimeLayout); err != nil {
			return nil, err
		}
		sites = append(sites, site)
	}

	return sites, nil
}

func (s *SQLiteStorage) AddBlogPost(post models.BlogPost) error {
	_, err := s.db.Exec(`
		INSERT INTO blog_posts (blog_id, url, title, post_time, created_at)
		VALUES (?, ?, ?, ?, ?)
	`, post.BlogID, post.URL, post.Title, post.PostTime, time.Now())
	return err
}

func (s *SQLiteStorage) GetBlogSites() ([]models.BlogSite, error) {
	rows, err := s.db.Query("SELECT id, url, title_selector, time_selector, time_layout FROM blog_sites")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sites []models.BlogSite
	for rows.Next() {
		var site models.BlogSite
		if err := rows.Scan(&site.ID, &site.URL, &site.TitleSelector, &site.TimeSelector, &site.TimeLayout); err != nil {
			return nil, err
		}
		sites = append(sites, site)
	}
	return sites, nil
}
