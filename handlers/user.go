package handlers

import (
	"net/http"
	"show-calendar/request"
	"show-calendar/service"
	"show-calendar/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	baseHandler Handler
	service     service.UserService
}

func (handler *UserHandler) LikeShow(c *gin.Context) {
	var request request.UserLikeShowRequest
	if err := c.ShouldBind(&request); err != nil {
		handler.baseHandler.handleError(c, err)
		return
	}
	userData, exists := c.Get("userData")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "未登入",
		})
		c.Abort()
		return
	}
	handler.baseHandler.handleErrorAndReturn(c, handler.service.LikeShow(&request, userData.(utils.UserData).UserId), func() {
		handler.baseHandler.sendResponse(c, http.StatusCreated, "新增成功", nil)
	})
}

func NewUserHandler(handler Handler, service service.UserService) UserHandler {
	return UserHandler{handler, service}
}
