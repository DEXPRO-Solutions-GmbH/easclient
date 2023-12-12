package easclient_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/DEXPRO-Solutions-GmbH/easclient"
	"github.com/joho/godotenv"
	"gopkg.in/resty.v1"
)

var (
	DefaultClient       *easclient.StoreClient
	DefaultServerClient *easclient.ServerClient
)

// init is run before any tests are executed.
// it loads the environment variables from .env and creates
// a default client for all tests to use.
func init() {
	err := godotenv.Load(".env")
	if errors.Is(err, os.ErrNotExist) {
		return
	}

	client := resty.New()
	client.SetHostURL(fmt.Sprintf("http://%s/eas/archives/%s", os.Getenv("EAS_HOST"), os.Getenv("EAS_STORE")))
	client.SetBasicAuth(os.Getenv("EAS_USER"), os.Getenv("EAS_PASSWORD"))
	DefaultClient = easclient.NewStoreClient(client)

	serverClient := resty.New()
	serverClient.SetHostURL(fmt.Sprintf("http://%s/eas/archives", os.Getenv("EAS_HOST")))
	serverClient.SetBasicAuth(os.Getenv("EAS_USER"), os.Getenv("EAS_PASSWORD"))
	DefaultServerClient = easclient.NewServerClient(serverClient)
}

func testPrelude(t *testing.T) {
	if os.Getenv("GITHUB_ACTION") != "" {
		t.Skip("Tests can't currently be run in the GitHub Action environment. We will first have to make it possible to run the EAS or a mocked variant in CI")
	}
}