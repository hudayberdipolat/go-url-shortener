package dto

import "github.com/hudayberdipolat/go-url-shortener/internal/models"

type AuthResponse struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	FullName    string `json:"full_name"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	AccessToken string `json:"access_token"`
}

func NewAuthResponse(user models.User, accessToken string) AuthResponse {
	return AuthResponse{
		ID:          user.ID,
		Username:    user.Username,
		FullName:    user.FullName,
		CreatedAt:   user.CreatedAt.Format("01-02-2006"),
		UpdatedAt:   user.UpdatedAt.Format("01-02-2006"),
		AccessToken: accessToken,
	}
}
