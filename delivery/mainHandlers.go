package delivery

import (
	"encoding/json"
	"net/http"
	"serv/domain/model"
	uc "serv/usecase"
)

// @title Reozon API
// @version 1.0
// @description Reazon back server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath  /api/v1

type WebHandler struct {
	usHandler   uc.UserHandler
	prodHandler uc.ProductHandler
}

func NewWebHandler(uh *uc.UserHandler, ph *uc.ProductHandler) *WebHandler {
	return &WebHandler{
		usHandler:   *uh,
		prodHandler: *ph,
	}
}

func ReturnErrorJSON(w http.ResponseWriter, err error, errCode int) {
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(&model.Error{Error: err.Error()})
	return
}
