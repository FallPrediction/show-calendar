package router

import (
	"show-calendar/handlers"
	"show-calendar/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(registerHandler handlers.RegisterHandler, authenticateHandler handlers.AuthenticateHandler, showHandler handlers.ShowHandler) *gin.Engine {
	router := gin.Default()
	apis := router.Group("api")
	{
		apis.POST("/register", registerHandler.Create)
		apis.POST("/login", authenticateHandler.Login)
		apis.GET("/shows/:id", showHandler.Show)
	}
	authorized := apis.Group("/")
	authorized.Use(middleware.AuthRequired())
	{
		authorized.POST("/shows", showHandler.Create)
	}
	return router
}
