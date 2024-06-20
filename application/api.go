package application

import (
	"fmt"
	"log"
	"net/http"

	"github.com/avshetty1980/good-growth-interview/database"
	"github.com/avshetty1980/good-growth-interview/handler"
)

type APIServer struct {
	addr  string
	store database.Store
}

func NewAPIServer(addr string, store database.Store) *APIServer {
	return &APIServer{
		addr:  addr,
		store: store,
	}
}

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer next.ServeHTTP(w, r)

		log.Printf("method %s, path: %s", r.Method, r.URL.Path)
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	messageService := handler.NewMessageService(s.store)
	messageService.RegisterRoutes(router)

	fmt.Println("loadroutes here")
	server := &http.Server{
		Addr:    s.addr,
		Handler: RequestLoggerMiddleware(router),
	}

	log.Printf("Server started at port %s", s.addr)
	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	return nil
}
