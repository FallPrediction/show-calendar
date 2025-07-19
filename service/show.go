package service

import (
	"souflair/models"
	"souflair/repository"
	"souflair/request"
	"souflair/utils"
	"strings"
	"time"

	"gorm.io/datatypes"
)

type ShowService struct {
	repository repository.ShowRepository
}

func (s *ShowService) Show(id string) (models.Show, error) {
	return s.repository.Show(id)
}

func (s *ShowService) CreateShowAndEvent(request *request.CreateShowRequest) (models.Show, models.Event, error) {
	startDate := s.getStartDate(request.StartDate)
	endDate := s.getEndDate(request.EndDate)
	meta, err := (&utils.OpenGraph{}).Fetch(request.TicketUrl)
	if err != nil {
		return models.Show{}, models.Event{}, err
	}
	show := models.Show{
		Name:      request.Name,
		TicketUrl: request.TicketUrl,
		StartDate: startDate,
		EndDate:   endDate,
	}
	event := models.Event{
		ShowId:        show.Id,
		Name:          s.getName(meta, request.Name),
		OgImage:       meta.Image,
		OgUrl:         s.getUrl(meta.Url, request.TicketUrl),
		OgTitle:       meta.Title,
		OgDescription: meta.Description,
		StartDate:     startDate,
		EndDate:       endDate,
	}
	return show, event, s.repository.CreateShowAndEvent(&show, &event)
}

// If the parameter is empty, return today's date.
func (s *ShowService) getStartDate(startDate string) datatypes.Date {
	var result time.Time
	if startDate != "" {
		result, _ = time.Parse(time.DateOnly, startDate)
	} else {
		result = time.Now()
	}
	return datatypes.Date(result)
}

// If the parameter is empty, return the date in 365 days.
func (s *ShowService) getEndDate(endDate string) datatypes.Date {
	var result time.Time
	if endDate != "" {
		result, _ = time.Parse(time.DateOnly, endDate)
	} else {
		result = time.Now().Add(time.Hour * 24 * 365)
	}
	return datatypes.Date(result)
}

func (s *ShowService) getName(meta utils.OpenGraphMeta, showName string) string {
	description, _, _ := strings.Cut(meta.Description, "\n")
	name := meta.Title
	if len(description) > 0 {
		name += " " + description
	}
	if len(name) == 0 {
		name = showName
	}
	nameRunes := []rune(name)
	return string(nameRunes[:min(50, len(nameRunes))])
}

func (s *ShowService) getUrl(eventUrl string, ticketUrl string) string {
	if len(eventUrl) == 0 {
		return ticketUrl
	}
	return eventUrl
}

func NewShowService(repo repository.ShowRepository) ShowService {
	return ShowService{repo}
}
