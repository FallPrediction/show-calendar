package resource

import "show-calendar/models"

type Show struct {
	model models.Show
}

func (r *Show) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Id":        r.model.Id,
		"Name":      r.model.Name,
		"TicketUrl": r.model.TicketUrl,
		"StartDate": r.model.StartDate,
		"EndDate":   r.model.EndDate,
	}
}

func NewShow(model models.Show) Show {
	return Show{model}
}
