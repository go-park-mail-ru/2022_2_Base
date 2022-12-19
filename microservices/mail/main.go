package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	mail "serv/microservices/mail/gen_files"

	"google.golang.org/grpc"
	"gopkg.in/gomail.v2"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	conf "serv/config"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
)

func main() {
	lis, err := net.Listen("tcp", ":8084")
	if err != nil {
		log.Println("cant listen port", err)
	}

	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
	grpc_prometheus.Register(server)
	mail.RegisterMailServiceServer(server, NewMailManager())
	http.Handle("/metrics", promhttp.Handler())
	log.Println("starting server at :8084")
	server.Serve(lis)
}

const sessKeyLen = 10

type MailManager struct {
	mail.UnimplementedMailServiceServer
}

func NewMailManager() *MailManager {
	return &MailManager{}
}

func (mm *MailManager) SendMail(ctx context.Context, in *mail.Mail) (*mail.Nothing, error) {
	log.Println("call SendMail", in)
	var header string = "Письмо"
	var textbody string = "This is the body of the mail"
	switch in.Type {
	case "orderstatus":
		header = "Изменение статуса заказа"
		switch *in.OrderStatus {
		case "created":
			textbody = "Заказ номер " + fmt.Sprintf("%d", *in.OrderID) + " оформлен"
			log.Println(textbody)
		}
	case "promocode":
		header = "Получен новый промокод"
		textbody = "Ваш новый промокод: " + *in.Promocode
	case "greeting":
		header = "Приветствие"
		textbody = "Здравствуйте, " + in.Username
	}
	msg := gomail.NewMessage()
	msg.SetHeader("From", "Musicialbaum@mail.ru")
	msg.SetHeader("To", in.Useremail)
	msg.SetHeader("Subject", header)
	msg.SetBody("text/html", "<b>"+textbody+"</b>")
	//msg.Attach("/home/User/cat.jpg")
	n := gomail.NewDialer("smtp.mail.ru", 587, "Musicialbaum@mail.ru", conf.MailPassword)
	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		log.Println(err)
		return &mail.Nothing{IsSuccessful: false}, err
	}
	return &mail.Nothing{IsSuccessful: true}, nil
}
