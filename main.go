package main

import (
	"context"
	"log"
	"net/http"
	"os"

	_ "serv/docs"
	"serv/repository"

	"github.com/gorilla/mux"

	deliv "serv/delivery"
	usecase "serv/usecase"

	conf "serv/config"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/jackc/pgx/v5/pgxpool"
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

//172.23.0.2:5432 local db
//172.28.0.2:5431
func main() {
	myRouter := mux.NewRouter()
	//urlDB := "postgres://" + conf.DBSPuser + ":" + conf.DBPassword + "@" + conf.DBHost + ":" + conf.DBPort + "/" + conf.DBName
	urlDB := "postgres://" + os.Getenv("TEST_POSTGRES_USER") + ":" + os.Getenv("TEST_POSTGRES_PASSWORD") + "@" + os.Getenv("TEST_DATABASE_HOST") + ":" + os.Getenv("TEST_DATABASE_PORT") + "/" + os.Getenv("TEST_POSTGRES_DB")
	log.Println("conn: ", urlDB)
	//db, err := sql.Open("pgx", urlDB)
	db, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println("could not connect to database")
	} else {
		log.Println("database is reachable")
	}
	defer db.Close()

	// if err := db.Ping(); err != nil {
	// 	log.Println("unable to reach database ", err)
	// } else {
	// 	log.Println("database is reachable")
	// }

	userStore := repository.NewUserStore(db)
	productStore := repository.NewProductStore(db)

	userUsecase := usecase.NewUserUsecase(userStore)
	productUsecase := usecase.NewProductUsecase(productStore)

	userHandler := deliv.NewUserHandler(userUsecase)
	productHandler := deliv.NewProductHandler(productUsecase)

	myRouter.HandleFunc(conf.PathLogin, userHandler.Login).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathLogOut, userHandler.Logout).Methods(http.MethodDelete, http.MethodOptions)
	myRouter.HandleFunc(conf.PathSignUp, userHandler.SignUp).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathSessions, userHandler.GetSession).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathMain, productHandler.GetHomePage).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathProfile, userHandler.GetUser).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathProfile, userHandler.ChangeProfile).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathAvatar, userHandler.SetAvatar).Methods(http.MethodPost, http.MethodOptions)
	myRouter.PathPrefix(conf.PathDocs).Handler(httpSwagger.WrapHandler)
	myRouter.Use(loggingAndCORSHeadersMiddleware)
	http.ListenAndServe(conf.Port, myRouter)
}
