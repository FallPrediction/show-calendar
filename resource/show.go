package resource

import "show-calendar/models"

type Show struct {
	model models.Show
}

func (resource *Show) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Id":        resource.model.Id,
		"Name":      resource.model.Name,
		"TicketUrl": resource.model.TicketUrl,
		"StartDate": resource.model.StartDate,
		"EndDate":   resource.model.EndDate,
	}
}

func NewShow(model models.Show) Show {
	return Show{model}
}
