package easclient

import (
	"github.com/go-json-experiment/json"
	"gopkg.in/resty.v1"
)

func copyRestyClient(c *resty.Client) *resty.Client {
	// dereference the pointer and copy the value
	cc := *c
	return &cc
}

func adaptRestyClient(c *resty.Client) {
	c.JSONUnmarshal = func(data []byte, v interface{}) error {
		opts := []json.Options{
			json.MatchCaseInsensitiveNames(false),
		}
		return unmarshalJSON(data, v, opts...)
	}
}
