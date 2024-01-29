package handler

import "github.com/gofiber/fiber/v2"

type GlobalRedirectURLHandler interface {
	RedirectLongURL(ctx *fiber.Ctx) error
}
