package repository

import (
	"errors"
	"github.com/hudayberdipolat/go-url-shortener/internal/models"
	"gorm.io/gorm"
)

type authRepositoryImp struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return authRepositoryImp{
		db: db,
	}
}

func (a authRepositoryImp) Create(user models.User) error {
	if err := a.db.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("Bu username ady eýýäm ulanylýar!!!")
		}
		return err
	}
	return nil
}

func (a authRepositoryImp) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := a.db.Where("username =?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
