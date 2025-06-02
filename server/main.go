package main

import (
	"fmt"
	"http3tcc/domain/enum"
	"http3tcc/infra/server"
	"io"
	"log"
	"os"
)

func main() {

	log.SetOutput(io.Discard)

	var server server.Server

	options := map[string]func(){
		enum.Http1: server.StartHTTP1,
		enum.Http3: server.StartHTTP3,
	}

	callback, exists := options[os.Args[1]]

	if !exists {
		fmt.Println("Opção não existe")
		return
	}

	callback()

}
