package cases

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Http1Case(url string, request uint64) {

	start := time.Now()
	for i := uint64(0); i < request; i++ {
		_, err := http.Get(url)

		if err != nil {
			log.Fatalf("Failed to make request: %v", err)
		}

	}
	end := time.Since(start).Seconds()

	fmt.Println(request, "Requests em", end, "Segundos")

}
