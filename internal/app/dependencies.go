package app

import (
	"github.com/hudayberdipolat/go-url-shortener/pkg/config"
	"github.com/hudayberdipolat/go-url-shortener/pkg/database/dbConfig"
	CustomHttp "github.com/hudayberdipolat/go-url-shortener/pkg/http"
	"gorm.io/gorm"
	"net/http"
)

type Dependencies struct {
	DB         *gorm.DB
	Config     *config.Config
	HttpClient *http.Client
}

func GetDependencies() (*Dependencies, error) {
	getConfig, err := config.GetConfig()
	if err != nil {
		return nil, err
	}
	newDBConfig := dbConfig.NewDBConnection(getConfig)
	db, errDB := newDBConfig.GetDBConnection()
	if errDB != nil {
		return nil, errDB
	}
	httpClient := CustomHttp.NewHttpClient()

	return &Dependencies{
		DB:         db,
		Config:     getConfig,
		HttpClient: httpClient,
	}, nil
}
