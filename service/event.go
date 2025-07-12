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

func (s *EventService) GetByShowId(id string, request *request.GetEventByShowIdRequest) ([]models.Event, int64, error) {
	if request.CurrentPage == 0 {
		request.CurrentPage = 1
	}
	if request.PerPage == 0 {
		request.PerPage = 20
	}
	return s.repository.GetByShowId(id, request)
}

func (s *EventService) GetLatestEvents() ([]models.Event, error) {
	lastMonth := time.Now().AddDate(0, -1, 0)
	startDate := lastMonth.AddDate(0, 0, -lastMonth.Day()+1)
	return s.repository.Index(startDate)
}

func (s *EventService) GetLatestEventEachShow() ([]models.Event, error) {
	return s.repository.GetLatestEventEachShow()
}

func (s *EventService) Create(request *request.CreateEventRequest) (models.Event, error) {
	meta, err := (&utils.OpenGraph{}).Fetch(request.Url)
	if err != nil {
		return models.Event{}, err
	}
	name := s.getName(request.Name, meta)
	startDate := s.getStartDate(request.StartDate)
	endDate := s.getEndDate(startDate, request.EndDate)
	event := models.Event{
		ShowId:        request.ShowId,
		Name:          name,
		OgImage:       meta.Image,
		OgUrl:         meta.Url,
		OgTitle:       meta.Title,
		OgDescription: meta.Description,
		StartDate:     startDate,
		EndDate:       endDate,
	}
	return event, s.repository.Create(&event)
}

// If the parameter is empty, return today's date.
func (s *EventService) getStartDate(startDate string) datatypes.Date {
	var result time.Time
	if startDate != "" {
		result, _ = time.Parse(time.DateOnly, startDate)
	} else {
		result = time.Now()
	}
	return datatypes.Date(result)
}

// If the parameter is empty, return start date.
func (s *EventService) getEndDate(startDate datatypes.Date, endDate string) datatypes.Date {
	var result time.Time
	if endDate != "" {
		result, _ = time.Parse(time.DateOnly, endDate)
		return datatypes.Date(result)
	}
	return startDate
}

func (s *EventService) getName(name string, meta utils.OpenGraphMeta) string {
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
