package main

import (
	"log"
	"net/http"

	_ "serv/docs"

	"github.com/gorilla/mux"

	deliv "serv/delivery"
	usecase "serv/usecase"

	conf "serv/config"

	httpSwagger "github.com/swaggo/http-swagger"

	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func loggingAndCORSHeadersMiddleware(next http.Handler) http.Handler {
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

	urlDB := "postgres://" + conf.DBSPuser + ":" + conf.DBPassword + "@" + conf.DBHost + ":" + conf.DBPort + "/" + conf.DBName
	db, err := sql.Open("pgx", urlDB)
	if err != nil {
		log.Println("could not connect to database")
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Println("unable to reach database ", err)
	}
	log.Println("database is reachable")

	userHandler := usecase.NewUserHandler(db)
	productHandler := usecase.NewProductHandler(db)

	webHandler := deliv.NewWebHandler(userHandler, productHandler)

	myRouter.HandleFunc(conf.PathLogin, webHandler.Login).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathLogOut, webHandler.Logout).Methods(http.MethodDelete, http.MethodOptions)
	myRouter.HandleFunc(conf.PathSignUp, webHandler.SignUp).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathSessions, webHandler.GetSession).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathMain, webHandler.GetHomePage).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathProfile, webHandler.GetUser).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathProfile, webHandler.ChangeProfile).Methods(http.MethodPost, http.MethodOptions)
	myRouter.PathPrefix(conf.PathDocs).Handler(httpSwagger.WrapHandler)
	myRouter.Use(loggingAndCORSHeadersMiddleware)
	http.ListenAndServe(conf.Port, myRouter)
}
