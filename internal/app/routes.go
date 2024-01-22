package app

import (
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	contractService "github.com/Jumaniyozov/go-rest-template/internal/contracts/service"
	middlewares "github.com/Jumaniyozov/go-rest-template/internal/middleware"
	"github.com/Jumaniyozov/go-rest-template/internal/transport/rest"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
)

func SetupoRouter(c *config.Config, l *zerolog.Logger, s contractService.ServiceI) *httprouter.Router {
	router := httprouter.New()
	handlers := rest.NewHandler(c, l, s)

	router.GET("/swagger/:any", handlers.SwaggerHandler.Init)
	router.GET("/users", handlers.UserHandler.ListAllUsers)
	router.GET("/p/users", middlewares.RequirePermission(handlers.UserHandler.ListAllUsers, "users:list", s))

	return router
}
