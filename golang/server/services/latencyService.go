package services

import (
	"fmt"
	"log"
	"os/exec"
)

type Latency struct {
}

func (latency *Latency) SetLatency(delay string) {
	_, err := exec.LookPath("tc")

	if err != nil {
		fmt.Println("Latency cannot be setted")
	}

	log.Println("latency with variation of", delay)
	cmd := exec.Command("tc", "qdisc", "change", "dev", "eth0", "root", "netem", "delay", delay)
	cmd.Run()
}

func (latency *Latency) SetPacketLoss(loss string) {
	_, err := exec.LookPath("tc")

	if err != nil {
		fmt.Println("Latency cannot be setted")
	}

	log.Println("packet loss of added", loss)
	cmd := exec.Command("tc", "qdisc", "change", "dev", "eth0", "root", "netem", "loss", loss)
	cmd.Run()
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
