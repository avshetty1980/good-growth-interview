package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer next.ServeHTTP(w, r)

		log.Printf("method %s, path: %s", r.Method, r.URL.Path)
	}
}

func (s *APIServer) Run(ctx context.Context) error {
	router := loadRoutes()
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
