package main

import (
	"backend/internal/api"
	"backend/internal/config"
	"backend/internal/db"
	"fmt"
	"log"
)

func main() {
	fmt.Println("hello from GoLang")
	pgConfig, err := config.LoadPgConfig()
	if err != nil {
		log.Fatalf("invalid PostgreSQL config: %v", err)
	}
	db.InitDB(pgConfig)

	apiConfig, err := config.LoadApiConfig()
	if err != nil {
		log.Fatalf("invalid API config: %v", err)
	}
	api.Serve(apiConfig)
}
