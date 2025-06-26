package repository

import (
	"show-calendar/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (repository *UserRepository) Create(user *models.User) error {
	return repository.db.Create(&user).Error
}

func (repository *UserRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := repository.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (repository *UserRepository) LikeShow(userId, showId uint32) error {
	var count int64
	if repository.db.Table("user_shows").Where("user_id", userId).Where("show_id", showId).Count(&count); count > 0 {
		return nil
	}
	return repository.db.Model(&models.User{Id: userId}).Association("Shows").Append(&models.Show{Id: showId})
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db}
}
