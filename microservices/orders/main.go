package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	orderdl "serv/microservices/orders/delivery"
	orders "serv/microservices/orders/gen_files"
	orderst "serv/microservices/orders/repository"
	orderuc "serv/microservices/orders/usecase"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
)

func main() {
	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Println("cant listen port", err)
	}
	urlDB := "postgres://" + os.Getenv("TEST_POSTGRES_USER") + ":" + os.Getenv("TEST_POSTGRES_PASSWORD") + "@" + os.Getenv("TEST_DATABASE_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("TEST_POSTGRES_DB")
	config, _ := pgxpool.ParseConfig(urlDB)
	config.MaxConns = 120
	db, err := pgxpool.New(context.Background(), config.ConnString())
	log.Println("conn: ", urlDB)
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
	go func() {
		log.Println("starting collect mectrics :8093")
		err = http.ListenAndServe(":8093", nil)
		if err != nil {
			log.Println("cant serve metrics", err)
		}
	}()

	log.Println("starting server at :8083")
	err = server.Serve(lis)
	if err != nil {
		log.Println("cant serve", err)
	}
}
