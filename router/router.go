package router

import (
	"show-calendar/handlers"
	"show-calendar/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(registerHandler handlers.RegisterHandler, authenticateHandler handlers.AuthenticateHandler, showHandler handlers.ShowHandler, eventHandler handlers.EventHandler) *gin.Engine {
	router := gin.Default()
	apis := router.Group("api")
	{
		apis.POST("/register", registerHandler.Create)
		apis.POST("/login", authenticateHandler.Login)
		apis.GET("/shows/:id", showHandler.Show)
		apis.GET("/shows/:id/events", eventHandler.GetByShowId)
		apis.GET("/events/home", eventHandler.GetLatestEvent)
	}
	authorized := apis.Group("/")
	authorized.Use(middleware.AuthRequired())
	{
		authorized.POST("/shows", showHandler.Create)
		authorized.POST("/events", eventHandler.Create)
	}
	return router
}
