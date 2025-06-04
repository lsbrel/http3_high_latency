package services

import (
	"log"
	"math/rand"
	"time"
)

type Latency struct {
	duration time.Duration
}

func (latency *Latency) SetLatency(duration int) {
	latency.duration = time.Duration(duration)
}

func (*Latency) Wait() {
	random := rand.Intn(150)
	var latency time.Duration = time.Duration(random)
	log.Println("Latency added (ms):", random)
	time.Sleep(latency * time.Millisecond)
}
