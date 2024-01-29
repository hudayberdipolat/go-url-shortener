package constructor

import (
	"github.com/hudayberdipolat/go-url-shortener/internal/domain/globalRedirectUrl/handler"
	"github.com/hudayberdipolat/go-url-shortener/internal/domain/globalRedirectUrl/repository"
	"github.com/hudayberdipolat/go-url-shortener/internal/domain/globalRedirectUrl/service"
	"gorm.io/gorm"
)

var globalURLRepository repository.GlobalRedirectUrlRepository
var globalURLService service.GlobalRedirectUrlService
var GlobalURLHandler handler.GlobalRedirectURLHandler

func GlobalURLRequirementsCreator(db *gorm.DB) {
	globalURLRepository = repository.NewGlobalRedirectUrlRepository(db)
	globalURLService = service.NewGlobalRedirectUrlService(globalURLRepository)
	GlobalURLHandler = handler.NewGlobalRedirectURLHandler(globalURLService)
}
