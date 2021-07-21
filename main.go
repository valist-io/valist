package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/valist-io/registry/internal/core"
	"github.com/valist-io/registry/internal/http"
)

func main() {
	client, err := core.NewClient()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	server := http.NewServer(client)

	log.Println("starting web server")
	go server.ListenAndServe()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.Shutdown(ctx)
}
