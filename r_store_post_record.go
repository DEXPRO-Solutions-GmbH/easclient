package easclient

import (
	"context"
	"fmt"
)

type PostRecordResponse struct {
	Records []struct {
		Id   string `json:"id"`
		Link struct {
			Type  string `json:"type"`
			Title string `json:"title"`
			Href  string `json:"href"`
		} `json:"link"`
	} `json:"records"`
}

func (c *StoreClient) PostRecord(ctx context.Context, request *RecordRequest) (*PostRecordResponse, error) {
	req, err := c.newRequest(ctx)
	if err != nil {
		return nil, err
	}

	xml, err := renderRecordTemplate(request)
	if err != nil {
		return nil, err
	}

	req.SetMultiValueFormData(map[string][]string{
		"record": {xml},
	})

	var result PostRecordResponse

	req.SetResult(&result)
	res, err := req.Post("/record")
	if err != nil {
		return nil, err
	}

	if status := res.StatusCode(); status != 200 {
		return nil, fmt.Errorf("unexpected response status %v", status)
	}

	return &result, nil
}
