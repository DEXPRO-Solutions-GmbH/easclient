package easclient

import (
	"context"
	"strings"

	"github.com/google/uuid"
)

type PostRecordResponse struct {
	Records []struct {
		Id   uuid.UUID `json:"id"`
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

	req.SetMultipartField("record", "record.xml", "application/xml", strings.NewReader(xml))

	var result PostRecordResponse

	req.SetResult(&result)
	res, err := req.Post("/record")
	if err != nil {
		return nil, err
	}

	if _, err := isErrorResponse(res); err != nil {
		return nil, err
	}

	return &result, nil
}
