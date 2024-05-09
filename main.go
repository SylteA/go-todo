package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"todo/api"
)

var port = flag.String("port", "8080", "http service address")
var host = flag.String("host", "localhost", "http service host")

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	go func() {
		<-ctx.Done()
		log.Println("shutting down")
		os.Exit(0)
	}()

	s := api.NewServer(1)
	s.AddRoutes()

	log.Println("listening on " + *host + ":" + *port)
	return http.ListenAndServe(net.JoinHostPort(*host, *port), s.Router)
}

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		log.Println(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
