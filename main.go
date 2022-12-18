package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "serv/docs"
	"serv/repository"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	muxprom "gitlab.com/msvechla/mux-prometheus/pkg/middleware"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	deliv "serv/delivery"
	usecase "serv/usecase"

	conf "serv/config"

	httpSwagger "github.com/swaggo/http-swagger"

	auth "serv/microservices/auth/gen_files"
	orders "serv/microservices/orders/gen_files"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
)

func loggingAndCORSHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI, r.Method)

		//for local tests
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

var (
	sessManager auth.AuthCheckerClient
)
var (
	ordersManager orders.OrdersWorkerClient
)

func main() {
	myRouter := mux.NewRouter()
	//urlDB := "postgres://" + conf.DBSPuser + ":" + conf.DBPassword + "@" + conf.DBHost + ":" + conf.DBPort + "/" + conf.DBName
	urlDB := "postgres://" + os.Getenv("TEST_POSTGRES_USER") + ":" + os.Getenv("TEST_POSTGRES_PASSWORD") + "@" + os.Getenv("TEST_DATABASE_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("TEST_POSTGRES_DB")
	log.Println("conn: ", urlDB)
	db, err := sql.Open("pgx", urlDB)
	if err != nil {
		log.Println("could not connect to database")
	} else {
		log.Println("database is reachable")
	}
	defer db.Close()

	grcpConnAuth, err := grpc.Dial(
		"auth:8082",
		//"localhost:8082",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor),
	)
	if err != nil {
		log.Println("cant connect to grpc auth")
	} else {
		log.Println("connected to grpc auth service")
	}
	defer grcpConnAuth.Close()

	grcpConnOrders, err := grpc.Dial(
		"orders:8083",
		//"localhost:8083",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor),
	)
	if err != nil {
		log.Println("cant connect to grpc orders")
	} else {
		log.Println("connected to grpc orders service")
	}
	defer grcpConnOrders.Close()

	sessManager = auth.NewAuthCheckerClient(grcpConnAuth)
	ordersManager = orders.NewOrdersWorkerClient(grcpConnOrders)

	userStore := repository.NewUserStore(db)
	productStore := repository.NewProductStore(db)

	userUsecase := usecase.NewUserUsecase(userStore, sessManager)
	productUsecase := usecase.NewProductUsecase(productStore, ordersManager)

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
	myRouter.HandleFunc(conf.PathSeacrh, productHandler.GetProductsBySearch).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathSuggestions, productHandler.GetSuggestions).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathRecommendations, productHandler.GetRecommendations).Methods(http.MethodGet, http.MethodOptions)

	userRouter.HandleFunc(conf.PathProfile, userHandler.GetUser).Methods(http.MethodGet, http.MethodOptions)
	userRouter.HandleFunc(conf.PathProfile, userHandler.ChangeProfile).Methods(http.MethodPost, http.MethodOptions)
	userRouter.HandleFunc(conf.PathAvatar, userHandler.SetAvatar).Methods(http.MethodPost, http.MethodOptions)
	userRouter.HandleFunc(conf.PathPassword, userHandler.ChangePassword).Methods(http.MethodPost, http.MethodOptions)

	myRouter.HandleFunc(conf.PathComments, orderHandler.GetComments).Methods(http.MethodGet, http.MethodOptions)
	userRouter.HandleFunc(conf.PathMakeComment, orderHandler.CreateComment).Methods(http.MethodPost, http.MethodOptions)

	cartRouter.HandleFunc("", orderHandler.GetCart).Methods(http.MethodGet, http.MethodOptions)
	cartRouter.HandleFunc("", orderHandler.UpdateCart).Methods(http.MethodPost, http.MethodOptions)
	cartRouter.HandleFunc(conf.PathAddItemToCart, orderHandler.AddItemToCart).Methods(http.MethodPost, http.MethodOptions)
	cartRouter.HandleFunc(conf.PathDeleteItemFromCart, orderHandler.DeleteItemFromCart).Methods(http.MethodPost, http.MethodOptions)
	cartRouter.HandleFunc(conf.PathMakeOrder, orderHandler.MakeOrder).Methods(http.MethodPost, http.MethodOptions)
	cartRouter.HandleFunc(conf.PathGetOrders, orderHandler.GetOrders).Methods(http.MethodGet, http.MethodOptions)

	myRouter.PathPrefix(conf.PathDocs).Handler(httpSwagger.WrapHandler)
	myRouter.Use(loggingAndCORSHeadersMiddleware)

	instrumentation := muxprom.NewDefaultInstrumentation()
	myRouter.Use(instrumentation.Middleware)
	myRouter.Path("/metrics").Handler(promhttp.Handler())

	amw := deliv.NewAuthMiddleware(userUsecase)

	userRouter.Use(amw.CheckAuthMiddleware)
	cartRouter.Use(amw.CheckAuthMiddleware)

	http.ListenAndServe(conf.Port, myRouter)
}
