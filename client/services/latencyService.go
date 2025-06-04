package services

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

type Latency struct {
}

func (latency *Latency) Wait() {
	time.Sleep(1 * time.Millisecond)
}

func (latency *Latency) SetLatency() {
	_, err := exec.LookPath("tc")

	if err != nil {
		fmt.Println("Latency cannot be setted")
	}

	log.Println("latency with variation of 20ms to 100ms")
	cmd1 := exec.Command("tc", "qdisc change dev eth0 root netem delay 100ms 50ms")
	cmd1.Run()

	log.Println("packet loss of 1% added")
	cmd2 := exec.Command("tc", "qdisc change dev eth0 root netem loss 90%")
	cmd2.Run()

}

func (latency *Latency) RemoveLatency() {
	_, err := exec.LookPath("tc")

	if err != nil {

		fmt.Println("Latency cannot be setted")
	}

	log.Println("clear the network latency")
	cmd := exec.Command("tc", "qdisc del dev eth0 root netem")
	cmd.Run()
}
