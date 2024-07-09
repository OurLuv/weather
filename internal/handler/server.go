package handler

import (
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func (s *Server) Start() error {
	err := s.server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func NewServer(r http.Handler) *Server {
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return &Server{
		server: s,
	}
}
