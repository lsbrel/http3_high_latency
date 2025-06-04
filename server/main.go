package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"main/typos"
	"main/usecases"
)

func main() {

	servertype, verbose := parseFlags()

	if !*verbose {
		log.SetOutput(io.Discard)
	}

	switch *servertype {
	case string(typos.HTTP1):
		usecases.Http1Usecase()
	case string(typos.HTTP3):
		usecases.Http3Usecase()
	default:
		fmt.Println("No server type selected, exiting")
	}

}

func parseFlags() (*string, *bool) {
	serverType := flag.String("server", "none", "No server selected. Ex.: ./main --server={http1,http3}")
	verboseFlag := flag.Bool("verbose", false, "No verbose selected")
	flag.Parse()

	return serverType, verboseFlag
}
