package server

import (
	"fmt"
	"log"
	"net/http"
	"gvc/internal/config"
	"gvc/internal/service/api"
)


func Start(cfg *config.Config) {
	http.HandleFunc("/api", api.Handlers)

	addr := fmt.Sprintf(":%s", cfg.Port)
	
	log.Printf("Starting server on %s", addr)

	http.ListenAndServe(addr, nil)
}
