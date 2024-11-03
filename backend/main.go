package main

import (
	"backend/internal/config"
	"backend/internal/db"
	"fmt"
)

func main() {
	fmt.Println("hello from GoLang")
	pgConfig, err := config.LoadPgConfig()
	if err != nil {
		panic(err)
	}
	db.InitDB(pgConfig)
}
