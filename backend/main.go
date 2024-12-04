package main

import (
	"backend/internal/api"
	"backend/internal/config"
	"backend/internal/db"
	"log"
)

func main() {
	pgConfig, err := config.LoadPgConfig()
	if err != nil {
		log.Fatalf("invalid API config: %v", err)
	}
	apiConfig, err := config.LoadApiConfig()
	if err != nil {
		log.Fatalf("invalid API config: %v", err)
	}

	db.InitDB(pgConfig)
	api.Serve(apiConfig)
}
