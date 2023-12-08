package easclient_test

import (
	"errors"
	"fmt"
	"os"

	"github.com/DEXPRO-Solutions-GmbH/easclient"
	"github.com/joho/godotenv"
	"gopkg.in/resty.v1"
)

var DefaultClient *easclient.StoreClient

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
}
