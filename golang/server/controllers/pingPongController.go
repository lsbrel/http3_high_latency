package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func PingPongController(w http.ResponseWriter, r *http.Request) {
	log.Println("PingPongController: Message received")

	response := map[string]string{
		"message": "pong",
		"time":    time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
