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
	cfg      *config.Config
	logger   *zerolog.Logger
	response *response.Response
	service  service.ServiceI
}

func (u *userHandler) ListAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := u.service.UserService().ListAllUsers()
	if err != nil {
		u.response.FetchError(w, "error while fetching users")
	}

	resp := response.Message{
		Message: "Success",
		Code:    http.StatusOK,
		Success: true,
		Data:    users,
	}

	err = u.response.WriteJSON(w, http.StatusOK, resp, nil)
	if err != nil {
		u.response.InternalErrorMessage(w, err)
	}
}
