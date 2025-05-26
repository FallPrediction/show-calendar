package service

import (
	"show-calendar/models"
	"show-calendar/repository"
	"show-calendar/request"
	"show-calendar/utils"
	"time"

	"gorm.io/datatypes"
)

type ShowService struct {
	repository repository.ShowRepository
}

func (service *ShowService) CreateShowAndEvent(request *request.CreateShowRequest) (models.Show, models.Event, error) {
	startDate, _ := time.Parse(time.DateOnly, request.StartDate)
	endDate, _ := time.Parse(time.DateOnly, request.EndDate)
	meta, err := (&utils.OpenGraph{}).Fetch(request.TicketUrl)
	if err != nil {
		return models.Show{}, models.Event{}, err
	}
	show := models.Show{
		Name:      request.Name,
		TicketUrl: request.TicketUrl,
		StartDate: datatypes.Date(startDate),
		EndDate:   datatypes.Date(endDate),
	}
	event := models.Event{
		OgImage:       meta.Image,
		OgUrl:         meta.Url,
		OgTitle:       meta.Title,
		OgDescription: meta.Description,
		StartDate:     datatypes.Date(startDate),
		EndDate:       datatypes.Date(endDate),
	}
	return show, event, service.repository.CreateShowAndEvent(&show, &event)
}

func NewShowService(repo repository.ShowRepository) ShowService {
	return ShowService{repo}
}
