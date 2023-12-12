package easclient

import "gopkg.in/resty.v1"

type ServerClient struct {
	c *resty.Client
}

func NewServerClient(c *resty.Client) *ServerClient {
	return &ServerClient{c: c}
}
