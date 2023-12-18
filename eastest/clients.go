package eastest

import (
	"fmt"
	"os"
	"sync"

	"github.com/DEXPRO-Solutions-GmbH/easclient"
	"gopkg.in/resty.v1"
)

var (
	defaultClient       *easclient.StoreClient
	defaultClientOnce   sync.Once
	defaultServerClient *easclient.ServerClient
	defaultServerOnce   sync.Once
)

// DefaultClient returns an a client based on ENV variables: EAS_HOST, EAS_STORE, EAS_USER, EAS_PASSWORD
func DefaultClient() *easclient.StoreClient {
	defaultClientOnce.Do(func() {
		client := resty.New()
		client.SetHostURL(fmt.Sprintf("http://%s/eas/archives/%s", os.Getenv("EAS_HOST"), os.Getenv("EAS_STORE")))
		client.SetBasicAuth(os.Getenv("EAS_USER"), os.Getenv("EAS_PASSWORD"))
		defaultClient = easclient.NewStoreClient(client)
	})

	return defaultClient
}

// DefaultServerClient returns an a client based on ENV variables: EAS_HOST, EAS_USER, EAS_PASSWORD
func DefaultServerClient() *easclient.ServerClient {
	defaultServerOnce.Do(func() {
		serverClient := resty.New()
		serverClient.SetHostURL(fmt.Sprintf("http://%s/eas/archives", os.Getenv("EAS_HOST")))
		serverClient.SetBasicAuth(os.Getenv("EAS_USER"), os.Getenv("EAS_PASSWORD"))
		defaultServerClient = easclient.NewServerClient(serverClient)
	})

	return defaultServerClient
}
