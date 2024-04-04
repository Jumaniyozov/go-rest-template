package user

import (
	"context"
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	service "github.com/Jumaniyozov/go-rest-template/internal/services"
	"github.com/Jumaniyozov/go-rest-template/pkg/response"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
	"net/http"
)

type User struct {
	Cfg      *config.Config
	Logger   *zerolog.Logger
	Response *response.Response
	Service  service.Service
}

func (u *User) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := u.Service.User.List(context.Background())
	if err != nil {
		u.Response.FetchError(w, "error while fetching users")
	}

	resp := response.Message{
		Message: "Success",
		Code:    http.StatusOK,
		Success: true,
		Data:    users,
	}

	err = u.Response.WriteJSON(w, http.StatusOK, resp, nil)
	if err != nil {
		u.Response.InternalErrorMessage(w, err)
	}
}
