package main

import (
	"show-calendar/config"
	"show-calendar/handlers"
	"show-calendar/initialize"
	"show-calendar/repository"
	"show-calendar/router"
	"show-calendar/rules"
	"show-calendar/service"
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
	rules.BindValidator(translator)
	router := router.NewRouter(registerHandler, authenticateHandler)

	router.Run()
}
