package resource

import (
	"show-calendar/config"
	"show-calendar/models"
)

type EventResource struct{}

func (r *EventResource) ToSlice(models []models.Event) []map[string]interface{} {
	result := make([]map[string]interface{}, len(models))
	for i, m := range models {
		if len(m.OgImage) > 0 {
			m.OgImage = config.Endpoint + m.OgImage
		}
		result[i] = map[string]interface{}{
			"Id":            m.Id,
			"Name":          m.Name,
			"ShowId":        m.ShowId,
			"OgImage":       m.OgImage,
			"OgUrl":         m.OgUrl,
			"OgTitle":       m.OgTitle,
			"OgDescription": m.OgDescription,
			"StartDate":     m.StartDate,
			"EndDate":       m.EndDate,
		}
	}
	return result
}

func (r *EventResource) ToMap(model models.Event) map[string]interface{} {
	if len(model.OgImage) > 0 {
		model.OgImage = config.Endpoint + model.OgImage
	}
	return map[string]interface{}{
		"Id":            model.Id,
		"Name":          model.Name,
		"ShowId":        model.ShowId,
		"OgImage":       model.OgImage,
		"OgUrl":         model.OgUrl,
		"OgTitle":       model.OgTitle,
		"OgDescription": model.OgDescription,
		"StartDate":     model.StartDate,
		"EndDate":       model.EndDate,
	}
}

func NewEventResource() EventResource {
	return EventResource{}
}
