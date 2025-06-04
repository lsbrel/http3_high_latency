package main

import (
	"client/usecases"
	"flag"
	"fmt"
	"io"
	"log"
)

type Flags string

const (
	LatencyDisabled Flags = "nolatency"
	LatencyEnabled  Flags = "latency"
	TLS             Flags = "tls"
)

func main() {

	usecaseFlag := flag.String("usecase", "none", "No use case selected. Ex.: ./main --use-case={disabled,enabled,tls}")
	verbose := flag.Bool("verbose", false, "No verbose selected")
	flag.Parse()

	if !*verbose {
		log.SetOutput(io.Discard)
	}

	switch *usecaseFlag {
	case string(LatencyDisabled):
		usecases.NoLatencyUsecase()
	case string(LatencyEnabled):
		usecases.LatencyUsecase()
	case string(TLS):
		fmt.Println("TLS")
	default:
		fmt.Println("No use case selected. Ex.: ./main --use-case={nolatency,latency,tls}")
	}

}
