package service

import (
	"errors"
	"github.com/hudayberdipolat/go-url-shortener/internal/domain/user/dto"
	"github.com/hudayberdipolat/go-url-shortener/internal/domain/user/repository"
	"github.com/hudayberdipolat/go-url-shortener/internal/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type userServiceImp struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return userServiceImp{
		userRepo: repo,
	}
}

func (u userServiceImp) GetUserData(userID int) (*dto.UserResponse, error) {
	user, err := u.userRepo.GetUser(userID)
	if err != nil {
		return nil, err
	}
	userResponse := dto.NewUserResponse(*user)
	return &userResponse, nil
}

func (u userServiceImp) UpdateUser(userID int, update dto.UpdateUserData) error {
	updateUser := models.User{
		Username:  update.Username,
		FullName:  update.FullName,
		UpdatedAt: time.Now(),
	}
	if err := u.userRepo.ChangeUserData(userID, updateUser); err != nil {
		return err
	}
	return nil
}

func (u userServiceImp) ChangePassword(userID int, changePassword dto.ChangeUserPassword) error {
	user, err := u.userRepo.GetUser(userID)
	if err != nil {
		return err
	}
	errOldPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(changePassword.OldPassword))
	if errOldPassword != nil {
		return errors.New("Köne password nädogry!!!")
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(changePassword.Password), 16)
	if errChangePassword := u.userRepo.ChangeUserPassword(user.ID, string(hashPassword)); errChangePassword != nil {
		return err
	}
	return nil
}
