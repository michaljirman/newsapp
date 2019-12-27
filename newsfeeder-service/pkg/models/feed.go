package models

import (
	"fmt"
	"time"
)

type Feed struct {
	ID       uint64    `json:"feed_id" db:"feed_id"`
	Created  time.Time `json:"created_at" db:"created_at"`
	Updated  time.Time `json:"updated_at" db:"updated_at"`
	Category string    `json:"category" db:"category"`
	Provider string    `json:"provider" db:"provider"`
	URL      string    `json:"url" db:"url"`
}

type Article struct {
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	Link              string    `json:"link"`
	Published         time.Time `json:"published"`
	GUID              string    `json:"guid"`
	ThumbnailImageURL string    `json:"thumbnail_image_url"`
	HTMLContent       string    `json:"html_content"`
}

func (a Article) String() string {
	return fmt.Sprintf(`
	Article
		Title: %s
		Description: %s
		Link: %s
		Published: %+v
		GUID: %s
		ThumbnailImageURL: %s
		HTMLContent length: %d
	`, a.Title, a.Description, a.Link, a.Published, a.GUID, a.ThumbnailImageURL, len(a.HTMLContent))
}
