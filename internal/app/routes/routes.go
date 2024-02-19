package routes

import (
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	middlewares "github.com/Jumaniyozov/go-rest-template/internal/middleware"
	service "github.com/Jumaniyozov/go-rest-template/internal/services"
	"github.com/Jumaniyozov/go-rest-template/internal/transport/rest"
	"github.com/Jumaniyozov/go-rest-template/pkg/response"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
)

type Router struct {
	Config  *config.Config
	Logger  *zerolog.Logger
	Service service.ServiceI
}

func New(c *config.Config, l *zerolog.Logger, s service.ServiceI) *Router {
	return &Router{
		Config:  c,
		Logger:  l,
		Service: s,
	}
}

func (r *Router) CreateHttpRouter() *httprouter.Router {
	router := httprouter.New()
	handlers := rest.NewHandler(r.Config, r.Logger, r.Service)
	resp := response.NewResponse(r.Logger)

	router.GET("/swagger/:any", handlers.SwaggerHandler.Init)
	router.GET("/users", handlers.UserHandler.ListAllUsers)
	router.GET("/p/users", middlewares.RequirePermission(handlers.UserHandler.ListAllUsers, resp, "users:list", r.Service))

	return router
}
