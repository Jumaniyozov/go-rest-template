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

func RequirePermission(h httprouter.Handle, requiredPermission string, srv contractService.ServiceI) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		user := r.Context().Value(AuthorizationPayloadKey)

		permissions, err := srv.AuthService().GetAllPermissions(user.(int))
		//permissions, err := []string{"admin", "user"}
		if err != nil {
			rsp := response.Message{
				Message: "Error while getting permissions",
				Code:    http.StatusInternalServerError,
				Data:    nil,
			}
			err := response.WriteJSON(w, http.StatusInternalServerError, rsp, nil)
			if err != nil {
				response.InternalErrorMessage(w, err, nil)
			}
		}
		// Get the slice of permissions for the user.
		//permissions, err := app.models.Permissions.GetAllForUser(user.ID)
		//if err != nil {
		//	app.serverErrorResponse(w, r, err)
		//	return
		//}
		// Check if the slice includes the required permission. If it doesn't, then
		// return a 403 Forbidden response.
		if !PermissionsInclude(requiredPermission, permissions) {
			rsp := response.Message{
				Message: "You don't have required permission",
				Code:    http.StatusBadRequest,
				Data:    nil,
			}
			err := response.WriteJSON(w, http.StatusBadRequest, rsp, http.Header{
				"WWW-Authenticate": []string{"Basic realm=Restricted"},
			})
			if err != nil {
				response.InternalErrorMessage(w, err, nil)
			}
		}

		h(w, r, ps)
	}
}
