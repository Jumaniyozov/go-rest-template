package rest

import (
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	service "github.com/Jumaniyozov/go-rest-template/internal/services"
	"github.com/rs/zerolog"
)

type Handler struct {
	cfg            *config.Config
	log            *zerolog.Logger
	SwaggerHandler *swaggerHandler
	UserHandler    *userHandler
}

func NewHandler(
	cfg *config.Config,
	log *zerolog.Logger,
	service service.ServiceI,
) *Handler {
	sh := &swaggerHandler{}
	uh := &userHandler{
		cfg:     cfg,
		logger:  log,
		service: service,
	}

	return &Handler{
		cfg:            cfg,
		log:            log,
		SwaggerHandler: sh,
		UserHandler:    uh,
	}
}
