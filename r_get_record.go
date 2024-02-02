package easclient

import (
	"context"
	"encoding/xml"

	"github.com/google/uuid"
)

func (c *StoreClient) GetRecord(ctx context.Context, id uuid.UUID) (*Record, error) {
	req, err := c.newRequestXML(ctx)
	if err != nil {
		return nil, err
	}

	type RecordResponse struct {
		XMLName xml.Name `xml:"records"`
		Record  *Record  `xml:"record"`
	}

	var result RecordResponse

	req.SetResult(&result)
	res, err := req.Get("/record/" + id.String())
	if err != nil {
		return nil, err
	}

	if _, err := isErrorResponse(res); err != nil {
		return nil, err
	}

	return result.Record, nil
}
