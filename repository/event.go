package repository

import (
	"show-calendar/models"
	"show-calendar/request"
	"time"

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

func (repository *EventRepository) Index(startDate, endDate time.Time) ([]models.Event, error) {
	var events []models.Event
	err := repository.db.Order("start_date asc").
		Where("start_date <= ?", endDate).
		Where("end_date >= ?", startDate).
		Find(&events).Error
	return events, err
}

func (repository *EventRepository) GetLatestEventEachShow() ([]models.Event, error) {
	var events []models.Event
	err := repository.db.Table("events AS e").
		Joins("INNER JOIN shows AS s ON e.show_id = s.id").
		Where("NOW() <= s.end_date").
		Where(
			"(s.id, e.updated_at) IN (?)",
			repository.db.Select("show_id", "MAX(updated_at)").Table("events").Group("show_id")).
		Find(&events).Error
	return events, err
}

func (repository *EventRepository) Create(event *models.Event) error {
	return repository.db.Create(&event).Error
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return EventRepository{db}
}
