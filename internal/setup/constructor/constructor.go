package constructor

import (
	"github.com/hudayberdipolat/go-url-shortener/internal/app"
	authConstructor "github.com/hudayberdipolat/go-url-shortener/internal/domain/auth/constructor"
	urlConstructor "github.com/hudayberdipolat/go-url-shortener/internal/domain/url/constructor"
	userConstructor "github.com/hudayberdipolat/go-url-shortener/internal/domain/user/constructor"
)

func Build(dependencies *app.Dependencies) {
	authConstructor.AuthRequirementsCreator(dependencies.DB)
	userConstructor.UserRequirementsCreator(dependencies.DB)
	urlConstructor.URLRequirementsCreator(dependencies.DB)
}
