package dto

import "github.com/hudayberdipolat/go-url-shortener/internal/models"

type UserResponse struct {
	ID        int
	Username  string
	FullName  string
	CreatedAt string
	UpdatedAt string
}

func NewUserResponse(user models.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		FullName:  user.FullName,
		CreatedAt: user.CreatedAt.Format("01-02-2006"),
		UpdatedAt: user.UpdatedAt.Format("01-02-2006"),
	}
}
