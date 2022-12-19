package main

import (
	"context"
	"log"
	"net"
	"net/http"

	mail "serv/microservices/mail/gen_files"

	"google.golang.org/grpc"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	//"github.com/emersion/go-smtp"

	//"github.com/emersion/go-smtp"
	gomail "gopkg.in/gomail.v2"

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

	// // Setup an unencrypted connection to a local mail server.
	// c, err := smtp.Dial("localhost:25")
	// if err != nil {
	// 	return err
	// }
	// defer c.Close()

	// // Set the sender and recipient, and send the email all in one step.
	// to := []string{"recipient@example.net"}
	// msg := strings.NewReader("To: recipient@example.net\r\n" +
	// 	"Subject: discount Gophers!\r\n" +
	// 	"\r\n" +
	// 	"This is the email body.\r\n")
	// err = c.SendMail("sender@example.org", to, msg)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("starting server at :8082")
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
			textbody = "Заказ номер " + string(*in.OrderID) + "оформлен"
		}
	case "promocode":
		header = "Получен новый промокод"
		textbody = "Ваш новый промокод: " + *in.Promocode
	case "greeting":
		header = "Приветствие"
	}
	msg := gomail.NewMessage()
	// msg.SetHeader("From", "<paste your gmail account here>")
	// msg.SetHeader("To", "<paste the email address you want to send to>")
	// msg.SetHeader("Subject", "<paste the subject of the mail>")
	// msg.SetBody("text/html", "<b>This is the body of the mail</b>")
	// msg.Attach("/home/User/cat.jpg")
	// n := gomail.NewDialer("smtp.gmail.com", 587, "<paste your gmail account here>", "<paste Google password or app password here>")
	msg.SetHeader("From", "Musicialbaum@mail.ru")
	msg.SetHeader("To", "Scorpion1remeres@gmail.com")
	//msg.SetHeader("To", in.Useremail)
	msg.SetHeader("Subject", header)
	msg.SetBody("text/html", "<b>"+textbody+"</b>")
	//msg.Attach("/home/User/cat.jpg")
	//n := gomail.NewDialer("smtp.gmail.com", 587, "Musicialbaum@mail.ru", "Musicial2022")
	n := gomail.NewDialer("smtp.mail.ru", 587, "Musicialbaum@mail.ru", conf.MailPassword)
	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		//panic(err)
		log.Println(err)
		return &mail.Nothing{IsSuccessful: false}, err
	}
	return &mail.Nothing{IsSuccessful: true}, nil
}
