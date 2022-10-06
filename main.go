package main

import (
	"net/http"

	_ "serv/docs"

	"github.com/gorilla/mux"

	handlers "serv/handlers"

	conf "serv/config"

	gorHandlers "github.com/gorilla/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

var headersOk = gorHandlers.AllowedHeaders([]string{"X-Requested-With"})
var originsOk = gorHandlers.AllowedOrigins([]string{"http://89.208.198.137:8081", "http://127.0.0.1:8080"})
var methodsOk = gorHandlers.AllowedMethods([]string{"DELETE", "GET", "HEAD", "POST", "PUT", "OPTIONS"})
var credsOk = gorHandlers.AllowCredentials()

func main() {
	myRouter := mux.NewRouter()
	userHandler := handlers.NewUserHandler()
	productHandler := handlers.NewProductHandler()

	myRouter.HandleFunc(conf.PathLogin, userHandler.Login).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathLogOut, userHandler.Logout).Methods(http.MethodDelete, http.MethodOptions)
	myRouter.HandleFunc(conf.PathSignUp, userHandler.SignUp).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathSessions, userHandler.GetSession).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathMain, productHandler.GetHomePage).Methods(http.MethodGet, http.MethodOptions)
	myRouter.PathPrefix(conf.PathDocs).Handler(httpSwagger.WrapHandler)
	myRouter.Use(mux.CORSMethodMiddleware(myRouter))
	http.ListenAndServe(conf.Port, gorHandlers.CORS(originsOk, headersOk, methodsOk, credsOk)(myRouter))
}
