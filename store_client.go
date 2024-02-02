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

func (c *StoreClient) newRequestJSON(ctx context.Context) (*resty.Request, error) {
	return newRequestJSON(ctx, c.c)
}

func (c *StoreClient) newRequestXML(ctx context.Context) (*resty.Request, error) {
	return newRequestXML(ctx, c.c)
}

func newRequestJSON(ctx context.Context, c *resty.Client) (*resty.Request, error) {
	return newRequest(ctx, c, "application/json")
}

func newRequestXML(ctx context.Context, c *resty.Client) (*resty.Request, error) {
	return newRequest(ctx, c, "application/xml")
}

func newRequest(ctx context.Context, c *resty.Client, contentType string) (*resty.Request, error) {
	claims := UserClaimsFromContext(ctx)
	if claims == nil {
		return nil, errors.New("missing user claims in context object")
	}

	req := c.NewRequest()
	req.SetContext(ctx)
	if contentType != "" {
		req.SetHeader("Accept", contentType)
	}
	claims.SetOnHeader(req.Header)

	return req, nil
}
