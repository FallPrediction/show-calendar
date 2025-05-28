package handlers

import (
	"net/http"
	"show-calendar/request"
	"show-calendar/resource"
	"show-calendar/service"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	baseHandler Handler
	service     service.EventService
}

func (handler *EventHandler) GetByShowId(c *gin.Context) {
	if err := c.ShouldBindUri(&struct {
		id uint32 `form:"id" binding:"required,exists=shows id"`
	}{}); err != nil {
		handler.baseHandler.handleError(c, err)
		return
	}
	var request request.GetEventByShowIdRequest
	if err := c.ShouldBind(&request); err != nil {
		handler.baseHandler.handleError(c, err)
		return
	}

	events, count, err := handler.service.GetByShowId(c.Param("id"), &request)

	handler.baseHandler.handleErrorAndReturn(c, err, func() {
		resource := resource.NewEventSlice(events)
		handler.baseHandler.sendResponseWithPagination(c, http.StatusOK, "成功", resource.ToSlice(), request.CurrentPage, request.PerPage, int(count))
	})
}

func (handler *EventHandler) Index(c *gin.Context) {
	var request request.IndexEventsRequest
	if err := c.ShouldBind(&request); err != nil {
		handler.baseHandler.handleError(c, err)
		return
	}

	events, err := handler.service.Index(&request)

	handler.baseHandler.handleErrorAndReturn(c, err, func() {
		resource := resource.NewEventSlice(events)
		handler.baseHandler.sendResponse(c, http.StatusOK, "成功", resource.ToSlice())
	})
}

func (handler *EventHandler) GetLatestEvent(c *gin.Context) {
	events, err := handler.service.GetLatestEvent()
	handler.baseHandler.handleErrorAndReturn(c, err, func() {
		resource := resource.NewEventSlice(events)
		handler.baseHandler.sendResponse(c, http.StatusOK, "成功", resource.ToSlice())
	})
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
