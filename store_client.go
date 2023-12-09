package easclient

import (
	"context"
	"errors"

	"gopkg.in/resty.v1"
)

type StoreClient struct {
	c *resty.Client
}

func NewStoreClient(c *resty.Client) *StoreClient {
	return &StoreClient{c: c}
}

func (c *StoreClient) newRequest(ctx context.Context) (*resty.Request, error) {
	claims := UserClaimsFromContext(ctx)
	if claims == nil {
		return nil, errors.New("missing user claims in context object")
	}

	req := c.c.NewRequest()
	req.SetContext(ctx)
	req.SetHeader("Accept", "application/json")
	claims.SetOnHeader(req.Header)

	return req, nil
}
