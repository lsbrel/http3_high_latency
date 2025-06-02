package main

import (
	"client/domain/cases"
	"io"
	"log"
	"os"
)

func main() {

	log.SetOutput(io.Discard)

	typeCase := os.Args[1]

	var requests []uint64 = []uint64{1, 10, 100, 500, 1000, 2500, 5000, 10000}

	if typeCase == "http1" {
		for _, request := range requests {
			cases.Http1Case("http://0.0.0.0:8080/ping", request)
		}
	} else if typeCase == "http3" {
		for _, request := range requests {
			cases.Http3Case("https://0.0.0.0:4242/ping", request)
		}
	}

}
