package repository

import (
	"souflair/models"
	"souflair/request"
	"time"

	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func (r *EventRepository) GetByShowId(id string, request *request.GetEventByShowIdRequest) ([]models.Event, int64, error) {
	var events []models.Event
	var count int64
	r.db.Model(&models.Event{}).Count(&count)
	err := r.db.Order("start_date desc").Limit(request.PerPage).
		Offset((request.CurrentPage-1)*request.PerPage).
		Where("show_id = ?", id).
		Find(&events).Error
	return events, count, err
}

func (r *EventRepository) Index(startDate time.Time) ([]models.Event, error) {
	var events []models.Event
	err := r.db.Order("start_date desc").
		Where("end_date >= ?", startDate).
		Find(&events).Error
	return events, err
}

func (r *EventRepository) GetLatestEventEachShow() ([]models.Event, error) {
	var events []models.Event
	err := r.db.Table("events AS e").
		Joins("INNER JOIN shows AS s ON e.show_id = s.id").
		Where("NOW() <= s.end_date").
		Where(
			"(s.id, e.updated_at) IN (?)",
			r.db.Select("show_id", "MAX(updated_at)").Table("events").Group("show_id")).
		Find(&events).Error
	return events, err
}

func (r *EventRepository) Create(event *models.Event) error {
	return r.db.Create(&event).Error
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return EventRepository{db}
}
