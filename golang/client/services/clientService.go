package services

import (
	"client/typos"
	"crypto/tls"
	"errors"
	"log"
	"net/http"

	"github.com/quic-go/quic-go/http3"
)

type Client struct {
	service *http.Client
	Version typos.ProtocolVersion
}

func (client *Client) Init(version typos.ProtocolVersion) (*http.Client, error) {
	client.Version = version

	if client.Version == typos.V1 {
		log.Println("Client: Http1 service created")
		client.service = &http.Client{}

	} else if client.Version == typos.V3 {
		log.Println("Client: Http3 service created")

		client.service = &http.Client{
			Transport: &http3.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
	} else {
		return nil, errors.New("cannot create client for the specified version")
	}

	return client.service, nil
}

func (client *Client) UpdateClient(version typos.ProtocolVersion) (*http.Client, error) {
	if version == client.Version {
		return client.service, nil
	}
	return client.Init(version)
}
