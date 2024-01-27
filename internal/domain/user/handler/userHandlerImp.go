package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/go-url-shortener/internal/domain/user/dto"
	"github.com/hudayberdipolat/go-url-shortener/internal/domain/user/service"
	"github.com/hudayberdipolat/go-url-shortener/internal/utils/response"
	"github.com/hudayberdipolat/go-url-shortener/internal/utils/validate"
	"net/http"
)

type userHandlerImp struct {
	userService service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return userHandlerImp{
		userService: service,
	}
}

func (u userHandlerImp) GetUser(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(int)

	userResponse, err := u.userService.GetUserData(userID)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "bad request", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "user data", userResponse)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (u userHandlerImp) UpdateUserData(ctx *fiber.Ctx) error {
	var updateUser dto.UpdateUserData
	userID := ctx.Locals("user_id").(int)

	if err := ctx.BodyParser(&updateUser); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// validate
	if errValidate := validate.ValidateStruct(updateUser); errValidate != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", errValidate.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	if err := u.userService.UpdateUser(userID, updateUser); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "user data can't updated", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "userData updated successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (u userHandlerImp) UpdateUserPassword(ctx *fiber.Ctx) error {
	var updateUserPassword dto.ChangeUserPassword
	userID := ctx.Locals("user_id").(int)
	if err := ctx.BodyParser(&updateUserPassword); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// validate
	if errValidate := validate.ValidateStruct(updateUserPassword); errValidate != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", errValidate.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	if err := u.userService.ChangePassword(userID, updateUserPassword); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "password can't updated!!!", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, " password changed successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)

}
