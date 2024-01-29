package middlewares

import (
	contractService "github.com/Jumaniyozov/go-rest-template/internal/contracts/service"
	"github.com/Jumaniyozov/go-rest-template/internal/models"
	"github.com/Jumaniyozov/go-rest-template/pkg/response"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func PermissionsInclude(code string, p []models.Permissions) bool {

	for _, val := range p {
		if code == val.Permission {
			return true
		}
	}
	return false
}

func RequirePermission(h httprouter.Handle, resp *response.Response, requiredPermission string, srv contractService.ServiceI) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		user := r.Context().Value(AuthorizationPayloadKey)

		permissions, err := srv.AuthService().GetAllPermissions(user.(int))

		if err != nil {
			resp.InternalServerError(w, "Error while getting permissions")
		}

		if !PermissionsInclude(requiredPermission, permissions) {
			resp.PermissionDenied(w, "You don't have required permission")
		}

		h(w, r, ps)
	}
}
