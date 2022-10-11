package handlers

import (
	"encoding/json"
	"net/http"
	"regexp"
	baseErrors "serv/errors"
	"serv/model"
	"time"

	"github.com/google/uuid"
)

func NewUserHandler() *UserHandler {
	return &UserHandler{
		sessions: make(map[string]string),
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

func ReturnErrorJSON(w http.ResponseWriter, err error, errCode int) {
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(&model.Error{Error: err.Error()})
	return
}

// LogIn godoc
// @Summary Logs in and returns the authentication  cookie
// @Description Log in user
// @ID login
// @Accept  json
// @Produce  json
// @Param user body model.UserCreateParams true "UserDB params"
// @Success 201 {object} model.Response "OK"
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /login [post]
func (api *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	decoder := json.NewDecoder(r.Body)
	var req model.UserCreateParams
	err := decoder.Decode(&req)
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}
	user, err := api.GetUserByUsername(req.Email)
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}
	if user.Password != req.Password {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}
	newUUID := uuid.New()
	api.sessions[newUUID.String()] = user.Email

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    newUUID.String(),
		Expires:  time.Now().Add(10 * time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(&model.Response{})
}

// LogOut godoc
// @Summary Logs out user
// @Description Logs out user
// @ID logout
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Response "OK"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Router /logout [delete]
func (api *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	if _, ok := api.sessions[session.Value]; !ok {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	delete(api.sessions, session.Value)

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
	json.NewEncoder(w).Encode(&model.Response{})
}

// SignUp godoc
// @Summary Signs up and returns the authentication  cookie
// @Description Sign up user
// @ID signup
// @Accept  json
// @Produce  json
// @Param user body model.UserCreateParams true "UserDB params"
// @Success 201 {object} model.Response "OK"
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 409 {object} model.Error "Conflict - UserDB already exists"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /signup [post]
func (api *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	decoder := json.NewDecoder(r.Body)
	var req model.UserCreateParams
	err := decoder.Decode(&req)
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	user, err := api.GetUserByUsername(req.Email)
	if err != nil && err != baseErrors.ErrNotFound404 {
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	if user.Email != "" {
		ReturnErrorJSON(w, baseErrors.ErrConflict409, 409)
		return
	}

	//validation
	match, _ := regexp.MatchString(`^(.+)@(.+)$`, req.Email)
	if !match {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	if len(req.Password) < 6 {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	_, err = api.AddUser(&req)
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	newUUID := uuid.New()
	api.sessions[newUUID.String()] = user.Email

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    newUUID.String(),
		Expires:  time.Now().Add(10 * time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(&model.Response{})
}

// GetSession godoc
// @Summary Checks if user has active session
// @Description Checks if user has active session
// @ID session
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Response "OK"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Router /session [get]
func (api *UserHandler) GetSession(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}
	if _, ok := api.sessions[session.Value]; !ok {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}
	http.SetCookie(w, r.Cookies()[0])
	json.NewEncoder(w).Encode(&model.Response{})
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
// @Router /products [get]
func (api *ProductHandler) GetHomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	products, err := api.GetProducts()
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	if len(products) == 0 {
		ReturnErrorJSON(w, baseErrors.ErrNotFound404, 404)
		return
	}

	json.NewEncoder(w).Encode(&model.Response{Body: products})
}

// GetUser godoc
// @Summary Get current user
// @Description gets user by username in cookies
// @ID getUser
// @Accept  json
// @Produce  json
// @Success 200 {object} model.UserProfile
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /profile [get]
func (api *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	user, err := api.GetUserByUsername(api.sessions[session.Value])
	if err != nil && err != baseErrors.ErrNotFound404 {
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	userProfile := model.UserProfile{Email: user.Email, Username: user.Username, Phone: "111", Avatar: ""}
	json.NewEncoder(w).Encode(&model.Response{Body: userProfile})
}
