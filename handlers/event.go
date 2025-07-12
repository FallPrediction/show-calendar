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

func (h *EventHandler) GetByShowId(c *gin.Context) {
	if err := c.ShouldBindUri(&struct {
		id uint32 `form:"id" binding:"required,exists=shows id"`
	}{}); err != nil {
		h.baseHandler.handleError(c, err)
		return
	}
	var request request.GetEventByShowIdRequest
	if err := c.ShouldBind(&request); err != nil {
		h.baseHandler.handleError(c, err)
		return
	}

	events, count, err := h.service.GetByShowId(c.Param("id"), &request)

	h.baseHandler.handleErrorAndReturn(c, err, func() {
		resource := resource.NewEventResource()
		h.baseHandler.sendResponseWithPagination(c, http.StatusOK, "成功", resource.ToSlice(events), request.CurrentPage, request.PerPage, int(count))
	})
}

func (h *EventHandler) GetLatestEvents(c *gin.Context) {
	events, err := h.service.GetLatestEvents()

	h.baseHandler.handleErrorAndReturn(c, err, func() {
		resource := resource.NewEventResource()
		h.baseHandler.sendResponse(c, http.StatusOK, "成功", resource.ToSlice(events))
	})
}

func (h *EventHandler) GetHomeEvents(c *gin.Context) {
	events, err := h.service.GetLatestEventEachShow()
	h.baseHandler.handleErrorAndReturn(c, err, func() {
		resource := resource.NewEventResource()
		h.baseHandler.sendResponse(c, http.StatusOK, "成功", resource.ToSlice(events))
	})
}

func (h *EventHandler) Create(c *gin.Context) {
	var request request.CreateEventRequest
	if err := c.ShouldBind(&request); err != nil {
		h.baseHandler.handleError(c, err)
		return
	}
	event, err := h.service.Create(&request)
	h.baseHandler.handleErrorAndReturn(c, err, func() {
		resource := resource.NewEventResource()
		h.baseHandler.sendResponse(c, http.StatusCreated, "成功", map[string]interface{}{"evnet": resource.ToMap(event)})
	})
}

func NewEventHandler(handler Handler, service service.EventService) EventHandler {
	return EventHandler{handler, service}
}
