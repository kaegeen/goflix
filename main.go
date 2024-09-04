package main

import (
	"fmt"
	"log"
	"os"

	"goflix/internal/config"
	"goflix/internal/router"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	r := router.NewRouter(cfg)

	log.Printf("Starting server on %s", cfg.ServerAddress)
	if err := r.Start(); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
