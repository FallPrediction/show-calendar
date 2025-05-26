package repository

import (
	"show-calendar/models"

	"gorm.io/gorm"
)

type ShowRepository struct {
	db *gorm.DB
}

func (repository *ShowRepository) CreateShowAndEvent(show *models.Show, event *models.Event) error {
	return repository.db.Transaction(func(tx *gorm.DB) error {
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
