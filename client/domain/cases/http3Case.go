package cases

import (
	"client/infra/client"
	"fmt"
	"log"
	"time"
)

func Http3Case(url string, request uint64) {

	client := client.InitClient()

	start := time.Now()
	for i := uint64(0); i < request; i++ {
		chooseLatency()
		_, err := client.Get(url)

		if err != nil {
			log.Fatalf("Failed to make request: %v", err)
		}

	}
	end := time.Since(start).Seconds()

	fmt.Println(request, "Requests em", end, "Segundos")
}
