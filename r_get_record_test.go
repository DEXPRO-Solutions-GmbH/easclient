package easclient_test

import (
	"context"
	"testing"

	"github.com/DEXPRO-Solutions-GmbH/easclient"
	"github.com/DEXPRO-Solutions-GmbH/easclient/eastest"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestStoreClient_GetRecord(t *testing.T) {
	eastest.SkipInCI(t)

	ctx := context.Background()
	user := easclient.NewUserClaims("test@dexpro.de")
	ctx = user.SetOnContext(ctx)

	record, err := DefaultClient.GetRecord(ctx, uuid.MustParse("990ac5bf-1df5-45b8-82ca-41120621f826"))
	require.NoError(t, err)
	require.NotNil(t, record)
}
