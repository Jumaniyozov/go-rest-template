package rest

import (
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	service "github.com/Jumaniyozov/go-rest-template/internal/services"
	"github.com/Jumaniyozov/go-rest-template/pkg/response"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
	"net/http"
)

type userHandler struct {
	cfg     *config.Config
	logger  *zerolog.Logger
	service service.ServiceI
}

func (u *userHandler) ListAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := u.service.UserService().ListAllUsers()
	if err != nil {
		u.logger.Error().Err(err).Msg("Error while fetching users")
		resp := response.Message{
			Message: "Error while fetching users",
			Code:    http.StatusInternalServerError,
			Data:    nil,
		}

		err := response.WriteJSON(w, http.StatusInternalServerError, resp, nil)
		if err != nil {
			response.InternalErrorMessage(w, err, u.logger)
		}
	}

	resp := response.Message{
		Message: "Success",
		Code:    http.StatusOK,
		Data:    users,
	}

	err = response.WriteJSON(w, http.StatusOK, resp, nil)
	if err != nil {
		response.InternalErrorMessage(w, err, u.logger)
	}
}
