package delivery

import (
	"encoding/json"
	"log"
	"net/http"
	baseErrors "serv/domain/errors"
	"serv/domain/model"

	"github.com/microcosm-cc/bluemonday"
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

type OrderHandler struct {
	usHandler UserHandler
	prHandler ProductHandler
}

func NewOrderHandler(us *UserHandler, pr *ProductHandler) *OrderHandler {
	return &OrderHandler{
		usHandler: *us,
		prHandler: *pr,
	}
}

func ReturnErrorJSON(w http.ResponseWriter, err error, errCode int) {
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(&model.Error{Error: err.Error()})
	return
}

// GetCart godoc
// @Summary gets user's cart
// @Description gets user's cart
// @ID GetCart
// @Accept  json
// @Produce  json
// @Tags Order
// @Success 200 {object} model.Order
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /cart [get]
func (api *OrderHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	sanitizer := bluemonday.UGCPolicy()
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}
	usName, err := api.usHandler.usecase.GetSession(session.Value)
	UserData, err := api.usHandler.usecase.GetUserByUsername(usName)
	if err != nil {
		log.Println("db error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	if UserData.Email == "" {
		log.Println("error user not found")
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	cart, err := api.prHandler.usecase.GetCart(UserData.ID)
	if err != nil {
		log.Println("db error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	cart.OrderStatus = sanitizer.Sanitize(cart.OrderStatus)
	cart.PaymentStatus = sanitizer.Sanitize(cart.PaymentStatus)
	cart.Adress = sanitizer.Sanitize(cart.Adress)
	for _, prod := range cart.Items {
		prod.Name = sanitizer.Sanitize(prod.Name)
		prod.Description = sanitizer.Sanitize(prod.Description)
		prod.Imgsrc = sanitizer.Sanitize(prod.Imgsrc)
	}

	json.NewEncoder(w).Encode(cart)
}

// UpdateCart godoc
// @Summary updates user's cart
// @Description updates user's cart
// @ID UpdateCart
// @Accept  json
// @Produce  json
// @Tags Order
// @Param items body model.ProductCart true "ProductCart items"
// @Success 200 {object} model.Product
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /cart [post]
func (api *OrderHandler) UpdateCart(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req model.ProductCart
	err := decoder.Decode(&req)
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}
	usName, err := api.usHandler.usecase.GetSession(session.Value)
	UserData, err := api.usHandler.usecase.GetUserByUsername(usName)
	if err != nil {
		log.Println("db error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	if UserData.Email == "" {
		log.Println("error user not found")
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	err = api.prHandler.usecase.UpdateOrder(UserData.ID, &req.Items)
	if err != nil {
		log.Println("db error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	json.NewEncoder(w).Encode(&model.Response{})
}

// MakeOrder godoc
// @Summary makes user's order
// @Description makes user's order
// @ID MakeOrder
// @Accept  json
// @Produce  json
// @Tags Order
// @Success 200 {object} model.Response "OK"
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /makeorder [post]
func (api *OrderHandler) MakeOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	usName, err := api.usHandler.usecase.GetSession(session.Value)
	UserData, err := api.usHandler.usecase.GetUserByUsername(usName)
	if err != nil {
		log.Println("db error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	if UserData.Email == "" {
		log.Println("error user not found")
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	err = api.prHandler.usecase.MakeOrder(UserData.ID)
	if err != nil {
		log.Println("db error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	json.NewEncoder(w).Encode(&model.Response{})
}
