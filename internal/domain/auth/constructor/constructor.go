package constructor

import (
	"github.com/hudayberdipolat/go-url-shortener/internal/domain/auth/handler"
	"github.com/hudayberdipolat/go-url-shortener/internal/domain/auth/repository"
	"github.com/hudayberdipolat/go-url-shortener/internal/domain/auth/service"
	"gorm.io/gorm"
)

var authRepo repository.AuthRepository
var authService service.AuthService
var AuthHandler handler.AuthHandler

func AuthRequirementsCreator(db *gorm.DB) {
	authRepo = repository.NewAuthRepository(db)
	authService = service.NewAuthService(authRepo)
	AuthHandler = handler.NewAuthHandler(authService)
}
