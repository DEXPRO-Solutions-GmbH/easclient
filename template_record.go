package easclient

import (
	"bytes"
	"text/template"
)

const recordTemplateStr = `<?xml version="1.0"?>
<records xmlns="http://namespace.otris.de/2010/09/archive/recordExtern">
    <record>{{range $key, $value := .Fields}}
		<field name="{{$key}}">{{$value}}</field>{{end}}{{range $key, $value := .Attachments}}
		<attachment>
			<name>{{$value.Name}}</name>
			<path>{{$value.Path}}</path>
			<size>{{$value.Size}}</size>
			<register>{{$value.Register}}</register>
			<author>{{$value.Author}}</author>
		</attachment>{{end}}
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
	Fields      map[string]string
	Attachments []*RecordRequestAttachment
}

// RecordRequestAttachment is used in a RecordRequest to specify
// spooled attachments to be added to the record.
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
