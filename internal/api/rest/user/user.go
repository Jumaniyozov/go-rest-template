package user

import (
	"context"
	service "github.com/Jumaniyozov/go-rest-template/internal/services"
	"github.com/Jumaniyozov/go-rest-template/pkg/response"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

type UserImplementation struct {
	service service.User
}

func NewImplementation(uS service.User) *UserImplementation {
	return &Implementation{
		noteService: noteService,
	}
}

const requestTimeout = 5 * time.Second

func (u *UserImplementation) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	users, err := u.service.List(ctx)
	if err != nil {
		u.logger.Error().Err(err).Msg("error while fetching users")
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
		u.logger.Error().Err(err).Msg("error while fetching users")
		u.response.InternalErrorMessage(w, err)
	}
}
