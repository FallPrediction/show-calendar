package repository

import (
	"souflair/models"

	"gorm.io/gorm"
)

type ShowRepository struct {
	db *gorm.DB
}

func (r *ShowRepository) Show(id string) (models.Show, error) {
	var show models.Show
	err := r.db.Where("id = ?", id).First(&show).Error
	return show, err
}

func (r *ShowRepository) CreateShowAndEvent(show *models.Show, event *models.Event) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&show).Error; err != nil {
			return err
		}
		event.ShowId = show.Id
		if err := tx.Create(&event).Error; err != nil {
			return err
		}
		return nil
	})
}

func NewShowRepository(db *gorm.DB) ShowRepository {
	return ShowRepository{db}
}
