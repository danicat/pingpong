package main

func (s *Server) setupRoutes() {
	s.mux.HandleFunc("/ping", s.handlePing())
}
