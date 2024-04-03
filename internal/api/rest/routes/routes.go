package routes

import (
	"github.com/Jumaniyozov/go-rest-template/internal/api/rest"
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	middlewares "github.com/Jumaniyozov/go-rest-template/internal/middleware"
	service "github.com/Jumaniyozov/go-rest-template/internal/services"
	"github.com/Jumaniyozov/go-rest-template/pkg/response"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
)

type Router struct {
	Config *config.Config
	Logger *zerolog.Logger
}

func New(c *config.Config, l *zerolog.Logger, s service.ServiceI) *Router {
	return &Router{
		Config: c,
		Logger: l,
	}
}

func (r *Router) CreateHttpRouter() *httprouter.Router {
	router := httprouter.New()
	handlers := rest.New(r.Config, r.Logger, r.Service)
	resp := response.New(r.Logger)

	router.GET("/swagger/:any", middlewares.RequestLogger(handlers.Swagger.Init, r.Logger))
	router.GET("/api/v1/users", middlewares.RequestLogger(handlers.User.List, r.Logger))
	router.GET("/api/v1/p/users", middlewares.RequestLogger(middlewares.RequirePermission(handlers.User.List, resp, "users:list", r.Service), r.Logger))

	return router
}
