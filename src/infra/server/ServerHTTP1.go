package server

import (
	"fmt"
	"http3tcc/domain/controller"
	"net/http"
)

func startHTTP1Server() {

	fmt.Println("Iniciando servidor HTTP/1 na porta http://localhost:8080")

	http.HandleFunc("/ping", controller.BasicController)
	http.ListenAndServe(":8080", nil)

}
