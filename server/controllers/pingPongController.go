package controllers

import (
	"log"
	"net/http"
)

func PingPongController(w http.ResponseWriter, r *http.Request) {
	log.Println("Mensagem recebida")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "pong"}`))
}
