package controller

import "net/http"

type Transport struct {
	Client *http.Client
}

func NewTransport() *Transport {
	return &Transport{
		Client: &http.Client{},
	}
}
