package easclient

import (
	"time"

	"github.com/google/uuid"
)

type HeaderFields struct {
	DocumentType           string    `json:"_documentType"`
	MasterId               uuid.UUID `json:"_masterId"`
	ArchiveDateTime        time.Time `json:"_archiveDateTime"`
	Id                     uuid.UUID `json:"_id"`
	Version                string    `json:"_version"`
	ArchiverLogin          string    `json:"_archiverLogin"`
	InitialArchiverLogin   string    `json:"_initialArchiverLogin"`
	InitialArchiveDateTime time.Time `json:"_initialArchiveDateTime"`
}

type RecordFields map[string]string

type Record struct {
	HeaderFields HeaderFields        `json:"headerFields"`
	RecordFields RecordFields        `json:"recordFields"`
	Attachments  []*RecordAttachment `json:"attachments"`
}

// GetHeaderField returns either the value of the given header field or an empty string if the field does not exist.
func (rec *Record) GetHeaderField(name string) string {
	return rec.RecordFields[name]
}
