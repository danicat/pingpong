package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "defines the address to listen for requests")

func main() {
	flag.Parse()

	s := NewServer()

	log.Printf("starting server on: %s", *addr)
	http.ListenAndServe(*addr, s)
}
