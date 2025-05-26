package resource

import (
	"show-calendar/models"
)

type EventSlice struct {
	modelSlice []models.Event
}

func (resource *EventSlice) ToSlice() []map[string]interface{} {
	result := make([]map[string]interface{}, len(resource.modelSlice))
	for i, m := range resource.modelSlice {
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

func NewEventSlice(models []models.Event) EventSlice {
	return EventSlice{models}
}
