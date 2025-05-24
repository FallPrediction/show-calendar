package repository

import (
	"show-calendar/models"

	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func (repository *EventRepository) Create(event *models.Event) error {
	return repository.db.Create(&event).Error
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return EventRepository{db}
}
