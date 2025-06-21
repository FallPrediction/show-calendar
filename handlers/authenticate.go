package handlers

import (
	"errors"
	"net/http"
	"show-calendar/request"
	"show-calendar/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	token, expires, err := handler.service.Login(&request)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		handler.baseHandler.sendResponse(c, http.StatusNotFound, "該信箱尚未註冊", nil)
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "authorization",
		Value:    token,
		Path:     "/",
		Expires:  expires,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
		HttpOnly: true,
	})

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "login",
		Value:    "1",
		Path:     "/",
		Expires:  expires,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
	})
	handler.baseHandler.handleErrorAndReturn(c, err, func() {
		handler.baseHandler.sendResponse(c, http.StatusOK, "登入成功", nil)
	})
}

func (handler *AuthenticateHandler) Logout(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:   "authorization",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	http.SetCookie(c.Writer, &http.Cookie{
		Name:   "login",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	handler.baseHandler.sendResponse(c, http.StatusOK, "登出成功", nil)
}

func NewAuthenticateHandler(handler Handler, service service.AuthenticateService) AuthenticateHandler {
	return AuthenticateHandler{handler, service}
}
