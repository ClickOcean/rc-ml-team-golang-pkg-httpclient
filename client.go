package httpclient

import (
	"net/http"
	"time"
)

type client struct {
	client *http.Client
}

func New() *client {
	return &client{
		client: &http.Client{},
	}
}

func (c *client) WithTimeout(sec int) *client {
	c.client.Timeout = time.Duration(sec) * time.Second
	return c
}

func (c client) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}
