package usecases

import (
	"client/services"
	"client/typos"
	"fmt"
	"log"
	"os"
	"time"
)

func NoLatencyUsecase(requestNumber uint64) {
	log.Println("NolatencyUsecase: no latency usecase execution started")
	client := services.Client{}

	/** HTTP3 */
	http3Service, err := client.Init(typos.V3)
	timer := services.Cronos{}

	if err != nil {
		fmt.Println("An error ocurred", err.Error())
		os.Exit(1)
	}

	log.Println("NolatencyUsecase: HTTP3 sending requests")
	timer.Begin()
	for i := uint64(0); i < requestNumber; i++ {
		http3Service.Get("https://0.0.0.0:4242/ping")
	}
	timer.Stop()
	var http3Time time.Duration = timer.GetTotal()

	/** HTTP1 */
	http1Service, err := client.UpdateClient(typos.V1)

	if err != nil {
		fmt.Println("An error ocurred", err.Error())
		os.Exit(1)
	}

	log.Println("NolatencyUsecase: HTTP1 sending requests")
	timer.Begin()
	for i := uint64(0); i < requestNumber; i++ {
		http1Service.Get("http://0.0.0.0:8080/ping")
	}
	timer.Stop()
	var http1Time time.Duration = timer.GetTotal()

	fmt.Println("Total time Spent (seconds), HTTP1:", http1Time.Seconds(), "HTTP3:", http3Time.Seconds())
	log.Println("NolatencyUsecase: no latency usecase execution finished")
}
