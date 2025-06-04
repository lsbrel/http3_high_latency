package usecases

import (
	"main/configs"
	"main/controllers"
	"main/services"
	"main/typos"
)

func Http3Usecase() {
	server := services.Server{}
	serverConfigs := configs.ServerConfig{Host: "0.0.0.0:4242", Cert: "./certs/cert.crt", Key: "./certs/key.key"}

	server.Init(typos.V3, serverConfigs)
	server.SetEndpoint(typos.Endpoint{Name: "/ping", Callback: controllers.PingPongController})

	server.Listen()
}
