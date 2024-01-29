package service

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-url-shortener/internal/domain/url/dto"
	"github.com/hudayberdipolat/go-url-shortener/internal/domain/url/repository"
	"github.com/hudayberdipolat/go-url-shortener/internal/models"
	"github.com/hudayberdipolat/go-url-shortener/internal/utils"
	"time"
)

type urlServiceImp struct {
	urlRepo repository.UrlRepository
}

func NewUrlService(repo repository.UrlRepository) UrlService {
	return urlServiceImp{
		urlRepo: repo,
	}
}

func (u urlServiceImp) GetAllUrl(userID int) ([]dto.UrlResponse, error) {
	urls, err := u.urlRepo.GetAllUrl(userID)
	if err != nil {
		return nil, err
	}
	var urlResponses []dto.UrlResponse
	for _, url := range urls {
		urlResponse := dto.NewUrlResponse(url)
		urlResponses = append(urlResponses, urlResponse)
	}
	return urlResponses, nil
}

func (u urlServiceImp) GetUrlByID(userID, urlID int) (*dto.UrlResponse, error) {
	url, err := u.urlRepo.GetUrlByID(userID, urlID)
	if err != nil {
		return nil, err
	}
	urlResponse := dto.NewUrlResponse(*url)
	return &urlResponse, err
}

func (u urlServiceImp) CreateUrl(ctx *fiber.Ctx, userID int, request dto.CreateUrlRequest) error {
	// create short url
	baseUrl := ctx.BaseURL()
	shortUrl := utils.CreateShortURL(baseUrl, request.LongUrl)
	createURL := models.Url{
		UrlName:      request.UrlName,
		ShortUrl:     shortUrl,
		LongUrl:      request.LongUrl,
		VisitedCount: 0,
		UserID:       userID,
		CreatedAt:    time.Now(),
	}
	if err := u.urlRepo.Create(createURL); err != nil {
		return err
	}
	return nil
}

func (u urlServiceImp) DeleteUrl(userID, urlID int) error {
	getURL, err := u.GetUrlByID(userID, urlID)
	if err != nil {
		return errors.New("Ýalňyşlyk döredi!!!")
	}
	if errDelete := u.urlRepo.Delete(userID, getURL.ID); errDelete != nil {
		return errDelete
	}
	return nil
}
