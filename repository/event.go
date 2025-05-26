package repository

import (
	"show-calendar/models"
	"show-calendar/request"

	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func (repository *EventRepository) GetByShowId(id string, request *request.GetEventByShowIdRequest) ([]models.Event, int64, error) {
	var events []models.Event
	var count int64
	repository.db.Model(&models.Event{}).Count(&count)
	err := repository.db.Order("updated_at desc").Limit(request.PerPage).
		Offset((request.CurrentPage-1)*request.PerPage).
		Where("show_id = ?", id).
		Find(&events).Error
	return events, count, err
}

func (repository *EventRepository) Create(event *models.Event) error {
	return repository.db.Create(&event).Error
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return EventRepository{db}
}
