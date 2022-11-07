package delivery

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	"time"

	"github.com/google/uuid"

	usecase "serv/usecase"
)

type SessionHandler struct {
	usecase usecase.UserUsecase
}

func NewSessionHandler(uuc *usecase.UserUsecase) *SessionHandler {
	return &SessionHandler{
		usecase: *uuc,
	}
}

// LogIn godoc
// @Summary Logs in and returns the authentication  cookie
// @Description Log in user
// @ID login
// @Accept  json
// @Produce  json
// @Tags User
// @Param user body model.UserLogin true "UserDB params"
// @Success 201 {object} model.Response "OK"
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /login [post]
func (api *SessionHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	decoder := json.NewDecoder(r.Body)
	var req model.UserLogin
	err := decoder.Decode(&req)
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}
	user, err := api.usecase.GetUserByUsername(req.Email)
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	if user.Email == "" {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}
	if user.Password != req.Password {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	newUUID := uuid.New()
	api.usecase.SetSession(newUUID.String(), user.Email)

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
// @Tags User
// @Success 200 {object} model.Response "OK"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Router /logout [delete]
func (api *SessionHandler) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	res, err := api.usecase.GetSession(session.Value)
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	api.usecase.DeleteSession(res)

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
// @Tags User
// @Param user body model.UserCreateParams true "UserDB params"
// @Success 201 {object} model.Response "OK"
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 409 {object} model.Error "Conflict - UserDB already exists"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /signup [post]
func (api *SessionHandler) SignUp(w http.ResponseWriter, r *http.Request) {
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

	user, err := api.usecase.GetUserByUsername(req.Email)
	if err != nil && err != baseErrors.ErrNotFound404 {
		log.Println("error ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	if user.Email != "" {
		log.Println("error user exists")
		ReturnErrorJSON(w, baseErrors.ErrConflict409, 409)
		return
	}

	//validation
	match, _ := regexp.MatchString(`^(.+)@(.+)$`, req.Email)
	if !match {
		log.Println("validation error")
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	if len(req.Password) < 6 {
		log.Println("validation error")
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	err = api.usecase.AddUser(&req)
	if err != nil {
		log.Println("error while adding user to db: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	newUUID := uuid.New()

	api.usecase.SetSession(newUUID.String(), req.Email)

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
// @Tags User
// @Success 200 {object} model.Response "OK"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Router /session [get]
func (api *SessionHandler) GetSession(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	_, err = api.usecase.GetSession(session.Value)
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}
	http.SetCookie(w, r.Cookies()[0])
	json.NewEncoder(w).Encode(&model.Response{})
}