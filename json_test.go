package easclient

import (
	stdjson "encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestUnmarshalOfSimilarKeys contains tests which show a challenge when dealing with the
// JSON representation of EAS responses:
//
// Standard fields are sometimes mixed with custom fields in the same
// object. The std lib's json package does not strictly enforce that only the field
// with a matching case is being used. Have a look at the tests and these links for details:
//
// - https://github.com/golang/go/issues/14750
// - https://github.com/go-json-experiment/json
func TestUnmarshalOfSimilarKeys(t *testing.T) {
	t.Run("std lib does unmarshal from exact key", func(t *testing.T) {
		t.Run("uses last key 1", func(t *testing.T) {
			attachmentStr := `{
				"id": "0dd018f8-bf23-455d-8214-44e76b24e5db",
				"Id": "00000000-0000-0000-0000-000000000000"
			}`

			attachment := RecordAttachment{}

			err := stdjson.Unmarshal([]byte(attachmentStr), &attachment)
			require.NoError(t, err)

			// This test asserts that the std lib behaves "weired" - given that the Id struct field
			// should be unmarshalled from the JSON field "id" and not "Id".
			//
			// In practice however, the std lib will use the last key it finds, accepting all case-variations.
			assert.Equal(t, "00000000-0000-0000-0000-000000000000", attachment.Id.String())
		})

		t.Run("uses last key 2", func(t *testing.T) {
			attachmentStr := `{
				"Id": "00000000-0000-0000-0000-000000000000",
				"id": "0dd018f8-bf23-455d-8214-44e76b24e5db"
			}`

			attachment := RecordAttachment{}

			err := stdjson.Unmarshal([]byte(attachmentStr), &attachment)
			require.NoError(t, err)

			// This test asserts that the std lib behaves "weired" - given that the Id struct field
			// should be unmarshalled from the JSON field "id" and not "Id".
			//
			// In practice however, the std lib will use the last key it finds, accepting all case-variations.
			assert.Equal(t, "0dd018f8-bf23-455d-8214-44e76b24e5db", attachment.Id.String())
		})
	})

	t.Run("v2 json lib does unmarshal from exact key", func(t *testing.T) {
		t.Run("uses correct key 1", func(t *testing.T) {
			attachmentStr := `{
				"id": "0dd018f8-bf23-455d-8214-44e76b24e5db",
				"Id": "00000000-0000-0000-0000-000000000000"
			}`

			attachment := RecordAttachment{}

			err := unmarshalJSON([]byte(attachmentStr), &attachment)
			require.NoError(t, err)

			// This test asserts that the std lib behaves "weired" - given that the Id struct field
			// should be unmarshalled from the JSON field "id" and not "Id".
			//
			// In practice however, the std lib will use the last key it finds, accepting all case-variations.
			assert.Equal(t, "0dd018f8-bf23-455d-8214-44e76b24e5db", attachment.Id.String())
		})

		t.Run("uses correct key 2", func(t *testing.T) {
			attachmentStr := `{
				"Id": "00000000-0000-0000-0000-000000000000",
				"id": "0dd018f8-bf23-455d-8214-44e76b24e5db"
			}`

			attachment := RecordAttachment{}

			err := unmarshalJSON([]byte(attachmentStr), &attachment)
			require.NoError(t, err)

			// This test asserts that the std lib behaves "weired" - given that the Id struct field
			// should be unmarshalled from the JSON field "id" and not "Id".
			//
			// In practice however, the std lib will use the last key it finds, accepting all case-variations.
			assert.Equal(t, "0dd018f8-bf23-455d-8214-44e76b24e5db", attachment.Id.String())
		})
	})
}
