package eastest

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/DEXPRO-Solutions-GmbH/easclient"
	"github.com/joho/godotenv"
	"gopkg.in/resty.v1"
)

var (
	defaultClient       *easclient.StoreClient
	defaultClientOnce   sync.Once
	defaultServerClient *easclient.ServerClient
	defaultServerOnce   sync.Once
)

func DefaultClient() *easclient.StoreClient {
	defaultClientOnce.Do(func() {
		client := resty.New()
		client.SetHostURL(fmt.Sprintf("http://%s/eas/archives/%s", os.Getenv("EAS_HOST"), os.Getenv("EAS_STORE")))
		client.SetBasicAuth(os.Getenv("EAS_USER"), os.Getenv("EAS_PASSWORD"))
		defaultClient = easclient.NewStoreClient(client)
	})

	return defaultClient
}

func DefaultServerClient() *easclient.ServerClient {
	defaultServerOnce.Do(func() {
		serverClient := resty.New()
		serverClient.SetHostURL(fmt.Sprintf("http://%s/eas/archives", os.Getenv("EAS_HOST")))
		serverClient.SetBasicAuth(os.Getenv("EAS_USER"), os.Getenv("EAS_PASSWORD"))
		defaultServerClient = easclient.NewServerClient(serverClient)
	})

	return defaultServerClient
}

// init is run before any tests are executed.
// it loads the environment variables from .env and creates
// a default client for all tests to use.
func init() {
	err := godotenv.Load(".env")
	if errors.Is(err, os.ErrNotExist) {
		return
	}
}
