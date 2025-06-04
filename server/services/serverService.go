package services

import (
	"fmt"
	"log"
	"main/configs"
	"main/typos"
	"net/http"
	"os"

	"github.com/quic-go/quic-go/http3"
)

type Server struct {
	version typos.ProtocolVersion
	config  configs.ServerConfig
	mux     http.ServeMux
}

func (server *Server) Init(version typos.ProtocolVersion, serverConfigs configs.ServerConfig) {
	server.version = version
	server.config = serverConfigs
	server.mux = *http.NewServeMux()
}

func (server *Server) Listen() {
	log.Println("Server: listen server process")

	if server.version == typos.V3 {
		server.listenHTTP3()
	} else if server.version == typos.V1 {
		server.listenHTTP1()
	} else {
		fmt.Println("Cannot start the propper server")
		os.Exit(1)
	}

}

func (server *Server) SetEndpoint(endpoint typos.Endpoint) {
	log.Println("Endpoint:/", endpoint.Name)
	server.mux.HandleFunc(endpoint.Name, endpoint.Callback)
}

func (server *Server) listenHTTP1() {
	fmt.Printf("Starting HTTP1 Server on http://%s/\n", server.config.Host)
	err := http.ListenAndServe(server.config.Host, &server.mux)
	if err != nil {
		fmt.Println("An error occurred:", err.Error())
		os.Exit(1)
	}
}

func (server *Server) listenHTTP3() {
	fmt.Printf("Starting HTTP3 Server on https://%s/\n", server.config.Host)

	err := http3.ListenAndServeQUIC(server.config.Host, server.config.Cert, server.config.Key, &server.mux)

	if err != nil {
		fmt.Println("An error occurred:", err.Error())
		os.Exit(1)
	}
}
