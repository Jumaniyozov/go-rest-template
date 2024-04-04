package middlewares

import (
	"context"
	"github.com/Jumaniyozov/go-rest-template/internal/models"
	service "github.com/Jumaniyozov/go-rest-template/internal/services"
	"github.com/Jumaniyozov/go-rest-template/pkg/response"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

const requestTimeout = 5 * time.Second

func PermissionsInclude(code string, p []models.Permissions) bool {
	for _, val := range p {
		if code == val.Permission {
			return true
		}
	}
	return false
}

func RequirePermission(h httprouter.Handle, srv *service.Service, resp *response.Response, requiredPermission string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		user := r.Context().Value(AuthorizationPayloadKey)

		if user == nil {
			resp.PermissionDenied(w, "Unauthorized")
		}

		ctx, cancel := context.WithTimeout(r.Context(), requestTimeout)
		defer cancel()

		permissions, err := srv.Auth.AllPermissions(ctx, user.(int))

		if err != nil {
			resp.InternalServerError(w, "Error while getting permissions")
		}

		if !PermissionsInclude(requiredPermission, permissions) {
			resp.PermissionDenied(w, "You don't have required permission")
		}

		h(w, r, ps)
	}
}
