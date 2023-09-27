package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"runtime"
)

// build flags
var Version string = "develop"

var addr = flag.String("addr", ":8080", "defines the address to listen for requests")

func main() {
	flag.Parse()

	ctx := context.Background()

	err := run(ctx, *addr)
	if err != nil {
		log.Fatalln(err)
	}
}

func run(ctx context.Context, addr string) error {
	log.Printf("GOMAXPROCS: %d", runtime.GOMAXPROCS(0))
	log.Printf("go runtime version: %s", runtime.Version())
	log.Printf("server version: %s", Version)

	go func() {
		s := NewServer()

		log.Printf("starting server on: %s", addr)
		err := http.ListenAndServe(addr, s)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	<-ctx.Done()
	return ctx.Err()
}
