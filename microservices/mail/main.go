package main

import (
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

	msg := gomail.NewMessage()
	// msg.SetHeader("From", "<paste your gmail account here>")
	// msg.SetHeader("To", "<paste the email address you want to send to>")
	// msg.SetHeader("Subject", "<paste the subject of the mail>")
	// msg.SetBody("text/html", "<b>This is the body of the mail</b>")
	// msg.Attach("/home/User/cat.jpg")

	// n := gomail.NewDialer("smtp.gmail.com", 587, "<paste your gmail account here>", "<paste Google password or app password here>")
	msg.SetHeader("From", "Musicialbaum@gmail.com")
	msg.SetHeader("To", "Scorpion1remeres@gmail.com")
	msg.SetHeader("Subject", "LOL")
	msg.SetBody("text/html", "<b>This is the body of the mail</b>")
	//msg.Attach("/home/User/cat.jpg")

	n := gomail.NewDialer("smtp.gmail.com", 587, "Musicialbaum@gmail.com", "Musicial2022")

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}

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
	server.Serve(lis)

	// auth := sasl.NewPlainClient("", "user@example.com", "password")

	// // Connect to the server, authenticate, set the sender and recipient,
	// // and send the email all in one step.
	// to := []string{"recipient@example.net"}
	// msg := strings.NewReader("To: recipient@example.net\r\n" +
	// 	"Subject: discount Gophers!\r\n" +
	// 	"\r\n" +
	// 	"This is the email body.\r\n")
	// err := smtp.SendMail("mail.example.com:25", auth, "sender@example.org", to, msg)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

const sessKeyLen = 10

type MailManager struct {
	mail.UnimplementedMailServiceServer
}

func NewMailManager() *MailManager {
	return &MailManager{}
}

// func (sm *SessionManager) Create(ctx context.Context, in *session.Session) (*session.SessionID, error) {
// 	log.Println("call Create", in)
// 	newUUID := uuid.New()
// 	id := &session.SessionID{
// 		ID: newUUID.String(),
// 	}
// 	sm.mu.Lock()
// 	sm.sessions[id.ID] = in
// 	sm.mu.Unlock()

// 	return id, nil
// }

// func (sm *SessionManager) Check(ctx context.Context, in *session.SessionID) (*session.Session, error) {
// 	log.Println("call Check", in)
// 	sm.mu.RLock()
// 	defer sm.mu.RUnlock()
// 	if sess, ok := sm.sessions[in.ID]; ok {
// 		return sess, nil
// 	}
// 	return nil, grpc.Errorf(codes.NotFound, "session not found")
// }

// func (sm *SessionManager) Delete(ctx context.Context, in *session.SessionID) (*session.Nothing, error) {
// 	log.Println("call Delete", in)
// 	sm.mu.Lock()
// 	defer sm.mu.Unlock()
// 	delete(sm.sessions, in.ID)
// 	return &session.Nothing{IsSuccessful: true}, nil
// }

// func (sm *SessionManager) ChangeEmail(ctx context.Context, in *session.NewLogin) (*session.Nothing, error) {
// 	log.Println("call ChangeEmail", in)
// 	sm.mu.Lock()
// 	defer sm.mu.Unlock()
// 	sm.sessions[in.ID] = &session.Session{Login: in.Login}
// 	return &session.Nothing{IsSuccessful: true}, nil
// }
