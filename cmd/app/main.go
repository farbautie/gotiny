package main

import (
	"log"

	"github.com/farbautie/gotiny/config"
	"github.com/farbautie/gotiny/internal/app"
)

func main() {
	cfg, err := config.DefineConfig()
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}
	app.Run(cfg)
}
