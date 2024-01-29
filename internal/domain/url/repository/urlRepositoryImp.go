package repository

import (
	"errors"
	"github.com/hudayberdipolat/go-url-shortener/internal/models"
	"gorm.io/gorm"
)

type urlRepositoryImp struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) UrlRepository {
	return urlRepositoryImp{
		db: db,
	}
}

func (u urlRepositoryImp) GetAllUrl(userID int) ([]models.Url, error) {
	var urls []models.Url
	if err := u.db.Where("user_id=?", userID).Find(&urls).Error; err != nil {
		return nil, err
	}
	return urls, nil
}

func (u urlRepositoryImp) GetUrlByID(userID, urlID int) (*models.Url, error) {
	var url models.Url
	if err := u.db.Where("id =?", urlID).Where("user_id=?", userID).First(&url).Error; err != nil {
		return nil, err
	}
	return &url, nil
}

func (u urlRepositoryImp) Create(url models.Url) error {
	if err := u.db.Create(&url).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("short url döretmekde näsazlyk döredi!!!")
		}
		return err
	}
	return nil
}

func (u urlRepositoryImp) Delete(userID, urlID int) error {
	var url models.Url
	err := u.db.Where("id =?", urlID).Where("user_id=?", userID).Unscoped().Delete(&url).Error
	if err != nil {
		return err
	}
	return nil
}
