package repository

import (
	"souflair/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(&user).Error
}

func (r *UserRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *UserRepository) LikeShow(userId, showId uint32) error {
	var count int64
	if r.db.Table("user_shows").Where("user_id", userId).Where("show_id", showId).Count(&count); count > 0 {
		return nil
	}
	return r.db.Model(&models.User{Id: userId}).Association("Shows").Append(&models.Show{Id: showId})
}

func (r *UserRepository) Unsubscribe(email string) error {
	return r.db.Table("users").Where("email = ?", email).Update("subscribe", false).Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db}
}
