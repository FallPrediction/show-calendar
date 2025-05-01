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

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db}
}
