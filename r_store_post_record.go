package easclient

import (
	"context"
	"encoding/xml"
	"strings"

	"github.com/google/uuid"
)

type PostRecordResponse struct {
	XMLName xml.Name `xml:"recordArchive"`
	ID      struct {
		Value uuid.UUID `xml:",chardata"`
		Type  string    `xml:"type,attr"`
		Href  string    `xml:"href,attr"`
	} `xml:"id"`
}

func (c *StoreClient) PostRecord(ctx context.Context, request *RecordRequest) (*PostRecordResponse, error) {
	req, err := c.newRequestXML(ctx)
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
