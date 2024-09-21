package models

import "time"

type BlogPost struct {
	ID        int       `json:"id"`
	BlogID    int       `json:"blog_id"`
	URL       string    `json:"url"`
	Title     string    `json:"title"`
	PostTime  time.Time `json:"post_time"`
	CreatedAt time.Time `json:"created_at"`
}
