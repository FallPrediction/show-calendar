package router

import (
	"show-calendar/handlers"

	"github.com/gin-gonic/gin"
)

func NewRouter(registerHandler handlers.RegisterHandler, authenticateHandler handlers.AuthenticateHandler) *gin.Engine {
	router := gin.Default()
	apis := router.Group("api")
	{
		apis.POST("/register", registerHandler.Create)
		apis.POST("/login", authenticateHandler.Login)
	}
	return router
}
