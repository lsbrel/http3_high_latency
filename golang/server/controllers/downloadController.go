package controllers

import (
	"log"
	"net/http"
)

func DownloadController(w http.ResponseWriter, r *http.Request) {
	log.Println("DownloadController: downlaod started")

	http.ServeFile(w, r, "/var/www/html/large_file.png")
}
