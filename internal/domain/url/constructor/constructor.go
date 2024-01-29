package constructor

import (
	"github.com/hudayberdipolat/go-url-shortener/internal/domain/url/handler"
	"github.com/hudayberdipolat/go-url-shortener/internal/domain/url/repository"
	"github.com/hudayberdipolat/go-url-shortener/internal/domain/url/service"
	"gorm.io/gorm"
)

var urlRepository repository.UrlRepository
var urlService service.UrlService
var URLHandler handler.UrlHandler

func URLRequirementsCreator(db *gorm.DB) {
	urlRepository = repository.NewUrlRepository(db)
	urlService = service.NewUrlService(urlRepository)
	URLHandler = handler.NewUrlHandler(urlService)
}
