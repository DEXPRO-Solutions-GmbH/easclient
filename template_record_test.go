package easclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_renderPutRecordTemplate(t *testing.T) {
	t.Run("simple record", func(t *testing.T) {
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
	})

	t.Run("record with attachments", func(t *testing.T) {
		req := &RecordRequest{
			Fields: map[string]string{
				"Creditor": "DE123456789",
				"Debitor":  "DE987654321",
			},
			Attachments: []*RecordRequestAttachment{
				{
					Name:     "Test.pdf",
					Path:     "ff0351a7-aa00-4269-ab49-fb4172e3193f.tmp",
					Size:     56865,
					Register: "Testregister",
					Author:   "Test",
				},
				{
					Name:     "Test.pdf",
					Path:     "ff0351a7-aa00-4269-ab49-fb4172e3193f.tmp",
					Size:     56865,
					Register: "",
					Author:   "",
				},
			},
		}

		xml, err := renderRecordTemplate(req)
		require.NoError(t, err)
		expected := `<?xml version="1.0"?>
<records xmlns="http://namespace.otris.de/2010/09/archive/recordExtern">
    <record>
		<field name="Creditor">DE123456789</field>
		<field name="Debitor">DE987654321</field>
		<attachment>
			<name>Test.pdf</name>
			<path>ff0351a7-aa00-4269-ab49-fb4172e3193f.tmp</path>
			<size>56865</size>
			<register>Testregister</register>
			<author>Test</author>
		</attachment>
		<attachment>
			<name>Test.pdf</name>
			<path>ff0351a7-aa00-4269-ab49-fb4172e3193f.tmp</path>
			<size>56865</size>
		</attachment>
    </record>
</records>
`
		assert.Equal(t, expected, xml)
	})

	t.Run("record with attachment with minimal fields", func(t *testing.T) {
		req := &RecordRequest{
			Fields: map[string]string{
				"Creditor": "DE123456789",
				"Debitor":  "DE987654321",
			},
			Attachments: []*RecordRequestAttachment{
				{
					Name: "Test.pdf",
					Path: "ff0351a7-aa00-4269-ab49-fb4172e3193f.tmp",
				},
			},
		}

		xml, err := renderRecordTemplate(req)
		require.NoError(t, err)
		expected := `<?xml version="1.0"?>
<records xmlns="http://namespace.otris.de/2010/09/archive/recordExtern">
    <record>
		<field name="Creditor">DE123456789</field>
		<field name="Debitor">DE987654321</field>
		<attachment>
			<name>Test.pdf</name>
			<path>ff0351a7-aa00-4269-ab49-fb4172e3193f.tmp</path>
		</attachment>
    </record>
</records>
`
		assert.Equal(t, expected, xml)
	})
}
