package delivery

import (
	"encoding/json"
	"net/http"
	"regexp"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	"time"

	"github.com/google/uuid"

	usecase "serv/usecase"
)

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
func (api *WebHandler) Login(w http.ResponseWriter, r *http.Request) {
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
	user, err := api.usHandler.GetUserByUsername(req.Email)
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
	usecase.SetSession(&api.usHandler, newUUID.String(), user.Email)

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
func (api *WebHandler) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	res, err := usecase.GetSession(&api.usHandler, session.Value)
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	usecase.DeleteSession(&api.usHandler, res)

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
func (api *WebHandler) SignUp(w http.ResponseWriter, r *http.Request) {
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

	user, err := api.usHandler.GetUserByUsername(req.Email)
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

	_, err = api.usHandler.AddUser(&req)
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	newUUID := uuid.New()
	usecase.SetSession(&api.usHandler, newUUID.String(), user.Email)

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
func (api *WebHandler) GetSession(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	_, err = usecase.GetSession(&api.usHandler, session.Value)
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}
	http.SetCookie(w, r.Cookies()[0])
	json.NewEncoder(w).Encode(&model.Response{})
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
func (api *WebHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}
	usName, err := usecase.GetSession(&api.usHandler, session.Value)
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	user, err := api.usHandler.GetUserByUsername(usName)
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	if user.Email == "" {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	userProfile := model.UserProfile{Email: user.Email, Username: user.Username, Phone: "", Avatar: ""}
	json.NewEncoder(w).Encode(&model.Response{Body: userProfile})
}

// ChangeUser godoc
// @Summary changes user parameters
// @Description changes user parameters
// @ID changeUserParameters
// @Accept  json
// @Produce  json
// @Param user body model.UserProfile true "UserDB params"
// @Success 200 {object} model.Response "OK"
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /profile [post]
func (api *WebHandler) ChangeProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	_, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req model.UserProfile
	err = decoder.Decode(&req)
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	oldUserData, err := api.usHandler.GetUserByUsername(req.Email)
	if oldUserData.Email == "" {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	count, err := api.usHandler.ChangeUser(oldUserData.Email, req)
	if err != nil && count == 0 {
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	json.NewEncoder(w).Encode(&model.Response{})
}
