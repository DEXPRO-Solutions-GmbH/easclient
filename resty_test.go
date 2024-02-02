package easclient

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/resty.v1"
)

func Test_copyRestyClient(t *testing.T) {
	t.Run("returns valid copy", func(t *testing.T) {
		original := resty.New()
		copied := copyRestyClient(original)

		require.NotSame(t, original, copied)

		t.Run("copy modifications do not affect original", func(t *testing.T) {
			copied.JSONUnmarshal = func(data []byte, v interface{}) error {
				return errors.New("this is some error")
			}

			require.NotSame(t, original.JSONUnmarshal, copied.JSONUnmarshal)
		})
	})
}
