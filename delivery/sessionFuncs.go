package delivery

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	baseErrors "serv/domain/errors"
	"serv/domain/model"
	"time"

	usecase "serv/usecase"

	"golang.org/x/crypto/argon2"
)

type SessionHandler struct {
	usecase usecase.UserUsecaseInterface
}

func NewSessionHandler(uuc usecase.UserUsecaseInterface) *SessionHandler {
	return &SessionHandler{
		usecase: uuc,
	}
}

func hashPass(salt []byte, plainPassword string) []byte {
	hashedPass := argon2.IDKey([]byte(plainPassword), []byte(salt), 1, 64*1024, 4, 32)
	return append(salt, hashedPass...)
}

func checkPass(passHash []byte, plainPassword string) bool {
	//salt := passHash[0:8]
	salt := []byte("Base2022")
	userPassHash := hashPass(salt, plainPassword)
	return bytes.Equal(userPassHash, passHash)
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
		log.Println("get UserLogin ", err)
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}
	user, err := api.usecase.GetUserByUsername(req.Email)
	if err != nil {
		log.Println("get GetUserByUsername ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	if user.Email == "" {
		log.Println("get Email ", err)
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	byteUserPass, err := base64.RawStdEncoding.DecodeString(user.Password)
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	if !checkPass(byteUserPass, req.Password) {
		log.Println("get Password ", err)
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	sess, err := api.usecase.SetSession(user.Email)
	if err != nil {
		log.Println("error with auth microservice: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sess.ID,
		Expires:  time.Now().Add(10 * time.Hour),
		HttpOnly: true,
		Secure:   true,
	}

	curSession := model.Session{ID: 0, UserUUID: sess.ID}
	hashTok := HashToken{Secret: []byte("Base")}
	token, err := hashTok.CreateCSRFToken(&curSession, time.Now().Add(10*time.Hour).Unix())

	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	w.Header().Set("csrf", token)

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
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}
	_, err = api.usecase.CheckSession(session.Value)
	if err != nil {
		log.Println("no sess ", err)
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}
	err = api.usecase.DeleteSession(session.Value)
	if err != nil {
		log.Println("error with auth microservice: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
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
		log.Println(err)
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
		log.Println("error user exists ", err)
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	//validation
	match, _ := regexp.MatchString(`^(.+)@(.+)$`, req.Email)
	if !match {
		log.Println("validation error ", err)
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	if len(req.Password) < 6 {
		log.Println("validation error ", err)
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	salt := []byte("Base2022")
	hashedPass := hashPass(salt, req.Password)
	b64Pass := base64.RawStdEncoding.EncodeToString(hashedPass)
	req.Password = b64Pass

	id, err := api.usecase.AddUser(&req)
	if err != nil {
		log.Println("error while adding user to db: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	RegisterMail := model.Mail{Type: "greeting", Username: req.Username, Useremail: req.Email}
	go func() {
		err = api.usecase.SendMail(RegisterMail)
		if err != nil {
			log.Println("error sending greeting email ", err)
			//ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
			//return
		}
	}()

	promo := api.usecase.GenPromocode(id)
	PromoMail := model.Mail{Type: "promocode", Username: req.Username, Useremail: req.Email, Promocode: promo}
	go func() {
		err = api.usecase.SendMail(PromoMail)
		if err != nil {
			log.Println("error sending promocode email ", err)
			//ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
			//return
		}
	}()

	sess, err := api.usecase.SetSession(req.Email)

	if err != nil {
		log.Println("error with auth microservice: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sess.ID,
		Expires:  time.Now().Add(10 * time.Hour),
		HttpOnly: true,
		Secure:   true,
	}

	curSession := model.Session{ID: 0, UserUUID: sess.ID}
	hashTok := HashToken{Secret: []byte("Base")}
	token, err := hashTok.CreateCSRFToken(&curSession, time.Now().Add(10*time.Hour).Unix())
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	w.Header().Set("csrf", token)

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
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	_, err = api.usecase.CheckSession(session.Value)
	if err != nil {
		log.Println("no sess ", err)
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		return
	}

	curSession := model.Session{ID: 0, UserUUID: session.Value}
	hashTok := HashToken{Secret: []byte("Base")}
	token, err := hashTok.CreateCSRFToken(&curSession, time.Now().Add(24*time.Hour).Unix())
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	w.Header().Set("csrf", token)

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    r.Cookies()[0].Value,
		Expires:  r.Cookies()[0].Expires,
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, cookie)
	json.NewEncoder(w).Encode(&model.Response{})
}
