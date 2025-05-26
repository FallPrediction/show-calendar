package handlers

import (
	"net/http"
	"show-calendar/request"
	"show-calendar/service"

	"github.com/gin-gonic/gin"
)

type ShowHandler struct {
	baseHandler Handler
	service     service.ShowService
}

func (handler *ShowHandler) Create(c *gin.Context) {
	var request request.CreateShowRequest
	if err := c.ShouldBind(&request); err != nil {
		handler.baseHandler.handleError(c, err)
		return
	}
	show, event, err := handler.service.CreateShowAndEvent(&request)
	handler.baseHandler.handleErrorAndReturn(c, err, func() {
		handler.baseHandler.sendResponse(c, http.StatusCreated, "成功", map[string]interface{}{"show": show, "evnet": event})
	})
}

func NewShowHandler(handler Handler, service service.ShowService) ShowHandler {
	return ShowHandler{handler, service}
}
