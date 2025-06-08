package services

import (
	"fmt"
	"log"
	"time"
)

type Cronos struct {
	start time.Time
	end   time.Time
	total time.Duration
}

func (cronos *Cronos) Begin() {
	log.Println("Cronos: time counter started")
	cronos.start = time.Now()
}

func (cronos *Cronos) Stop() {
	log.Println("Cronos: time counter stoped")
	cronos.end = time.Now()
	cronos.total = cronos.end.Sub(cronos.start)
}

func (cronos *Cronos) GetTotal() time.Duration {
	return cronos.total
}

func (cronos *Cronos) ShowDuration() {
	fmt.Println("Time Spent:", cronos.total.Seconds())
}
