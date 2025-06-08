package usecases

import (
	"main/configs"
	"main/controllers"
	"main/services"
	"main/typos"
)

func Http1Usecase() {
	server := services.Server{}
	serverConfigs := configs.ServerConfig{Host: "0.0.0.0:8080"}

	server.Init(typos.V1, serverConfigs)
	server.SetEndpoint(typos.Endpoint{Name: "/ping", Callback: controllers.PingPongController})

	server.Listen()
}
