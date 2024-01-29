package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-url-shortener/internal/domain/globalRedirectUrl/service"
	"github.com/hudayberdipolat/go-url-shortener/internal/utils/response"
	"net/http"
)

type globalRedirectURLHandlerImp struct {
	service service.GlobalRedirectUrlService
}

func NewGlobalRedirectURLHandler(service service.GlobalRedirectUrlService) GlobalRedirectURLHandler {
	return globalRedirectURLHandlerImp{
		service: service,
	}
}

func (g globalRedirectURLHandlerImp) RedirectLongURL(ctx *fiber.Ctx) error {
	getURL := ctx.Params("shortURL")

	if getURL == "" {
		errResponse := response.Error(http.StatusBadRequest, "error short URL", "", nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	longURl, err := g.service.GetLongURL(getURL)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "short url not found", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	return ctx.Redirect(longURl, http.StatusPermanentRedirect)
}
