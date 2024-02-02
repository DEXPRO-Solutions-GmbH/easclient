package easclient

import (
	"context"
	"errors"

	"gopkg.in/resty.v1"
)

type StoreClient struct {
	c *resty.Client
}

// NewStoreClient creates a new client for store interaction.
func NewStoreClient(c *resty.Client) *StoreClient {
	c = copyRestyClient(c)
	adaptRestyClient(c)
	return &StoreClient{c: c}
}

func (c *StoreClient) newRequest(ctx context.Context) (*resty.Request, error) {
	return newRequest(ctx, c.c)
}

func newRequest(ctx context.Context, c *resty.Client) (*resty.Request, error) {
	claims := UserClaimsFromContext(ctx)
	if claims == nil {
		return nil, errors.New("missing user claims in context object")
	}

	req := c.NewRequest()
	req.SetContext(ctx)
	req.SetHeader("Accept", "application/json")
	claims.SetOnHeader(req.Header)

	return req, nil
}
