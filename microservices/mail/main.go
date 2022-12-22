package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"path/filepath"

	mail "serv/microservices/mail/gen_files"

	"google.golang.org/grpc"
	"gopkg.in/gomail.v2"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"bytes"
	"html/template"
	conf "serv/config"
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
	//var textbody string = "This is the body of the mail"
	fp := filepath.Join("mails_templates", "mail_register", "index.html")
	i := info{Usename: in.Username, ImgEmailLogoSrc: "https://email.reazon.ru/mail.webp", ImgTGLogoSrc: "https://email.reazon.ru/telegram.webp", ImgGitLogoSrc: "https://email.reazon.ru/github.webp"}
	switch in.Type {
	case "orderstatus":
		header = "Изменение статуса заказа"
		switch *in.OrderStatus {
		case "created":
			//textbody = "Заказ номер " + fmt.Sprintf("%d", *in.OrderID) + " оформлен"
			i.BigImgSrc = "https://email.reazon.ru/delivery-img.webp"
			i.OrderID = fmt.Sprintf("%d", *in.OrderID)
			//log.Println(textbody)
			fp = filepath.Join("mails_templates", "mail_orderstatus", "index.html")
		}
	case "promocode":
		header = "Получен новый промокод"
		//textbody = "Ваш новый промокод: " + *in.Promocode
		i.BigImgSrc = "https://email.reazon.ru/gift.webp"
		i.Promocode = *in.Promocode
		fp = filepath.Join("mails_templates", "mail_promocode", "index.html")
	case "greeting":
		header = "Приветствие"
		i.BigImgSrc = "https://email.reazon.ru/girl.svg"
		//fp = filepath.Join("mails_templates", "mail_register", "index.html")
		//textbody = "Здравствуйте, " + in.Username
	}

	//fpReg := filepath.Join("mails_templates", "mail_register", "index.html")
	//fpOrder := filepath.Join("mails_templates", "mail_orderstatus", "index.html")
	//fpPromo := filepath.Join("mails_templates", "mail_promocode", "index.html")

	t := template.New(fp)
	//t := template.New("./mails_templates/mail_register/index.html")

	var err error
	//t, err = t.ParseFiles("./mails_templates/mail_register/index.html")
	t, err = t.ParseFiles(fp)
	if err != nil {
		log.Println(err)
	}

	var tpl bytes.Buffer

	//if err := t.ExecuteTemplate(&tpl, "./mails_templates/mail_register/index.html", i); err != nil {

	//log.Println(result)

	if err := t.ExecuteTemplate(&tpl, "index.html", i); err != nil {
		log.Println(err)
	}

	result := tpl.String()

	msg := gomail.NewMessage()
	msg.SetHeader("From", "Musicialbaum@mail.ru")
	//msg.SetHeader("To", in.Useremail)
	msg.SetHeader("To", "Scorpion1remeres@gmail.com")
	msg.SetHeader("Subject", header)
	//msg.SetBody("text/html", "<b>"+textbody+"</b>")
	msg.SetBody("text/html", result)
	//msg.Attach("/home/User/cat.jpg")
	n := gomail.NewDialer("smtp.mail.ru", 587, "Musicialbaum@mail.ru", conf.MailPassword)
	//Send the email
	if err := n.DialAndSend(msg); err != nil {
		log.Println(err)
		return &mail.Nothing{IsSuccessful: false}, err
	}
	return &mail.Nothing{IsSuccessful: true}, nil
}
