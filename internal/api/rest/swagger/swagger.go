package swagger

import (
	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

type Swagger struct{}

func (s *Swagger) Init(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	httpSwagger.WrapHandler(res, req)
}
