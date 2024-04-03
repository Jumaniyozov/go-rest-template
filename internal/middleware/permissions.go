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

func PermissionsInclude(code string, p []models.Permissions) bool {

	for _, val := range p {
		if code == val.Permission {
			return true
		}
	}
	return false
}

func RequirePermission(h httprouter.Handle, resp *response.Response, requiredPermission string, srv service.ServiceI) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		user := r.Context().Value(AuthorizationPayloadKey)

		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		permissions, err := srv.AuthService().AllPermissions(ctx, user.(int))

		if err != nil {
			resp.InternalServerError(w, "Error while getting permissions")
		}

		if !PermissionsInclude(requiredPermission, permissions) {
			resp.PermissionDenied(w, "You don't have required permission")
		}

		h(w, r, ps)
	}
}
