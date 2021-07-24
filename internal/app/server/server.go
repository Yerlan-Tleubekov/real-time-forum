package server

import (
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(handler *http.ServeMux) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         mainConfig.Addr,
			Handler:      handler,
			WriteTimeout: 10 * time.Second,
			ReadTimeout:  10 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	log.Println("starting api server at", s.httpServer.Addr)

	return s.httpServer.ListenAndServe()
}
