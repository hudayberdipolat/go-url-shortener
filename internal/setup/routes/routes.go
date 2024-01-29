package routes

import (
	"github.com/gofiber/fiber/v2"
	authConstructor "github.com/hudayberdipolat/go-url-shortener/internal/domain/auth/constructor"
	GlobalURLConstructor "github.com/hudayberdipolat/go-url-shortener/internal/domain/globalRedirectUrl/constructor"
	urlConstructor "github.com/hudayberdipolat/go-url-shortener/internal/domain/url/constructor"
	userConstructor "github.com/hudayberdipolat/go-url-shortener/internal/domain/user/constructor"
	"github.com/hudayberdipolat/go-url-shortener/internal/middleware"
)

func Routes(app *fiber.App) {

	// global URL Redirect route
	app.Get("/:shortURL", GlobalURLConstructor.GlobalURLHandler.RedirectLongURL)

	// api client
	apiClient := app.Group("/api")
	//  auth routes
	authRoute := apiClient.Group("/auth")
	authRoute.Post("/register", authConstructor.AuthHandler.Register)
	authRoute.Post("/login", authConstructor.AuthHandler.Login)

	// userData routes
	userRoute := apiClient.Group("/user")
	userRoute.Use(middleware.AuthMiddleware)
	userRoute.Get("/", userConstructor.UserHandler.GetUser)
	userRoute.Put("/update-data", userConstructor.UserHandler.UpdateUserData)
	userRoute.Put("/change-password", userConstructor.UserHandler.UpdateUserPassword)

	// url routes
	urlRoute := userRoute.Group("/urls")
	urlRoute.Get("/", urlConstructor.URLHandler.GetAll)
	urlRoute.Get("/:urlID", urlConstructor.URLHandler.GetOne)
	urlRoute.Post("/create", urlConstructor.URLHandler.Create)
	urlRoute.Delete("/:urlID/delete", urlConstructor.URLHandler.Delete)

}
