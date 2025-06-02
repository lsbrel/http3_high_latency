package server

import (
	"fmt"

	"http3tcc/domain/controller"
	"http3tcc/infra/entities"
	"http3tcc/infra/router"

	"github.com/quic-go/quic-go/http3"
)

func startHTTP3Server() {

	var serverConfig = entities.ServerConfig{
		Host: "0.0.0.0:4242", Cert: "./certs/cert.crt", Key: "./certs/key.key"}

	var routes = []entities.Route{
		{Name: "/ping", Callback: controller.BasicController},
	}

	fmt.Printf("Iniciando servidor HTTP/3 na porta https://%s\n", serverConfig.Host)

	router.AddRoute(routes)

	err := http3.ListenAndServeQUIC(serverConfig.Host, serverConfig.Cert, serverConfig.Key, router.GetMuxer())

	if err != nil {
		fmt.Printf("Failed to configure HTTP/3: %v", err)
	}

}
