package easclient

import (
	"bytes"
	"text/template"
)

const recordTemplateStr = `<?xml version="1.0"?>
<records xmlns="http://namespace.otris.de/2010/09/archive/recordExtern">
    <record>{{range $key, $value := .Fields}}
		<field name="{{$key}}">{{$value}}</field>{{end}}
    </record>
</records>
`

var recordTempalte *template.Template

func init() {
	t, err := template.New("putRecord").Parse(recordTemplateStr)
	if err != nil {
		panic(err)
	}
	recordTempalte = t
}

type RecordRequest struct {
	Fields map[string]string
}

func renderRecordTemplate(request *RecordRequest) (string, error) {
	buf := new(bytes.Buffer)
	err := recordTempalte.Execute(buf, request)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
