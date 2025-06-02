package cases

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func Http1Case(url string, request uint64) {

	start := time.Now()
	for i := uint64(0); i < request; i++ {
		chooseLatency()
		_, err := http.Get(url)

		if err != nil {
			log.Fatalf("Failed to make request: %v", err)
		}

	}
	end := time.Since(start).Seconds()

	fmt.Println(request, "Requests em", end, "Segundos")

}

func chooseLatency() {
	choose := rand.Intn(3)

	if choose == 2 {
		// fmt.Println("Latência de 50 milisegundos adicionada")
		time.Sleep(50 * time.Millisecond)
	} else if choose == 1 {
		// fmt.Println("Latência de 15 milisegundos adicionada")
		time.Sleep(15 * time.Millisecond)
	} else {
		return
	}
}
