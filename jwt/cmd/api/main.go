package main

import (
	"log"

	"github.com/maretrodep/base-auth-go/jwt/config"
	"github.com/maretrodep/base-auth-go/jwt/internal/database"
	"github.com/maretrodep/base-auth-go/jwt/internal/server"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db, err := database.NewConnection(&cfg.DB)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	s := server.NewServer(db, &cfg.Server, &cfg.Auth)
	if err := s.Start(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
