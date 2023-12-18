package eastest

import (
	"os"
	"testing"
)

// SkipInCI is used to skip tests if they can not be executed in the current environment.
func SkipInCI(t *testing.T) {
	if os.Getenv("GITHUB_ACTION") != "" {
		t.Skip("Tests can't currently be run in the GitHub Action environment. We will first have to make it possible to run the EAS or a mocked variant in CI")
	}
}
