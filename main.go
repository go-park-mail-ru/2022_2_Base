package main

import (
	"context"
	"log"
	"net/http"

	_ "serv/docs"
	"serv/repository"

	"github.com/gorilla/mux"

	deliv "serv/delivery"
	usecase "serv/usecase"

	conf "serv/config"
	baseErrors "serv/domain/errors"
	"serv/domain/model"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/jackc/pgx/v5/pgxpool"
)

func loggingAndCORSHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI, r.Method)

		//for tests on local server
		origin := r.Header.Get("Origin")
		if origin == "http://89.208.198.137:8081" || origin == "http://127.0.0.1:8081" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		for header := range conf.Headers {
			w.Header().Set(header, conf.Headers[header])
		}

		next.ServeHTTP(w, r)
	})
}

type authenticationMiddleware struct {
	userUsecase usecase.UserUsecase
}

func WithUser(ctx context.Context, user *model.UserProfile) context.Context {
	return context.WithValue(ctx, "userdata", user)
}

func (amw *authenticationMiddleware) checkAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := r.Cookie("session_id")
		if err == http.ErrNoCookie {
			log.Println("no session")
			deliv.ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
			return
		}
		usName, err := amw.userUsecase.GetSession(session.Value)
		if err != nil {
			log.Println("no session2")
			deliv.ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
			return
		}

		// hashTok := HashToken{Secret: []byte("Base")}
		// token := r.Header.Get("csrf")
		// curSession := model.Session{ID: 0, UserUUID: session.Value}
		// flag, err := hashTok.CheckCSRFToken(&curSession, token)
		// if err != nil || !flag {
		// 	log.Println("no csrf token")
		// 	ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
		// 	return
		// }

		user, err := amw.userUsecase.GetUserByUsername(usName)
		if err != nil {
			log.Println("err get user ", err)
			deliv.ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
			return
		}
		addresses, err := amw.userUsecase.GetAddressesByUserID(user.ID)
		if err != nil {
			log.Println("err get adresses ", err)
			deliv.ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
			return
		}
		payments, err := amw.userUsecase.GetPaymentMethodByUserID(user.ID)
		if err != nil {
			log.Println("err get payments ", err)
			deliv.ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
			return
		}

		if user.Email == "" {
			deliv.ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 401)
			return
		}

		userData := model.UserProfile{ID: user.ID, Email: user.Email, Username: user.Username, Phone: *user.Phone, Avatar: *user.Avatar, Address: addresses, PaymentMethods: payments}

		next.ServeHTTP(w, r.WithContext(WithUser(r.Context(), &userData)))
	})
}

func main() {
	myRouter := mux.NewRouter()
	urlDB := "postgres://" + conf.DBSPuser + ":" + conf.DBPassword + "@" + conf.DBHost + ":" + conf.DBPort + "/" + conf.DBName
	//urlDB := "postgres://" + os.Getenv("TEST_POSTGRES_USER") + ":" + os.Getenv("TEST_POSTGRES_PASSWORD") + "@" + os.Getenv("TEST_DATABASE_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("TEST_POSTGRES_DB")
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

	userRouter := myRouter.PathPrefix("/api/v1/user").Subrouter()
	cartRouter := myRouter.PathPrefix("/api/v1/cart").Subrouter()

	myRouter.HandleFunc(conf.PathLogin, sessionHandler.Login).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathLogOut, sessionHandler.Logout).Methods(http.MethodDelete, http.MethodOptions)
	myRouter.HandleFunc(conf.PathSignUp, sessionHandler.SignUp).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathSessions, sessionHandler.GetSession).Methods(http.MethodGet, http.MethodOptions)

	myRouter.HandleFunc(conf.PathProductByID, productHandler.GetProductByID).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathMain, productHandler.GetHomePage).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathCategory, productHandler.GetProductsByCategory).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathComments, orderHandler.GetComments).Methods(http.MethodGet, http.MethodOptions)

	userRouter.HandleFunc(conf.PathProfile, userHandler.GetUser).Methods(http.MethodGet, http.MethodOptions)
	userRouter.HandleFunc(conf.PathProfile, userHandler.ChangeProfile).Methods(http.MethodPost, http.MethodOptions)
	userRouter.HandleFunc(conf.PathAvatar, userHandler.SetAvatar).Methods(http.MethodPost, http.MethodOptions)
	userRouter.HandleFunc(conf.PathMakeComment, orderHandler.CreateComment).Methods(http.MethodPost, http.MethodOptions)

	cartRouter.HandleFunc("", orderHandler.GetCart).Methods(http.MethodGet, http.MethodOptions)
	cartRouter.HandleFunc("", orderHandler.UpdateCart).Methods(http.MethodPost, http.MethodOptions)
	cartRouter.HandleFunc(conf.PathAddItemToCart, orderHandler.AddItemToCart).Methods(http.MethodPost, http.MethodOptions)
	cartRouter.HandleFunc(conf.PathDeleteItemFromCart, orderHandler.DeleteItemFromCart).Methods(http.MethodPost, http.MethodOptions)
	cartRouter.HandleFunc(conf.PathMakeOrder, orderHandler.MakeOrder).Methods(http.MethodPost, http.MethodOptions)
	cartRouter.HandleFunc(conf.PathGetOrders, orderHandler.GetOrders).Methods(http.MethodGet, http.MethodOptions)

	myRouter.PathPrefix(conf.PathDocs).Handler(httpSwagger.WrapHandler)
	myRouter.Use(loggingAndCORSHeadersMiddleware)

	amw := authenticationMiddleware{*userUsecase}
	userRouter.Use(amw.checkAuthMiddleware)
	cartRouter.Use(amw.checkAuthMiddleware)

	http.ListenAndServe(conf.Port, myRouter)
}
