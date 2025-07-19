package handlers

import (
	"net/http"
	"souflair/request"
	"souflair/service"
	"souflair/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	baseHandler Handler
	service     service.UserService
}

func (h *UserHandler) LikeShow(c *gin.Context) {
	var request request.UserLikeShowRequest
	if err := c.ShouldBind(&request); err != nil {
		h.baseHandler.handleError(c, err)
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
	h.baseHandler.handleErrorAndReturn(c, h.service.LikeShow(&request, userData.(utils.UserData).UserId), func() {
		h.baseHandler.sendResponse(c, http.StatusCreated, "新增成功", nil)
	})
}

func (h *UserHandler) Unsubscribe(c *gin.Context) {
	h.baseHandler.handleErrorAndReturn(c, h.service.Unsubscribe(c.Query("token")), func() {
		h.baseHandler.sendResponse(c, http.StatusOK, "取消訂閱成功", nil)
	})
}

func NewUserHandler(handler Handler, service service.UserService) UserHandler {
	return UserHandler{handler, service}
}
