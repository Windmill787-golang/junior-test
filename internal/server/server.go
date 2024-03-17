package server

import (
	"net/http"
	"time"

	"github.com/Windmill787-golang/junior-test/internal/config"
	"github.com/Windmill787-golang/junior-test/internal/handler"
)

type Server struct {
	httpServer *http.Server
	port       string
	handler    *handler.Handler
}

func NewServer(handler *handler.Handler, config *config.Server) *Server {
	return &Server{
		httpServer: &http.Server{
			MaxHeaderBytes: config.MaxHeaderMB << 20, // 1 MB
			ReadTimeout:    time.Duration(config.ReadTimeOut) * time.Second,
			WriteTimeout:   time.Duration(config.WriteTimeOut) * time.Second,
		},
		port:    config.Port,
		handler: handler,
	}
}

func (s *Server) Shutdown() {
	//TODO: implement graceful shutdown
}

func (s *Server) Run() error {
	s.httpServer.Addr = ":" + s.port
	//s.httpServer.Handler = handler
	s.handler.InitRoutesNew()

	return s.httpServer.ListenAndServe()
}
