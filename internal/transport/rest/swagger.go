package rest

import (
	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

type swaggerHandler struct{}

func (s *swaggerHandler) Init(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	httpSwagger.WrapHandler(res, req)
}
