package main


import (
	"log"
	"gvc/internal/config"
	"gvc/internal/server"
)

func main(){
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	if err := server.Start(cfg); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
