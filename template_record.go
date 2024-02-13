package easclient

import (
	"bytes"
	"text/template"
)

const recordTemplateStr = `<?xml version="1.0"?>
<records xmlns="http://namespace.otris.de/2010/09/archive/recordExtern">
    <record>
		{{- if .Title}}
		<title>{{.Title}}</title>
		{{end}}
		{{- range $key, $value := .Fields}}
		<field name="{{$key}}">{{$value}}</field>{{end}}{{range $key, $value := .Attachments}}
		<attachment>
			<name>{{$value.Name}}</name>
			<path>{{$value.Path}}</path>

			{{- if $value.Size}}
			<size>{{$value.Size}}</size>
			{{- end}}

			{{- if $value.Register}}
			<register>{{$value.Register}}</register>
			{{- end}}

			{{- if $value.Author}}
			<author>{{$value.Author}}</author>
			{{- end}}
		</attachment>
		{{- end}}
    </record>
</records>
`

var recordTemplate *template.Template

func init() {
	t, err := template.New("putRecord").Parse(recordTemplateStr)
	if err != nil {
		panic(err)
	}
	recordTemplate = t
}

type RecordRequest struct {
	Title       string
	Fields      map[string]string
	Attachments []*RecordRequestAttachment
}

// RecordRequestAttachment is used in a RecordRequest to specify
// spooled attachments to be added to the record.
//
// The EAS expects at least Name and Path to be set.
type RecordRequestAttachment struct {
	Name     string
	Path     string
	Size     uint64
	Register string
	Author   string
}

func renderRecordTemplate(request *RecordRequest) (string, error) {
	buf := new(bytes.Buffer)
	err := recordTemplate.Execute(buf, request)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
