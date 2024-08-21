package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func New(port string, routes http.Handler) Server {
	server := Server{}

	server.httpServer = &http.Server{
		Addr:           ":" + port, // localhost:port
		Handler:        routes,
		MaxHeaderBytes: 1 << 28, // 1 Mb
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   60 * time.Second,
	}

	return server
}

func (s *Server) RegisterRouter(router http.Handler) {
	log.Println("Регистрирую роутеры")

}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	return s.httpServer.Shutdown(ctx)
}
