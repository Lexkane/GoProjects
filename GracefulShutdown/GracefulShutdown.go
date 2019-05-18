package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

func startHttpServer() *http.Server {
	server := &http.Server{Addr: ":8080", Handler: handler}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			// here should be error handling
		}
	}()

	// here is signal capturing
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// here we wait for SIGINT (pkill -2)
	<-stop

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if err := server.Shutdown(ctx); err != nil {
		// here we handle error
	}
}

func main() {
	startHttpServer()
	log.Printf("main: starting HTTP server")
}
