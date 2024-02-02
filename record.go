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

type RecordField struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

type Record struct {
	DocumentType           string         `xml:"documentType"`
	MasterId               uuid.UUID      `xml:"masterId"`
	ArchiveDateTime        time.Time      `xml:"archiveDateTime"`
	ID                     uuid.UUID      `xml:"id"`
	Version                string         `xml:"version"`
	ArchiverLogin          string         `xml:"archiverLogin"`
	Archiver               string         `xml:"archiver"`
	InitialArchiver        string         `xml:"initialArchiver"`
	InitialArchiverLogin   string         `xml:"initialArchiverLogin"`
	InitialArchiveDateTime time.Time      `xml:"initialArchiveDateTime"`
	Fields                 []*RecordField `xml:"field"`
	Attachment             struct {
		Name     string `xml:"name"`
		Size     int    `xml:"size"`
		Register string `xml:"register"`
		Author   string `xml:"author"`
		ID       string `xml:"id"`
	} `xml:"attachment"`
}
