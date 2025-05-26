package main

import (
	"fmt"
	"http3tcc/infra/server"
	"os"
)

func main() {

	var server server.Server

	options := map[string]func(){
		"http1": server.StartHTTP1,
		"http3": server.StartHTTP3,
	}

	callback, exists := options[os.Args[1]]

	if !exists {
		fmt.Println("Opção não existe")
		return
	}

	callback()

}
