package internal

import (
	"errors"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// TestPrelude is used by tests of this library to ensure that they behave as expected in the context
// of our environments.
func TestPrelude(t *testing.T) {
	// Having this function be called ensures that the init() function is called.
}

func init() {
	err := godotenv.Load(".env")
	if errors.Is(err, os.ErrNotExist) {
		return
	}
}
