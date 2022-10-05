package handlers

import (
	"encoding/json"
	"net/http"
	baseErrors "serv/errors"
	"serv/model"
	"time"

	"github.com/google/uuid"
)

func NewUserHandler() *UserHandler {
	return &UserHandler{
		sessions: make(map[string]uint),
		store:    *NewUserStore(),
	}
}
func NewProductHandler() *ProductHandler {
	return &ProductHandler{
		store: *NewProductStore(),
	}
}

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

// LogIn godoc
// @Summary Logs in and returns the authentication  cookie
// @Description Log in user
// @ID login
// @Accept  json
// @Produce  json
// @Param user body model.UserCreateParams true "UserDB params"
// @Success 201 {object} string
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /login [post, options]
func (api *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://89.208.198.137:8081")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	if r.Method == http.MethodOptions {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req model.UserCreateParams
	err := decoder.Decode(&req)
	//log.Println("0")
	if err != nil {
		http.Error(w, baseErrors.ErrBadRequest400.Error(), 400)
		return
	}
	//log.Println("1")
	user, err := api.GetUserByUsername(req.Username)
	if err != nil {
		http.Error(w, baseErrors.ErrBadRequest400.Error(), 400)
		return
	}
	//log.Println("2")
	if user.Password != req.Password {
		http.Error(w, baseErrors.ErrBadRequest400.Error(), 400)
		return
	}
	//log.Println("3")
	newUUID := uuid.New()
	api.sessions[newUUID.String()] = user.ID

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    newUUID.String(),
		Expires:  time.Now().Add(10 * time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(201)
	//w.WriteHeader("A")
	w.Header().Set("accept", "application/json")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://89.208.198.137:8081")
	json.NewEncoder(w).Encode(cookie)

}

// LogOut godoc
// @Summary Logs out user
// @Description Logs out user
// @ID logout
// @Accept  json
// @Produce  json
// @Success 200 {object} string "OK"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Router /logout [delete]
func (api *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {

	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		http.Error(w, baseErrors.ErrUnauthorized401.Error(), 401)
		return
	}

	if _, ok := api.sessions[session.Value]; !ok {
		http.Error(w, baseErrors.ErrUnauthorized401.Error(), 401)
		return
	}

	delete(api.sessions, session.Value)

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
}

// SignUp godoc
// @Summary Signs up and returns the authentication  cookie
// @Description Sign up user
// @ID signup
// @Accept  json
// @Produce  json
// @Param user body model.UserCreateParams true "UserDB params"
// @Success 201 {object} string
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 409 {object} model.Error "Conflict - UserDB already exists"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /signup [post]
func (api *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var req model.UserCreateParams
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, baseErrors.ErrBadRequest400.Error(), 400)
		return
	}

	user, err := api.GetUserByUsername(req.Username)
	if err != nil && err != baseErrors.ErrNotFound404 {
		http.Error(w, baseErrors.ErrServerError500.Error(), 500)
		return
	}

	if user.Username != "" {
		http.Error(w, baseErrors.ErrConflict409.Error(), 409)
		return
	}

	// add validation of name and pass

	_, err = api.AddUser(&req)
	if err != nil {
		http.Error(w, baseErrors.ErrServerError500.Error(), 500)
	}

	newUUID := uuid.New()
	api.sessions[newUUID.String()] = user.ID

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    newUUID.String(),
		Expires:  time.Now().Add(10 * time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(cookie)
}

// GetSession godoc
// @Summary Checks if user has active session
// @Description Checks if user has active session
// @ID session
// @Accept  json
// @Produce  json
// @Success 200 {object} string "OK"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Router /session [get]
func (api *UserHandler) GetSession(w http.ResponseWriter, r *http.Request) {

	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		http.Error(w, baseErrors.ErrUnauthorized401.Error(), 401)
		return
	}
	if _, ok := api.sessions[session.Value]; !ok {
		http.Error(w, baseErrors.ErrUnauthorized401.Error(), 401)
		return
	}
	json.NewEncoder(w).Encode(r.Cookies()[0])
}

type ProductCollection struct {
	Body interface{} `json:"body,omitempty"`
}

// GetHomePage godoc
// @Summary Gets products for main page
// @Description Gets products for main page
// @ID getMain
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Product
// @Failure 404 {object} model.Error "Products not found"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router / [get]
func (api *ProductHandler) GetHomePage(w http.ResponseWriter, r *http.Request) {

	products, err := api.GetProducts()
	if err != nil {
		http.Error(w, baseErrors.ErrServerError500.Error(), 500)
		return
	}
	if len(products) == 0 {
		http.Error(w, baseErrors.ErrNotFound404.Error(), 404)
		return
	}

	json.NewEncoder(w).Encode(&ProductCollection{Body: products})
}
