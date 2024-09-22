# Blog Scraper API

This Go application scrapes blog post titles from a list of blogs stored in an SQLite database and provides an API to manage blog sites and retrieve blog posts.

## Features

- Add a new blog site with specific CSS selectors for title and publication time.
- Remove a blog site by its ID.
- List all blog sites.
- Retrieve the latest blog posts.

## API Endpoints

| Method | Endpoint              | Description                        |
|--------|-----------------------|------------------------------------|
| POST   | `/blog-sites`          | Add a new blog site                |
| DELETE | `/blog-sites/:id`      | Remove a blog site by its ID       |
| GET    | `/blog-sites`          | List all blog sites                |
| GET    | `/blog-posts`          | Retrieve latest blog posts         |

## Sample Usage

### Add a Blog Site

```bash
curl -X POST http://localhost:8080/blog-sites \
-H "Content-Type: application/json" \
-d '{
  "url": "https://example-blog.com",
  "title_selector": ".post-title",
  "time_selector": ".post-date",
  "time_layout": "2006-01-02T15:04:05Z07:00"
}'
