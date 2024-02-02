package easclient

import (
	"context"
	"encoding/xml"
	"time"

	"github.com/google/uuid"
)

type RecordAttachment struct {
	Body            string    `json:"body"`
	Name            string    `json:"name"`
	Size            string    `json:"size"`
	Register        string    `json:"register"`
	Author          string    `json:"author"`
	Type            string    `json:"type"`
	DocumentType    string    `json:"documentType"`
	Id              uuid.UUID `json:"id"`
	FileId          uuid.UUID `json:"fileId"`
	MasterId        uuid.UUID `json:"masterId"`
	Version         string    `json:"version"`
	ArchiveDateTime time.Time `json:"archiveDateTime"`
}

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
