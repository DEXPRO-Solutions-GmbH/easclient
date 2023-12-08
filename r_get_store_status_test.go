package easclient_test

import (
	"context"
	"testing"

	"github.com/DEXPRO-Solutions-GmbH/easclient"
	"github.com/stretchr/testify/require"
)

func TestStoreClient_GetStoreStatus(t *testing.T) {
	ctx := context.Background()
	user := easclient.NewUserClaims("test@dexpro.de")
	ctx = user.SetOnContext(ctx)

	status, err := DefaultClient.GetStoreStatus(ctx)
	require.NoError(t, err)
	require.NotNil(t, status)
}
