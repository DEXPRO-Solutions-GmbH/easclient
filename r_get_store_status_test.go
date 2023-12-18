package easclient_test

import (
	"context"
	"testing"

	"github.com/DEXPRO-Solutions-GmbH/easclient"
	"github.com/DEXPRO-Solutions-GmbH/easclient/eastest"
	"github.com/stretchr/testify/require"
)

func TestStoreClient_GetStoreStatus(t *testing.T) {
	eastest.SkipInCI(t)

	ctx := context.Background()
	user := easclient.NewUserClaims("test@dexpro.de")
	ctx = user.SetOnContext(ctx)

	status, err := eastest.DefaultClient.GetStoreStatus(ctx)
	require.NoError(t, err)
	require.NotNil(t, status)
}
