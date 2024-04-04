package routes

import (
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	"github.com/Jumaniyozov/go-rest-template/internal/logger"
	middlewares "github.com/Jumaniyozov/go-rest-template/internal/middleware"
	service "github.com/Jumaniyozov/go-rest-template/internal/services"
	"github.com/Jumaniyozov/go-rest-template/internal/transport/rest"
	"github.com/Jumaniyozov/go-rest-template/pkg/response"
	"github.com/julienschmidt/httprouter"
)

type Router struct {
	Config  *config.Config
	Logger  *logger.Logger
	Service *service.Service
}

func New(c *config.Config, l *logger.Logger, s *service.Service) *Router {
	return &Router{
		Config:  c,
		Logger:  l,
		Service: s,
	}
}

func (r *Router) CreateHttpRouter() *httprouter.Router {
	router := httprouter.New()
	handlers := rest.New(r.Config, r.Logger, r.Service)
	resp := response.New(r.Logger)

	router.GET("/swagger/:any", middlewares.RequestLogger(handlers.Swagger.Init, r.Logger))
	router.GET("/api/v1/users", middlewares.RequestLogger(handlers.User.List, r.Logger))
	router.GET("/api/v1/p/users", middlewares.RequestLogger(middlewares.RequirePermission(handlers.User.List, r.Service, resp, "users:list"), r.Logger))

	return router
}
