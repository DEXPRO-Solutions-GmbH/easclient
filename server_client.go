package easclient

import "gopkg.in/resty.v1"

type ServerClient struct {
	c *resty.Client
}

// NewServerClient creates a new client for server interaction.
func NewServerClient(c *resty.Client) *ServerClient {
	c = copyRestyClient(c)
	adaptRestyClient(c)
	return &ServerClient{c: c}
}
