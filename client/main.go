package main

import (
	"client/usecases"
	"flag"
	"io"
	"log"
)

func main() {

	requests, verbose, host := parseFlags()

	if !*verbose {
		log.SetOutput(io.Discard)
	}

	usecases.RequestToServer(*requests, *host)

}

func parseFlags() (*uint64, *bool, *string) {
	requestsFlag := flag.Uint64("requests", 10000, "No requests number setted")
	verboseFlag := flag.Bool("verbose", false, "No verbose selected")
	hostFlag := flag.String("host", "https://0.0.0.0", "No Host informed ex.: https://lskr.com.br")

	flag.Parse()

	return requestsFlag, verboseFlag, hostFlag
}
