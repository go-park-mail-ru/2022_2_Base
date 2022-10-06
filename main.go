package main

import (
	"log"
	"net/http"

	_ "serv/docs"

	"github.com/gorilla/mux"

	handlers "serv/handlers"

	conf "serv/config"

	httpSwagger "github.com/swaggo/http-swagger"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		for header := range conf.Headers {
			w.Header().Set(header, conf.Headers[header])
		}
		next.ServeHTTP(w, r)
	})
}

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
	myRouter.Use(loggingMiddleware)
	http.ListenAndServe(conf.Port, myRouter)
}
