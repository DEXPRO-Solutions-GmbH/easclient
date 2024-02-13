package easclient

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecord_UnmarshalXML(t *testing.T) {
	input := `
<record>
        <type>Invoice</type>
        <title>[Rechnung] B.R.T. GmbH - 47110815</title>
        <createdDateTime>2024-02-08T11:16:26</createdDateTime>
        <lastModifiedDateTime>2024-02-08T11:17:14</lastModifiedDateTime>
        <creator>Pust, Simon</creator>
        <lastModifier>Pust, Simon</lastModifier>
        <creatorLogin>pust</creatorLogin>
        <masterId>4b4012ae-c33a-484d-b444-1567efd3bc89</masterId>
        <archiveDateTime>2024-02-08T11:32:25+01:00</archiveDateTime>
        <id>4b4012ae-c33a-484d-b444-1567efd3bc89</id>
        <version>0</version>
        <archiverLogin>job</archiverLogin>
        <archiver>Jobs, Steve</archiver>
        <initialArchiver>Jobs, Steve</initialArchiver>
        <initialArchiverLogin>job</initialArchiverLogin>
        <initialArchiveDateTime>2024-02-08T11:32:25+01:00</initialArchiveDateTime>
</record>`

	var record Record

	err := xml.Unmarshal([]byte(input), &record)
	require.NoError(t, err)

	assert.Equal(t, "Invoice", record.Type)
	assert.Equal(t, "[Rechnung] B.R.T. GmbH - 47110815", record.Title)
}
