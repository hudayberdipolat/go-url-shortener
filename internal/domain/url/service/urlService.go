package service

import (
	"github.com/hudayberdipolat/go-url-shortener/internal/domain/url/dto"
)

type UrlService interface {
	GetAllUrl(userID int) ([]dto.UrlResponse, error)
	GetUrlByID(userID, urlID int) (*dto.UrlResponse, error)
	CreateUrl(userID int, request dto.CreateUrlRequest) error
	DeleteUrl(userID, urlID int) error
}
