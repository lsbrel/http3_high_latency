package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"main/services"
	"main/typos"
	"main/usecases"
)

func main() {

	servertype, verbose, latency, loss := parseFlags()
	latencyService := services.Latency{}

	if !*verbose {
		log.SetOutput(io.Discard)
	}

	if latency != nil {
		latencyService.SetLatency(*latency)
	}

	if loss != nil {
		latencyService.SetPacketLoss(*loss)
	}

	switch *servertype {
	case string(typos.HTTP1):
		usecases.Http1Usecase()
	case string(typos.HTTP3):
		usecases.Http3Usecase()
	default:
		fmt.Println("No server type selected, exiting")
	}

	latencyService.RemoveLatency()

}

func parseFlags() (*string, *bool, *string, *string) {
	serverTypeFlag := flag.String("server", "none", "No server selected. Ex.: ./main --server={http1,http3}")
	verboseFlag := flag.Bool("verbose", false, "No verbose selected")
	latencyFlag := flag.String("latency", "0ms", "No latency has been defined, Ex.: --latency 100ms")
	lossFlag := flag.String("loss", "0%", "No latency has been defined, Ex.: --loss 10%")

	flag.Parse()

	return serverTypeFlag, verboseFlag, latencyFlag, lossFlag
}
