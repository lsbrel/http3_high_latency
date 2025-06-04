package main

import (
	"client/typos"
	"client/usecases"
	"flag"
	"fmt"
	"io"
	"log"
)

func main() {

	usecase, requests, verbose := parseFlags()

	if !*verbose {
		log.SetOutput(io.Discard)
	}

	switch *usecase {
	case string(typos.LatencyDisabled):
		usecases.NoLatencyUsecase(*requests)
	case string(typos.LatencyEnabled):
		usecases.LatencyUsecase(*requests)
	case string(typos.TLS):
		fmt.Println("TLS")
	default:
		fmt.Println("No use case selected. Ex.: ./main --use-case={nolatency,latency,tls}")
	}

}

func parseFlags() (*string, *uint64, *bool) {
	usecaseFlag := flag.String("usecase", "none", "No use case selected. Ex.: ./main --use-case={disabled,enabled,tls}")
	requestsFlag := flag.Uint64("requests", 10000, "No requests number setted")
	verboseFlag := flag.Bool("verbose", false, "No verbose selected")
	flag.Parse()

	return usecaseFlag, requestsFlag, verboseFlag
}
