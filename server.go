package main

import (
	"encoding/json"
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
	type response struct {
		Message string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request received: %s %s", r.Method, r.RequestURI)

		contentType := r.Header.Get("content-type")
		if contentType != "" && contentType != "application/json" {
			w.WriteHeader(http.StatusNotImplemented)
			log.Printf("content type not implemented: %s", contentType)
			return
		}

		switch r.Method {
		case http.MethodGet:
			w.Header().Set("content-type", "application/json")
			json.NewEncoder(w).Encode(response{Message: "pong"})
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}

	}
}
