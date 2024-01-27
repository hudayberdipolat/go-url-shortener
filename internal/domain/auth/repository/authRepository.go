package repository

import "github.com/hudayberdipolat/go-url-shortener/internal/models"

type AuthRepository interface {
	Create(user models.User) error
	GetUserByUsername(username string) (*models.User, error)
}
