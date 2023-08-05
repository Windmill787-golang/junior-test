package main

import "net/http"

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	//return http server
	return &Server{}
}

func (s *Server) Shutdown() {
	//implement graceful shutdown
}

func (s *Server) Run(port int) {
	//listen and server on specific port
}
