package usecases

import (
	"client/services"
	"client/typos"
	"fmt"
	"log"
	"os"
	"time"
)

func LatencyUsecase() {
	log.Println("LatencyUsecase: latency usecase execution started")
	client := services.Client{}
	latencyService := services.Latency{}

	/** HTTP3 */
	http3Service, err := client.Init(typos.V3)
	timer := services.Cronos{}

	if err != nil {
		fmt.Println("An error ocurred", err.Error())
		os.Exit(1)
	}

	timer.Begin()
	for i := uint64(0); i < uint64(10); i++ {
		latencyService.Wait()
		log.Println("LatencyUsecase: http3 request")
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

	timer.Begin()
	for i := uint64(0); i < uint64(10); i++ {
		latencyService.Wait()
		log.Println("LatencyUsecase: http1 request")
		http1Service.Get("http://0.0.0.0:8080/ping")
	}
	timer.Stop()
	var http1Time time.Duration = timer.GetTotal()

	fmt.Println("Total time Spent, HTTP1:", http1Time.Seconds(), "HTTP3:", http3Time.Seconds())
	log.Println("LatencyUsecase: latency usecase execution finished")
}
