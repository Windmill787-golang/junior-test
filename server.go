package main

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	return &Server{
		httpServer: &http.Server{
			MaxHeaderBytes: 1 << 20, // 1 MB
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
	}
}

func (s *Server) Shutdown() {
	//implement graceful shutdown
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer.Addr = ":" + port
	s.httpServer.Handler = handler
	return s.httpServer.ListenAndServe()
}
