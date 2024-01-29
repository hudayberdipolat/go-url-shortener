package models

import "time"

type Url struct {
	ID           int       `json:"id"`
	UrlName      string    `json:"url_name"`
	ShortUrl     string    `json:"short_url"`
	LongUrl      string    `json:"long_url"`
	VisitedCount int       `json:"visited_count"`
	UserID       int       `json:"user_id"`
	CreatedAt    time.Time `json:"created_at"`
}

func (*Url) TableName() string {
	return "urls"
}
