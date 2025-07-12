package handlers

import (
	"net/http"
	"show-calendar/request"
	"show-calendar/service"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	baseHandler Handler
	service     service.RegisterService
}

func (h *RegisterHandler) Create(c *gin.Context) {
	var request request.RegisterRequest
	if err := c.ShouldBind(&request); err != nil {
		h.baseHandler.handleError(c, err)
		return
	}
	h.baseHandler.handleErrorAndReturn(c, h.service.Create(&request), func() {
		h.baseHandler.sendResponse(c, http.StatusCreated, "註冊成功", nil)
	})
}

func NewRegisterHandler(handler Handler, service service.RegisterService) RegisterHandler {
	return RegisterHandler{handler, service}
}
