package easclient

import (
	"time"

	"github.com/google/uuid"
)

type Record struct {
	Type                   string              `xml:"type"`
	Title                  string              `xml:"title"`
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

func (r *Record) GetHeaderFieldVal(name string) string {
	for _, field := range r.Fields {
		if field.Name == name {
			return field.Value
		}
	}
	return ""
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
