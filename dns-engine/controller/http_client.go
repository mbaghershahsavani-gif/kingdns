package controller

import (
	"net/http"
)

type HTTPClient struct {
	BaseURL string
	Client  *http.Client
}

func (c HTTPClient) Ping() error {
	return nil
}
