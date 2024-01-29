package repository

import (
	"errors"
	"github.com/hudayberdipolat/go-url-shortener/internal/models"
	"gorm.io/gorm"
	"log"
)

type globalRedirectUrlRepositoryImp struct {
	db *gorm.DB
}

func NewGlobalRedirectUrlRepository(db *gorm.DB) GlobalRedirectUrlRepository {
	return globalRedirectUrlRepositoryImp{
		db: db,
	}
}

func (g globalRedirectUrlRepositoryImp) GetLongURL(shortUrl string) (string, error) {
	var url models.Url
	if err := g.db.Where("short_url=?", shortUrl).First(&url).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("Nädogry url!!! url not found")
		}
		return "", err
	}
	longURL := url.LongUrl
	log.Println(longURL)
	return longURL, nil
}
