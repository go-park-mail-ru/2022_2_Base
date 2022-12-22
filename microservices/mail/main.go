package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"

	mail "serv/microservices/mail/gen_files"

	"google.golang.org/grpc"
	"gopkg.in/gomail.v2"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"bytes"
	"html/template"
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
	err = server.Serve(lis)
	if err != nil {
		log.Println("cant serve", err)
	}
}

type MailManager struct {
	mail.UnimplementedMailServiceServer
}

func NewMailManager() *MailManager {
	return &MailManager{}
}

type info struct {
	Usename         string
	Promocode       string
	OrderID         string
	BigImgSrc       string
	ImgEmailLogoSrc string
	ImgTGLogoSrc    string
	ImgGitLogoSrc   string
}

func (mm *MailManager) SendMail(ctx context.Context, in *mail.Mail) (*mail.Nothing, error) {
	log.Println("call SendMail", in)
	var header string = "Письмо"
	//fp := filepath.Join("mails_templates", "mail_register", "index.html")
	fp := filepath.Join("microservices", "mail", "mails_templates", "mail_register", "index.html")
	i := info{Usename: in.Username, ImgEmailLogoSrc: "https://email.reazon.ru/mail.png", ImgTGLogoSrc: "https://email.reazon.ru/telegram.png", ImgGitLogoSrc: "https://email.reazon.ru/github.png"}
	switch in.Type {
	case "orderstatus":
		header = "Изменение статуса заказа"
		switch *in.OrderStatus {
		case "created":
			i.BigImgSrc = "https://email.reazon.ru/delivery-img.png"
			i.OrderID = fmt.Sprintf("%d", *in.OrderID)
			//fp = filepath.Join("mails_templates", "mail_orderstatus", "index.html")
			fp = filepath.Join("microservices", "mail", "mails_templates", "mail_orderstatus", "index.html")
		}
	case "promocode":
		header = "Получен новый промокод"
		i.BigImgSrc = "https://email.reazon.ru/gift.png"
		i.Promocode = *in.Promocode
		//fp = filepath.Join("mails_templates", "mail_promocode", "index.html")
		fp = filepath.Join("microservices", "mail", "mails_templates", "mail_promocode", "index.html")
	case "greeting":
		header = "Приветствие"
		i.BigImgSrc = "https://email.reazon.ru/girl.png"
	}

	t := template.New(fp)
	var err error
	t, err = t.ParseFiles(fp)
	if err != nil {
		log.Println(err)
	}

	var tpl bytes.Buffer
	if err := t.ExecuteTemplate(&tpl, "index.html", i); err != nil {
		log.Println(err)
	}
	result := tpl.String()

	msg := gomail.NewMessage()
	msg.SetHeader("From", "Musicialbaum@mail.ru")
	msg.SetHeader("To", in.Useremail)
	msg.SetHeader("Subject", header)
	msg.SetBody("text/html", result)
	//msg.Attach("/home/User/cat.jpg")
	n := gomail.NewDialer("smtp.mail.ru", 587, "Musicialbaum@mail.ru", os.Getenv("MAILPASSWORD"))
	//Send the email
	if err := n.DialAndSend(msg); err != nil {
		log.Println(err)
		return &mail.Nothing{IsSuccessful: false}, err
	}
	return &mail.Nothing{IsSuccessful: true}, nil
}
