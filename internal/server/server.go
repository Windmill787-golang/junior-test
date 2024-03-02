package server

import (
	"net/http"
	"time"

	"github.com/Windmill787-golang/junior-test/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(config *config.Server) *Server {
	return &Server{
		httpServer: &http.Server{
			MaxHeaderBytes: config.MaxHeaderMB << 20, // 1 MB
			ReadTimeout:    time.Duration(config.ReadTimeOut) * time.Second,
			WriteTimeout:   time.Duration(config.WriteTimeOut) * time.Second,
		},
	}
}

func (s *Server) Shutdown() {
	//TODO: implement graceful shutdown
}

func (s *Server) Run(config *config.Server, handler http.Handler) error {
	s.httpServer.Addr = ":" + config.Port
	s.httpServer.Handler = handler
	return s.httpServer.ListenAndServe()
}
