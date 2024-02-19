package rest

import (
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	service "github.com/Jumaniyozov/go-rest-template/internal/services"
	"github.com/Jumaniyozov/go-rest-template/pkg/response"
	"github.com/rs/zerolog"
)

type Handler struct {
	cfg *config.Config
	log *zerolog.Logger

	SwaggerHandler *swaggerHandler
	UserHandler    *userHandler
}

func NewHandler(
	cfg *config.Config,
	log *zerolog.Logger,
	service service.ServiceI,
) *Handler {
	resp := response.NewResponse(log)

	sh := &swaggerHandler{}
	uh := &userHandler{
		cfg:      cfg,
		logger:   log,
		service:  service,
		response: resp,
	}

	return &Handler{
		cfg:            cfg,
		log:            log,
		SwaggerHandler: sh,
		UserHandler:    uh,
	}
}
