package handlers

import (
	"net/http"
	"show-calendar/request"
	"show-calendar/resource"
	"show-calendar/service"

	"github.com/gin-gonic/gin"
)

type ShowHandler struct {
	baseHandler Handler
	service     service.ShowService
}

func (handler *ShowHandler) Show(c *gin.Context) {
	id := c.Param("id")
	show, err := handler.service.Show(id)
	handler.baseHandler.handleErrorAndReturn(c, err, func() {
		resourceObj := resource.NewShow(show)
		handler.baseHandler.sendResponse(c, http.StatusOK, "成功", resourceObj.ToMap())
	})
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
