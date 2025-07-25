package main

import (
	"souflair/config"
	"souflair/handlers"
	"souflair/initialize"
	"souflair/repository"
	"souflair/router"
	"souflair/rules"
	"souflair/service"
)

func main() {
	translator := initialize.NewTranslator()
	handler := handlers.NewBaseHandler(translator)
	pg := config.NewPg()
	db := initialize.NewDB(pg)
	userRepository := repository.NewUserRepository(db)
	registerService := service.NewRegisterService(userRepository)
	registerHandler := handlers.NewRegisterHandler(handler, registerService)
	authenticateService := service.NewAuthenticateService(userRepository)
	authenticateHandler := handlers.NewAuthenticateHandler(handler, authenticateService)
	showRepository := repository.NewShowRepository(db)
	showService := service.NewShowService(showRepository)
	showHandler := handlers.NewShowHandler(handler, showService)
	eventRepository := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepository)
	eventHandler := handlers.NewEventHandler(handler, eventService)
	userService := service.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(handler, userService)
	rules.BindValidator(translator)
	router := router.NewRouter(registerHandler, authenticateHandler, showHandler, eventHandler, userHandler)

	router.Run()
}
