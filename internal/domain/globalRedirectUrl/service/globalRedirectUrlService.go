package service

type GlobalRedirectUrlService interface {
	GetLongURL(shortUrl string) (string, error)
}
