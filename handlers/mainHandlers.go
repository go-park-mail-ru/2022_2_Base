package handlers

import (
	"encoding/json"
	"net/http"
	baseErrors "serv/errors"
	"serv/model"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func NewUserHandler() (*UserHandler, error) {
	return &UserHandler{
		sessions: make(map[string]uint, 10),
		store:    *NewUserStore(),
	}, nil
}

func (api *UserHandler) Root(w http.ResponseWriter, r *http.Request) {
	authorized := false
	session, err := r.Cookie("session_id")
	if err == nil && session != nil {
		_, authorized = api.sessions[session.Value]
	}

	if authorized {
		w.Write([]byte("autrorized"))
	} else {
		w.Write([]byte("not autrorized"))
	}
}

// GetUser godoc
// @Summary Get current user
// @Description gets user by username
// @ID getUser
// @Accept  json
// @Produce  json
// @Param username path string true "Username"
// @Success 200 {object} model.User
// @Failure 400 {object} model.Error "Bad request"
// @Failure 404 {object} model.Error "User not found"
// @Router /getuser/{username} [get]
func (api *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username, ok := vars["username"]
	if !ok {
		http.Error(w, baseErrors.ErrBadRequest400.Error(), 400)
		return
	}
	user, err := api.UserByUsername(username)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	//users, err := api.store.GetUsers()
	//log.Println("123", users)
	//user = model.User{ID: 4, Username: "22", Password: "33"}
	//ss, err := api.store.AddUser(&user)
	//log.Println("ss", ss)
	json.NewEncoder(w).Encode(user)
}

// LogIn godoc
// @Summary Logs in and returns the authentication  cookie
// @Description Log in user
// @ID login
// @Accept  json
// @Produce  json
// @Param user body model.UserCreateParams true "User params"
// @Success 201 {object} string
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /login [post]
func (api *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var req model.UserCreateParams
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, baseErrors.ErrBadRequest400.Error(), 400)
	}

	user, err := api.UserByUsername(req.Username)
	if err != nil {
		http.Error(w, baseErrors.ErrBadRequest400.Error(), 400)
		return
	}

	if user.Password != req.Password {
		http.Error(w, baseErrors.ErrBadRequest400.Error(), 400)
		return
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
// @Param user body model.UserCreateParams true "User params"
// @Success 201 {object} string
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 409 {object} model.Error "Conflict - User already exists"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /signup [post]
func (api *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var req model.UserCreateParams
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, baseErrors.ErrBadRequest400.Error(), 400)
	}

	user, err := api.UserByUsername(req.Username)
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
