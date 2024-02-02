package easclient

import (
	"time"

	"github.com/google/uuid"
)

type Record struct {
	DocumentType           string              `xml:"documentType"`
	MasterId               uuid.UUID           `xml:"masterId"`
	ArchiveDateTime        time.Time           `xml:"archiveDateTime"`
	ID                     uuid.UUID           `xml:"id"`
	Version                string              `xml:"version"`
	ArchiverLogin          string              `xml:"archiverLogin"`
	Archiver               string              `xml:"archiver"`
	InitialArchiver        string              `xml:"initialArchiver"`
	InitialArchiverLogin   string              `xml:"initialArchiverLogin"`
	InitialArchiveDateTime time.Time           `xml:"initialArchiveDateTime"`
	Fields                 []*RecordField      `xml:"field"`
	Attachments            []*RecordAttachment `xml:"attachment"`
}

type RecordField struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

type RecordAttachment struct {
	Name     string    `xml:"name"`
	Size     int       `xml:"size"`
	Register string    `xml:"register"`
	Author   string    `xml:"author"`
	ID       uuid.UUID `xml:"id"`
}
