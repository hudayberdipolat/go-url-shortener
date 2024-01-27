package service

import "github.com/hudayberdipolat/go-url-shortener/internal/domain/user/dto"

type UserService interface {
	GetUserData(userID int) (*dto.UserResponse, error)
	UpdateUser(userID int, update dto.UpdateUserData) error
	ChangePassword(userID int, changePassword dto.ChangeUserPassword) error
}
