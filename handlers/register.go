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

func (handler *RegisterHandler) Create(c *gin.Context) {
	var request request.RegisterRequest
	if err := c.ShouldBind(&request); err != nil {
		handler.baseHandler.handleError(c, err)
		return
	}
	handler.baseHandler.handleErrorAndReturn(c, handler.service.Create(&request), func() {
		handler.baseHandler.sendResponse(c, http.StatusCreated, "註冊成功", nil)
	})
}

func NewRegisterHandler(handler Handler, service service.RegisterService) RegisterHandler {
	return RegisterHandler{handler, service}
}
