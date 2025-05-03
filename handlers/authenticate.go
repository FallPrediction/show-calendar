package handlers

import (
	"show-calendar/request"
	"show-calendar/service"

	"github.com/gin-gonic/gin"
)

type AuthenticateHandler struct {
	baseHandler Handler
	service     service.AuthenticateService
}

func (handler *AuthenticateHandler) Login(c *gin.Context) {
	var request request.LoginRequest
	if err := c.ShouldBind(&request); err != nil {
		handler.baseHandler.handleError(c, err)
		return
	}
	token, err := handler.service.Login(&request)
	handler.baseHandler.handleErrorAndReturn(c, err, func() {
		handler.baseHandler.sendResponse(c, 200, "登入成功", map[string]string{"access_token": token, "token_type": "bearer"})
	})
}

func NewAuthenticateHandler(handler Handler, service service.AuthenticateService) AuthenticateHandler {
	return AuthenticateHandler{handler, service}
}
