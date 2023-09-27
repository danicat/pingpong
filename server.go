package main

import (
	"log"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	var s Server

	s.mux = http.NewServeMux()

	s.setupRoutes()

	return &s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.mux.ServeHTTP(w, req)
}

func (s *Server) handlePing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request received: %s %s", r.Method, r.RequestURI)
		w.Write([]byte(`{"message":"pong"}`))
	}
}
