package dto

import "github.com/hudayberdipolat/go-url-shortener/internal/models"

type UrlResponse struct {
	ID           int    `json:"id"`
	UrlName      string `json:"url_name"`
	ShortUrl     string `json:"short_url"`
	LongUrl      string `json:"long_url"`
	VisitedCount int    `json:"visited_count"`
	CreatedAt    string `json:"created_at"`
}

func NewUrlResponse(url models.Url) UrlResponse {
	return UrlResponse{
		ID:           url.ID,
		UrlName:      url.UrlName,
		ShortUrl:     url.ShortUrl,
		LongUrl:      url.LongUrl,
		VisitedCount: url.VisitedCount,
		CreatedAt:    url.CreatedAt.Format("01-02-2006"),
	}
}
