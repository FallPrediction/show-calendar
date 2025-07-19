package handlers

import (
	"net/http"
	"souflair/request"
	"souflair/resource"
	"souflair/service"

	"github.com/gin-gonic/gin"
)

type ShowHandler struct {
	baseHandler Handler
	service     service.ShowService
}

func (h *ShowHandler) Show(c *gin.Context) {
	id := c.Param("id")
	show, err := h.service.Show(id)
	h.baseHandler.handleErrorAndReturn(c, err, func() {
		resourceObj := resource.NewShow(show)
		h.baseHandler.sendResponse(c, http.StatusOK, "成功", resourceObj.ToMap())
	})
}

func (h *ShowHandler) Create(c *gin.Context) {
	var request request.CreateShowRequest
	if err := c.ShouldBind(&request); err != nil {
		h.baseHandler.handleError(c, err)
		return
	}
	show, event, err := h.service.CreateShowAndEvent(&request)
	h.baseHandler.handleErrorAndReturn(c, err, func() {
		resource := resource.NewEventResource()
		h.baseHandler.sendResponse(c, http.StatusCreated, "成功", map[string]interface{}{"show": show, "evnet": resource.ToMap(event)})
	})
}

func NewShowHandler(handler Handler, service service.ShowService) ShowHandler {
	return ShowHandler{handler, service}
}
