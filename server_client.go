package easclient

import "gopkg.in/resty.v1"

type ServerClient struct {
	c *resty.Client
}

// NewServerClient creates a new client for server interaction.
func NewServerClient(c *resty.Client) *ServerClient {
	return &ServerClient{c: c}
}
