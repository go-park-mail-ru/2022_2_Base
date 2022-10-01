package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	_ "serv/docs"

	models "serv/model"

	//jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

var SECRET_KEY = []byte("baseteam")

type Result struct {
	Body interface{} `json:"body,omitempty"`
	Err  string      `json:"err,omitempty"`
}

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type MyHandler struct {
	sessions map[string]uint
	users    map[string]*models.User
}

func NewMyHandler() *MyHandler {
	return &MyHandler{
		sessions: make(map[string]uint, 10),
		users: map[string]*models.User{
			"rvasily": {ID: 1, Username: "rvasily", Password: "love"},
		},
	}
}

type myError struct {
	Status int
	Error  string
}

// LogIn godoc
// @Summary Logs in and returns the authentication  cookie
// @Description log in user
// @ID login
// @Accept  json
// @Produce  json
// @Param user body models.UserCreateParams true "User params"
// @Success 201 {object} string
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /login [post]
func (api *MyHandler) Login(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var resp models.UserCreateParams
	err := decoder.Decode(&resp)
	if err != nil {
		http.Error(w, `Bad request`, 400)
	}

	user, ok := api.users[resp.Username]
	if !ok {
		http.Error(w, `Bad request`, 400)
		return
	}

	if user.Password != resp.Password {
		http.Error(w, `Bad request`, 400)
		return
	}

	SID := RandStringRunes(32)
	api.sessions[SID] = user.ID

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"username": resp.Username,
	// })
	// jwtToken, err := token.SignedString(SECRET_KEY)
	// log.Printf(jwtToken)
	// if err != nil {
	// 	http.Error(w, `Internal Server Error`, 500)
	// 	return
	// }

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    SID,
		Expires:  time.Now().Add(10 * time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)

	json.NewEncoder(w).Encode(cookie)
	//json.NewEncoder(w).Encode(jwtToken)
}

// LogOut godoc
// @Summary Logs out user
// @Description Logs out user
// @ID logout
// @Accept  json
// @Produce  json
// @Success 200 {object} string "OK"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Router /logout [post]
func (api *MyHandler) Logout(w http.ResponseWriter, r *http.Request) {

	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		http.Error(w, `Unauthorized - Access token is missing or invalid`, 401)
		return
	}

	if _, ok := api.sessions[session.Value]; !ok {
		http.Error(w, `Unauthorized - Access token is missing or invalid`, 401)
		return
	}

	delete(api.sessions, session.Value)

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
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
func (api *MyHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username, ok := vars["username"]
	if !ok {
		http.Error(w, `Username is missing in parameters`, 400)
		return
	}
	user, ok := api.users[username]
	if !ok {
		http.Error(w, `User not found`, 404)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (api *MyHandler) Root(w http.ResponseWriter, r *http.Request) {
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

func handleUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status": "ok"}`))
}

// @title Reazon API
// @version 1.0
// @description Reazon back server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /

//x// @BasePath /api/v1
func main() {
	myRouter := mux.NewRouter()
	api := NewMyHandler()
	myRouter.HandleFunc("/", api.Root)
	myRouter.HandleFunc("/login", api.Login)
	myRouter.HandleFunc("/logout", api.Logout)
	myRouter.HandleFunc("/getuser/{username}", api.GetUser)
	myRouter.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)
	http.ListenAndServe(":8080", myRouter)

}
