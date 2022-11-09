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
		log.Println(r.RequestURI, r.Method)
		for header := range conf.Headers {
			w.Header().Set(header, conf.Headers[header])
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	myRouter := mux.NewRouter()
	//urlDB := "postgres://" + conf.DBSPuser + ":" + conf.DBPassword + "@" + conf.DBHost + ":" + conf.DBPort + "/" + conf.DBName
	urlDB := "postgres://" + os.Getenv("TEST_POSTGRES_USER") + ":" + os.Getenv("TEST_POSTGRES_PASSWORD") + "@" + os.Getenv("TEST_DATABASE_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("TEST_POSTGRES_DB")
	log.Println("conn: ", urlDB)
	db, err := pgxpool.New(context.Background(), urlDB)
	if err != nil {
		log.Println("could not connect to database")
	} else {
		log.Println("database is reachable")
	}
	defer db.Close()

	userStore := repository.NewUserStore(db)
	productStore := repository.NewProductStore(db)

	userUsecase := usecase.NewUserUsecase(userStore)
	productUsecase := usecase.NewProductUsecase(productStore)

	userHandler := deliv.NewUserHandler(userUsecase)
	sessionHandler := deliv.NewSessionHandler(userUsecase)
	productHandler := deliv.NewProductHandler(productUsecase)

	orderHandler := deliv.NewOrderHandler(userHandler, productHandler)

	myRouter.HandleFunc(conf.PathLogin, sessionHandler.Login).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathLogOut, sessionHandler.Logout).Methods(http.MethodDelete, http.MethodOptions)
	myRouter.HandleFunc(conf.PathSignUp, sessionHandler.SignUp).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathSessions, sessionHandler.GetSession).Methods(http.MethodGet, http.MethodOptions)

	myRouter.HandleFunc(conf.PathProfile, userHandler.GetUser).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathProfile, userHandler.ChangeProfile).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathAvatar, userHandler.SetAvatar).Methods(http.MethodPost, http.MethodOptions)

	myRouter.HandleFunc(conf.PathMain, productHandler.GetHomePage).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathCategory, productHandler.GetProductsByCategory).Methods(http.MethodGet, http.MethodOptions)

	myRouter.HandleFunc(conf.PathCart, orderHandler.GetCart).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathCart, orderHandler.UpdateCart).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathAddItemToCart, orderHandler.AddItemToCart).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathDeleteItemFromCart, orderHandler.DeleteItemFromCart).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathMakeOrder, orderHandler.MakeOrder).Methods(http.MethodPost, http.MethodOptions)

	myRouter.PathPrefix(conf.PathDocs).Handler(httpSwagger.WrapHandler)
	myRouter.Use(loggingAndCORSHeadersMiddleware)
	http.ListenAndServe(conf.Port, myRouter)
}
