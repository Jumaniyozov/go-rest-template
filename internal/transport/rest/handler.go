package rest

import (
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	contractService "github.com/Jumaniyozov/go-rest-template/internal/contracts/service"
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
	service contractService.ServiceI,
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
