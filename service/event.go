package service

import (
	"show-calendar/models"
	"show-calendar/repository"
	"show-calendar/request"
	"show-calendar/utils"
	"strings"
	"time"

	"gorm.io/datatypes"
)

type EventService struct {
	repository repository.EventRepository
}

func (service *EventService) GetByShowId(id string, request *request.GetEventByShowIdRequest) ([]models.Event, int64, error) {
	if request.CurrentPage == 0 {
		request.CurrentPage = 1
	}
	if request.PerPage == 0 {
		request.PerPage = 20
	}
	return service.repository.GetByShowId(id, request)
}

func (service *EventService) GetLatestEvents(request *request.LatestEventsRequest) ([]models.Event, error) {
	startDate, err := time.Parse(time.DateOnly, request.Date)
	if err != nil {
		return nil, err
	}
	startDate = startDate.AddDate(0, 0, -startDate.Day()+1)
	endDate := startDate.AddDate(0, 1, -startDate.Day())
	return service.repository.Index(startDate, endDate)
}

func (service *EventService) GetLatestEventEachShow() ([]models.Event, error) {
	return service.repository.GetLatestEventEachShow()
}

func (service *EventService) Create(request *request.CreateEventRequest) (models.Event, error) {
	meta, err := (&utils.OpenGraph{}).Fetch(request.Url)
	if err != nil {
		return models.Event{}, err
	}
	name := service.getName(request.Name, meta)
	event := models.Event{
		ShowId:        request.ShowId,
		Name:          name,
		OgImage:       meta.Image,
		OgUrl:         meta.Url,
		OgTitle:       meta.Title,
		OgDescription: meta.Description,
		StartDate:     service.getStartDate(request.StartDate),
		EndDate:       service.getEndDate(request.EndDate),
	}
	return event, service.repository.Create(&event)
}

// If the parameter is empty, return today's date.
func (service *EventService) getStartDate(startDate string) datatypes.Date {
	var result time.Time
	if startDate != "" {
		result, _ = time.Parse(time.DateOnly, startDate)
	} else {
		result = time.Now()
	}
	return datatypes.Date(result)
}

// If the parameter is empty, return the date in 30 days.
func (service *EventService) getEndDate(endDate string) datatypes.Date {
	var result time.Time
	if endDate != "" {
		result, _ = time.Parse(time.DateOnly, endDate)
	} else {
		result = time.Now().Add(time.Hour * 24 * 30)
	}
	return datatypes.Date(result)
}

func (service *EventService) getName(name string, meta utils.OpenGraphMeta) string {
	if name == "" {
		description, _, _ := strings.Cut(meta.Description, "\n")
		s := []rune(meta.Title + " " + description)
		return string(s[:min(50, len(s))])
	}
	return name
}

func NewEventService(repo repository.EventRepository) EventService {
	return EventService{repo}
}
