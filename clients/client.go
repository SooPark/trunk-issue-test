package clients

import (
	"github.com/imroc/req/v3"
)

type Client struct {
	Client *req.Client
}

func NewClient(baseURL, apiKey string) *Client {
	c := req.C()
	return &Client{Client: c}
}
