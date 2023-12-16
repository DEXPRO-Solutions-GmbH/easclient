package easclient

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Record struct {
	HeaderFields struct {
		DocumentType           string    `json:"_documentType"`
		MasterId               uuid.UUID `json:"_masterId"`
		ArchiveDateTime        time.Time `json:"_archiveDateTime"`
		Id                     uuid.UUID `json:"_id"`
		Version                string    `json:"_version"`
		ArchiverLogin          string    `json:"_archiverLogin"`
		InitialArchiverLogin   string    `json:"_initialArchiverLogin"`
		InitialArchiveDateTime time.Time `json:"_initialArchiveDateTime"`
	} `json:"headerFields"`
	RecordFields map[string]string   `json:"recordFields"`
	Attachments  []*RecordAttachment `json:"attachments"`
}

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
	req, err := c.newRequest(ctx)
	if err != nil {
		return nil, err
	}

	var result Record

	req.SetResult(&result)
	res, err := req.Get("/record/" + id.String())
	if err != nil {
		return nil, err
	}

	if status := res.StatusCode(); status != 200 {
		return nil, fmt.Errorf("unexpected response status %v", status)
	}

	return &result, nil
}
