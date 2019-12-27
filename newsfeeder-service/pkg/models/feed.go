package models

import (
	"fmt"
	"time"
)

// Feed models a news feed.
type Feed struct {
	ID       uint64    `json:"feed_id" db:"feed_id"`
	Created  time.Time `json:"created_at" db:"created_at"`
	Updated  time.Time `json:"modified_at" db:"modified_at"`
	Category string    `json:"category" db:"category"`
	Provider string    `json:"provider" db:"provider"`
	URL      string    `json:"url" db:"url"`
}

// Article models a news feed's article.
type Article struct {
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	Link              string    `json:"link"`
	Published         time.Time `json:"published"`
	GUID              string    `json:"guid"`
	ThumbnailImageURL string    `json:"thumbnail_image_url"`
	HTMLContent       string    `json:"html_content,omitempty"`
}

// String returns string representation of an Article model.
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
