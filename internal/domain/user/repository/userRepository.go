package repository

import "github.com/hudayberdipolat/go-url-shortener/internal/models"

type UserRepository interface {
	GetUser(userID int) (*models.User, error)
	ChangeUserData(userId int, user models.User) error
	ChangeUserPassword(userId int, password string) error
}
