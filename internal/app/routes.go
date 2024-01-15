package app

import (
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	service "github.com/Jumaniyozov/go-rest-template/internal/services"
	"github.com/Jumaniyozov/go-rest-template/internal/transport/rest"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
)

func SetupoRouter(c *config.Config, l *zerolog.Logger, s service.ServiceI) *httprouter.Router {
	router := httprouter.New()
	handlers := rest.NewHandler(c, l, s)

	router.GET("/swagger/:any", handlers.SwaggerHandler.Init)
	router.GET("/users", handlers.UserHandler.ListAllUsers)

	return router
}
