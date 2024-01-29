package service

import "github.com/hudayberdipolat/go-url-shortener/internal/domain/globalRedirectUrl/repository"

type globalRedirectUrlServiceImp struct {
	repo repository.GlobalRedirectUrlRepository
}

func NewGlobalRedirectUrlService(repo repository.GlobalRedirectUrlRepository) GlobalRedirectUrlService {
	return globalRedirectUrlServiceImp{
		repo: repo,
	}
}

func (g globalRedirectUrlServiceImp) GetLongURL(shortUrl string) (string, error) {
	longURL, err := g.repo.GetLongURL(shortUrl)
	if err != nil {
		return "", err
	}
	return longURL, nil
}
