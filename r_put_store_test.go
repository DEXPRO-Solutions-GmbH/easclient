package easclient_test

import (
	"context"
	"testing"

	"github.com/DEXPRO-Solutions-GmbH/easclient"
	"github.com/DEXPRO-Solutions-GmbH/easclient/eastest"
	"github.com/stretchr/testify/require"
)

func TestStoreClient_PutStore(t *testing.T) {
	eastest.SkipInCI(t)

	ctx := context.Background()
	user := easclient.NewUserClaims("test@dexpro.de")
	ctx = user.SetOnContext(ctx)

	storeName := "random-store"

	err := DefaultServerClient.PutStore(ctx, storeName, &easclient.PutStoreRequest{
		ConfigurationTemplate: easclient.ConfigurationTemplate{
			Name: "default",
			Parameters: []easclient.ConfigurationParameter{
				{
					Name:  "STORE_NAME",
					Value: storeName,
				},
			},
		},
	})

	require.NoError(t, err)
}
