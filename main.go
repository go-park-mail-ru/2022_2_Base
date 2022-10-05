package main

import (
	"net/http"

	_ "serv/docs"

	"github.com/gorilla/mux"

	handlers "serv/handlers"

	conf "serv/config"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	myRouter := mux.NewRouter()
	userHandler := handlers.NewUserHandler()
	productHandler := handlers.NewProductHandler()

	myRouter.HandleFunc(conf.PathLogin, userHandler.Login).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathLogOut, userHandler.Logout).Methods("Delete")
	myRouter.HandleFunc(conf.PathSignUp, userHandler.SignUp).Methods("Post")
	myRouter.HandleFunc(conf.PathSessions, userHandler.GetSession).Methods("GET")
	myRouter.HandleFunc(conf.PathMain, productHandler.GetHomePage).Methods("GET")
	myRouter.PathPrefix(conf.PathDocs).Handler(httpSwagger.WrapHandler)
	myRouter.Use(mux.CORSMethodMiddleware(myRouter))
	http.ListenAndServe(conf.Port, myRouter)
}
