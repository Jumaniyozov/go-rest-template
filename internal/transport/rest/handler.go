package rest

import (
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	service "github.com/Jumaniyozov/go-rest-template/internal/services"
	"github.com/Jumaniyozov/go-rest-template/internal/transport/rest/swagger"
	"github.com/Jumaniyozov/go-rest-template/internal/transport/rest/user"
	"github.com/Jumaniyozov/go-rest-template/pkg/response"
	"github.com/rs/zerolog"
)

type Handler struct {
	Swagger *swagger.Swagger
	User    *user.User
}

func New(
	cfg *config.Config,
	log *zerolog.Logger,
	service service.ServiceI,
) *Handler {
	resp := response.New(log)

	sh := &swagger.Swagger{}
	uh := &user.User{
		Cfg:      cfg,
		Logger:   log,
		Service:  service,
		Response: resp,
	}

	return &Handler{
		Swagger: sh,
		User:    uh,
	}
}
