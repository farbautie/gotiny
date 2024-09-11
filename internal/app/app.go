package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/farbautie/gotiny/config"
	"github.com/farbautie/gotiny/pkg/database"
	"github.com/farbautie/gotiny/pkg/database/repositories"
	"github.com/farbautie/gotiny/pkg/server"
)

func Run(config *config.Config) {
	log.Printf("Connected to database")
	db, err := database.New(config, database.MaxOpenConns(config.Database.PoolSize))
	if err != nil {
		log.Fatalf("Error creating database pool: %s", err)
	}
	defer db.Close()
	rp := repositories.New(db)
	r := NewRouter(rp)

	log.Printf("Starting server on port %s", config.Http.Port)
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
