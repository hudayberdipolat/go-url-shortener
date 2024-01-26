package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	DbConfig   dbConfig   `json:"db_config"`
	HttpServer httpServer `json:"http_server"`
}
type dbConfig struct {
	DbHost     string `json:"db_host"`
	DbPort     string `json:"db_port"`
	DbUser     string `json:"db_user"`
	DbPassword string `json:"db_password"`
	DbName     string `json:"db_name"`
	DbSllMode  string `json:"db_sll_mode"`
}
type httpServer struct {
	ServerHost string `json:"http_host"`
	ServerPort string `json:"http_port"`
	AppName    string `json:"app_name"`
	AppHeader  string `json:"app_header"`
}

func GetConfig() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig("../.env", &cfg)
	if err != nil {
		return nil, err
	}

	cfg = Config{
		DbConfig: dbConfig{
			DbHost:     os.Getenv("DB_HOST"),
			DbPort:     os.Getenv("DB_PORT"),
			DbUser:     os.Getenv("DB_USER"),
			DbPassword: os.Getenv("DB_PASSWORD"),
			DbName:     os.Getenv("DB_NAME"),
			DbSllMode:  os.Getenv("DB_SLL_MODE"),
		},
		HttpServer: httpServer{
			ServerHost: os.Getenv("SERVER_HOST"),
			ServerPort: os.Getenv("SERVER_PORT"),
			AppName:    os.Getenv("APP_NAME"),
			AppHeader:  os.Getenv("APP_HEADER"),
		},
	}
	return &cfg, err
}
