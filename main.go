package main

import (
	"log"
	"net/http"

	_ "serv/docs"

	"github.com/gorilla/mux"

	handlers "serv/handlers"

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
	api, err := handlers.NewUserHandler()
	if err != nil {
		log.Println("1234")
	}
	myRouter.HandleFunc("/", api.Root).Methods("GET")
	myRouter.HandleFunc("/login", api.Login).Methods("Post")
	myRouter.HandleFunc("/logout", api.Logout).Methods("Delete")
	myRouter.HandleFunc("/signup", api.SignUp).Methods("Post")
	myRouter.HandleFunc("/getuser/{username}", api.GetUser).Methods("GET")
	myRouter.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)
	http.ListenAndServe(":8080", myRouter)

}
