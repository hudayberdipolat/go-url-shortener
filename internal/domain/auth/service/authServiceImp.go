package service

import (
	"errors"
	"github.com/hudayberdipolat/go-url-shortener/internal/domain/auth/dto"
	"github.com/hudayberdipolat/go-url-shortener/internal/domain/auth/repository"
	"github.com/hudayberdipolat/go-url-shortener/internal/models"
	"github.com/hudayberdipolat/go-url-shortener/pkg/jwtToken"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type authServiceImp struct {
	authRepo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return authServiceImp{
		authRepo: repo,
	}
}

func (a authServiceImp) Register(registerRequest dto.RegisterRequest) (*dto.AuthResponse, error) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), 16)
	user := models.User{
		Username:  registerRequest.Username,
		FullName:  registerRequest.FullName,
		Password:  string(hashPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	// register user
	if err := a.authRepo.Create(user); err != nil {
		return nil, err
	}
	getUser, err := a.authRepo.GetUserByUsername(registerRequest.Username)
	if err != nil {
		return nil, err
	}
	// generate token
	accessToken, err := jwtToken.GenerateToken(getUser.Username, getUser.ID)
	registerResponse := dto.NewAuthResponse(*getUser, accessToken)
	return &registerResponse, nil
}

func (a authServiceImp) Login(loginRequest dto.LoginRequest) (*dto.AuthResponse, error) {
	getUser, err := a.authRepo.GetUserByUsername(loginRequest.Username)
	if err != nil {
		return nil, errors.New("username ýa-da password nädogry!!!")
	}
	errLoginPassword := bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(loginRequest.Password))
	if errLoginPassword != nil {
		return nil, errors.New("username ýa-da password nädogry!!!")
	}
	// generate token
	accessToken, err := jwtToken.GenerateToken(getUser.Username, getUser.ID)
	loginResponse := dto.NewAuthResponse(*getUser, accessToken)
	return &loginResponse, nil
}
