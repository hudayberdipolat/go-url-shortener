package dto

type CreateUrlRequest struct {
	UrlName string `json:"url_name" validate:"required"`
	LongUrl string `json:"long_url" validate:"required"`
}
