package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/farbautie/gotiny/config"
	"github.com/farbautie/gotiny/pkg/server"
)

func Run(config *config.Config) {
	r := NewRouter()
	srv := server.New(r, server.Port(config.Http.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Printf("Received signal %s, shutting down...", s)
	case err := <-srv.Notify():
		log.Printf("Server stopped with error: %s", err)
	}

	if err := srv.Shutdown(); err != nil {
		log.Printf("Error shutting down server: %s", err)
	}
}
