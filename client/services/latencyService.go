package services

import (
	"log"
	"math/rand"
	"time"
)

type Latency struct {
	samples []time.Duration
}

func (latency *Latency) Wait(index uint64) {
	time.Sleep(latency.samples[index])
}

func (latency *Latency) GenerateLatencySample(samplesNumber uint64) {
	log.Println("Latency: generating latency Distribuition")
	for i := uint64(0); i < samplesNumber; i++ {
		latencyDuration := rand.Intn(150) * int(time.Millisecond)
		latency.samples = append(latency.samples, time.Duration(latencyDuration))
	}
}
