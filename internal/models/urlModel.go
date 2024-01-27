package models

import "time"

type Url struct {
	ID           int
	ShortUrl     string
	LongUrl      string
	VisitedCount int
	UserID       int
	CreatedAt    time.Time
}

func (*Url) TableName() string {
	return "urls"
}
