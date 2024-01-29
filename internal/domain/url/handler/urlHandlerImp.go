package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-url-shortener/internal/domain/url/dto"
	"github.com/hudayberdipolat/go-url-shortener/internal/domain/url/service"
	"github.com/hudayberdipolat/go-url-shortener/internal/utils/response"
	"github.com/hudayberdipolat/go-url-shortener/internal/utils/validate"
	"net/http"
	"strconv"
)

type urlHandlerImp struct {
	urlService service.UrlService
}

func NewUrlHandler(service service.UrlService) UrlHandler {
	return urlHandlerImp{
		urlService: service,
	}
}

func (u urlHandlerImp) GetAll(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(int)
	urls, err := u.urlService.GetAllUrl(userID)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "urls not found", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get all urls of user", urls)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (u urlHandlerImp) GetOne(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(int)
	urlID, _ := strconv.Atoi(ctx.Params("urlID"))
	urls, err := u.urlService.GetUrlByID(userID, urlID)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "url not found", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "get one url of user", urls)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (u urlHandlerImp) Create(ctx *fiber.Ctx) error {
	var createRequest dto.CreateUrlRequest
	userID := ctx.Locals("user_id").(int)
	if err := ctx.BodyParser(&createRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// validate
	if errValidate := validate.ValidateStruct(createRequest); errValidate != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", errValidate.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	if err := u.urlService.CreateUrl(userID, createRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "short URL can't created", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "short URL created successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (u urlHandlerImp) Delete(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(int)
	urlID, _ := strconv.Atoi(ctx.Params("urlID"))
	if err := u.urlService.DeleteUrl(userID, urlID); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "short URL can't deleted", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "short URL deleted successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
