package usecases

import (
	"main/configs"
	"main/controllers"
	"main/services"
	"main/typos"
)

func Http3Usecase() {
	server := services.Server{}
	serverConfigs := configs.ServerConfig{Host: "0.0.0.0:4242", Cert: "/etc/letsencrypt/live/lskr.com.br/fullchain.pem", Key: "/etc/letsencrypt/live/lskr.com.br/privkey.pem"}

	server.Init(typos.V3, serverConfigs)
	server.SetEndpoint(typos.Endpoint{Name: "/", Callback: controllers.PingPongController})

	server.Listen()
}
