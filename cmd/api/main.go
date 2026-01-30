package main

import (
	config "CRUD/internal/Config"
	"CRUD/internal/db"
	"CRUD/internal/server"
	"fmt"
	"log"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error")
	}

	client, database, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("db error")
	}

	defer func() {
		if err := db.Disconnect(client); err != nil {
			log.Printf("mongo disconnect error %v", err)
		}
	}()

	router := server.NewRouter(database)

	addr := fmt.Sprintf(":%s", cfg.ServerPort)

	if err := router.Run(addr); err != nil {
		log.Fatalf("server failed")
	}
}
