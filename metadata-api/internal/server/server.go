package server

import (
	"context"
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
		WriteTimeout:   5 * time.Second,
	}

	return server
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
	// return s.httpServer.ListenAndServeTLS("./certs/nginx.crt", "./certs/nginx.key")
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	return s.httpServer.Shutdown(ctx)
}
