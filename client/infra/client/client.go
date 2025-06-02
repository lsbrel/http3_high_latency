package client

import (
	"crypto/tls"
	"net/http"

	"github.com/quic-go/quic-go/http3"
)

func InitClient() *http.Client {
	rt := &http3.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	defer rt.Close()

	client := &http.Client{
		Transport: rt,
	}

	return client
}
