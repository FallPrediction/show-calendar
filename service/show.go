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
		Name:          s.getName(meta),
		OgImage:       meta.Image,
		OgUrl:         meta.Url,
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

func (service *ShowService) getName(meta utils.OpenGraphMeta) string {
	description, _, _ := strings.Cut(meta.Description, "\n")
	s := []rune(meta.Title + " " + description)
	return string(s[:min(50, len(s))])
}

func NewShowService(repo repository.ShowRepository) ShowService {
	return ShowService{repo}
}
