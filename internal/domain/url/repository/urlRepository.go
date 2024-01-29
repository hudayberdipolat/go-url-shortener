package repository

import "github.com/hudayberdipolat/go-url-shortener/internal/models"

type UrlRepository interface {
	GetAllUrl(userID int) ([]models.Url, error)
	GetUrlByID(userID, urlID int) (*models.Url, error)
	Create(url models.Url) error
	Delete(userID, urlID int) error
}
