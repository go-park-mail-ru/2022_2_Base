package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	session "serv/microservices/auth/gen_files"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
)

func main() {
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Println("cant listen port", err)
	}

	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
	grpc_prometheus.Register(server)
	session.RegisterAuthCheckerServer(server, NewSessionManager())
	http.Handle("/metrics", promhttp.Handler())
	log.Println("starting server at :8082")
	server.Serve(lis)
}

const sessKeyLen = 10

type SessionManager struct {
	session.UnimplementedAuthCheckerServer

	mu       sync.RWMutex
	sessions map[string]*session.Session
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		mu:       sync.RWMutex{},
		sessions: map[string]*session.Session{},
	}
}

func (sm *SessionManager) Create(ctx context.Context, in *session.Session) (*session.SessionID, error) {
	log.Println("call Create", in)
	newUUID := uuid.New()
	id := &session.SessionID{
		ID: newUUID.String(),
	}
	sm.mu.Lock()
	sm.sessions[id.ID] = in
	sm.mu.Unlock()

	return id, nil
}

func (sm *SessionManager) Check(ctx context.Context, in *session.SessionID) (*session.Session, error) {
	log.Println("call Check", in)
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	if sess, ok := sm.sessions[in.ID]; ok {
		return sess, nil
	}
	return nil, grpc.Errorf(codes.NotFound, "session not found")
}

func (sm *SessionManager) Delete(ctx context.Context, in *session.SessionID) (*session.Nothing, error) {
	log.Println("call Delete", in)
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.sessions, in.ID)
	return &session.Nothing{IsSuccessful: true}, nil
}