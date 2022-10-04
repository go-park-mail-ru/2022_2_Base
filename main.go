package main

import (
	"net/http"

	_ "serv/docs"

	"github.com/gorilla/mux"

	handlers "serv/handlers"

	conf "serv/config"

	httpSwagger "github.com/swaggo/http-swagger"
)

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
	userHandler := handlers.NewUserHandler()

	productHandler := handlers.NewProductHandler()

	//myRouter.HandleFunc("/", userHandler.Root).Methods("GET")
	myRouter.HandleFunc(conf.PathLogin, userHandler.Login).Methods("Post")
	myRouter.HandleFunc(conf.PathLogOut, userHandler.Logout).Methods("Delete")
	myRouter.HandleFunc(conf.PathSignUp, userHandler.SignUp).Methods("Post")
	myRouter.HandleFunc(conf.PathGetUser, userHandler.GetUser).Methods("GET")
	myRouter.HandleFunc(conf.PathSessions, userHandler.GetSession).Methods("GET")
	myRouter.HandleFunc(conf.PathMain, productHandler.GetHomePage).Methods("GET")
	myRouter.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)
	http.ListenAndServe(conf.Port, myRouter)

}
