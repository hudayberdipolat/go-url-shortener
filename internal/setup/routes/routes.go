package routes

import (
	"github.com/gofiber/fiber/v2"
	authConstructor "github.com/hudayberdipolat/go-url-shortener/internal/domain/auth/constructor"
	userConstructor "github.com/hudayberdipolat/go-url-shortener/internal/domain/user/constructor"
	"github.com/hudayberdipolat/go-url-shortener/internal/middleware"
)

func Routes(app *fiber.App) {
	apiClient := app.Group("/api")
	// user auth
	authRoute := apiClient.Group("/auth")
	authRoute.Post("/register", authConstructor.AuthHandler.Register)
	authRoute.Post("/login", authConstructor.AuthHandler.Login)

	// user data
	userRoute := apiClient.Group("/user")
	userRoute.Use(middleware.AuthMiddleware)
	userRoute.Get("/", userConstructor.UserHandler.GetUser)
	userRoute.Put("/update-data", userConstructor.UserHandler.UpdateUserData)
	userRoute.Put("/change-password", userConstructor.UserHandler.UpdateUserPassword)
}
