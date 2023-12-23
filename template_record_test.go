package easclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_renderPutRecordTemplate(t *testing.T) {
	req := &RecordRequest{
		Fields: map[string]string{
			"Creditor": "DE123456789",
			"Debitor":  "DE987654321",
		},
	}

	xml, err := renderRecordTemplate(req)
	require.NoError(t, err)
	expected := `<?xml version="1.0"?>
<records xmlns="http://namespace.otris.de/2010/09/archive/recordExtern">
    <record>
		<field name="Creditor">DE123456789</field>
		<field name="Debitor">DE987654321</field>
    </record>
</records>
`

	assert.Equal(t, expected, xml)
}
