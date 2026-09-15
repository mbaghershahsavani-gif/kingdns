package controller

import (
	"net/http"
)

type Client struct {
	BaseURL string
	HTTP    *http.Client
}

func NewClient(url string) *Client {
	return &Client{
		BaseURL: url,
		HTTP:    &http.Client{},
	}
}
