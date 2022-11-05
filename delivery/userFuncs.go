package delivery

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	"strconv"
	"time"

	"github.com/google/uuid"

	usecase "serv/usecase"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(uuc *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		usecase: *uuc,
	}
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

	_, err = api.usecase.AddUser(&req)
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

	_, err = api.usecase.GetSession(session.Value)
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
func (api *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		log.Println("no session")
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}
	usName, err := api.usecase.GetSession(session.Value)
	if err != nil {
		log.Println("no session2")
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	user, err := api.usecase.GetUserByUsername(usName)
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	if user.Email == "" {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	//userProfile := model.UserProfile{Email: user.Email, Username: user.Username, Phone: *user.Phone, Avatar: *user.Avatar}
	userProfile := model.UserProfile{Email: user.Email, Username: user.Username}
	if user.Phone != nil {
		userProfile.Phone = *user.Phone
	} else {
		userProfile.Phone = ""
	}
	if user.Avatar != nil {
		userProfile.Avatar = *user.Avatar
	} else {
		userProfile.Avatar = ""
	}
	json.NewEncoder(w).Encode(userProfile)
}

// ChangeUser godoc
// @Summary changes user parameters
// @Description changes user parameters
// @ID changeUserParameters
// @Accept  json
// @Produce  json
// @Param user body model.UserProfile true "UserProfile params"
// @Success 200 {object} model.Response "OK"
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /profile [post]
func (api *UserHandler) ChangeProfile(w http.ResponseWriter, r *http.Request) {
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

	oldUserData, err := api.usecase.GetUserByUsername(req.Email)
	if err != nil {
		log.Println("db error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	if oldUserData.Email == "" {
		log.Println("error user not found")
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	ansRowCount, err := api.usecase.ChangeUser(oldUserData.Email, req)
	if err != nil && ansRowCount == 0 {
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	json.NewEncoder(w).Encode(&model.Response{})
}

// SetAvatar godoc
// @Summary Set user's avatar
// @Description sets user's avatar
// @ID setAvatar
// @Accept  multipart/form-data
// @Produce  json
// @Param file formData file true "user's avatar"
// @Success 200 {object} model.Response "OK"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /avatar [post]
func (api *UserHandler) SetAvatar(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}
	usName, err := api.usecase.GetSession(session.Value)
	if err != nil {
		log.Println("no session")
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	user, err := api.usecase.GetUserByUsername(usName)
	if err != nil {
		log.Println("db error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	if user.Email == "" {
		log.Println("error user not found")
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	userDB := model.UserDB{ID: user.ID, Email: user.Email, Username: user.Username, Password: user.Password}

	//r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Println("error parse file")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	defer file.Close()
	fileName := "./img/avatars/avatar" + strconv.FormatUint(uint64(userDB.ID), 10) + ".jpg"
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println("error create/open file")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		log.Println("error copy to new file")
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	newUserData := model.UserProfile{Email: userDB.Email, Username: userDB.Username, Avatar: fileName[1:]}
	if userDB.Phone != nil {
		newUserData.Phone = *userDB.Phone
	} else {
		newUserData.Phone = ""
	}

	ansRowCount, err := api.usecase.ChangeUser(userDB.Email, newUserData)
	if err != nil && ansRowCount == 0 {
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	json.NewEncoder(w).Encode(&model.Response{})
}
