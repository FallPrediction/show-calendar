package router

import (
	"show-calendar/handlers"

	"github.com/gin-gonic/gin"
)

func NewRouter(registerHandler handlers.RegisterHandler) *gin.Engine {
	router := gin.Default()
	apis := router.Group("api")
	{
		apis.POST("/register", registerHandler.Create)
	}
	return router
}
