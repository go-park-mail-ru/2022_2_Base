package main

import (
	"database/sql"
	"log"
	"net"
	"net/http"
	"os"
	orderdl "serv/microservices/orders/delivery"
	orders "serv/microservices/orders/gen_files"
	orderst "serv/microservices/orders/repository"
	orderuc "serv/microservices/orders/usecase"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"

	//conf "serv/config"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
)

func main() {
	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Println("cant listen port", err)
	}
	//urlDB := "postgres://" + conf.DBSPuser + ":" + conf.DBPassword + "@" + conf.DBHost + ":" + conf.DBPort + "/" + conf.DBName
	urlDB := "postgres://" + os.Getenv("TEST_POSTGRES_USER") + ":" + os.Getenv("TEST_POSTGRES_PASSWORD") + "@" + os.Getenv("TEST_DATABASE_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("TEST_POSTGRES_DB")
	log.Println("conn: ", urlDB)
	db, err := sql.Open("pgx", urlDB)
	if err != nil {
		log.Println("could not connect to database: ", err)
	} else {
		log.Println("database is reachable")
	}
	defer db.Close()

	orderStore := orderst.NewOrderStore(db)

	orderUsecase := orderuc.NewOrderUsecase(orderStore)

	ordersManager := orderdl.NewOrdersManager(orderUsecase)

	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
	grpc_prometheus.Register(server)
	orders.RegisterOrdersWorkerServer(server, ordersManager)

	http.Handle("/metrics", promhttp.Handler())

	log.Println("starting server at :8083")
	err = server.Serve(lis)
	if err != nil {
		log.Println("cant serve", err)
	}
}
