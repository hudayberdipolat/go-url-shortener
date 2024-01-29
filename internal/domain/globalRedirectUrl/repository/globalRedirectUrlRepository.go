package repository

type GlobalRedirectUrlRepository interface {
	GetLongURL(shortUrl string) (string, error)
}
