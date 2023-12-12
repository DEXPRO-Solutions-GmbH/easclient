package easclient

import (
	"context"
	"fmt"
)

type ConfigurationParameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ConfigurationTemplate struct {
	Name       string                   `json:"name"`
	Parameters []ConfigurationParameter `json:"parameters"`
}

type PutStoreRequest struct {
	ConfigurationTemplate `json:"configurationTemplate"`
}

func (c *ServerClient) PutStore(ctx context.Context, storeName string, request *PutStoreRequest) error {
	req, err := newRequest(ctx, c.c)
	if err != nil {
		return err
	}

	req.SetBody(request)

	res, err := req.Put("/" + storeName)
	if err != nil {
		return err
	}

	if status := res.StatusCode(); status != 201 {
		return fmt.Errorf("unexpected response status %v", status)
	}

	return nil
}
