package handlers

import (
	// "net/http"
	// "show-calendar/request"
	"net/http"
	"show-calendar/request"
	"show-calendar/service"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	baseHandler Handler
	service     service.EventService
}

func (handler *EventHandler) Create(c *gin.Context) {
	var request request.CreateEventRequest
	if err := c.ShouldBind(&request); err != nil {
		handler.baseHandler.handleError(c, err)
		return
	}
	event, err := handler.service.Create(&request)
	handler.baseHandler.handleErrorAndReturn(c, err, func() {
		handler.baseHandler.sendResponse(c, http.StatusCreated, "成功", map[string]interface{}{"evnet": event})
	})
}

func NewEventHandler(handler Handler, service service.EventService) EventHandler {
	return EventHandler{handler, service}
}
