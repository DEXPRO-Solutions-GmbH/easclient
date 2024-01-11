package easclient_test

import (
	"context"
	"testing"

	"github.com/DEXPRO-Solutions-GmbH/easclient"
	"github.com/DEXPRO-Solutions-GmbH/easclient/eastest"
	"github.com/DEXPRO-Solutions-GmbH/easclient/internal"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestStoreClient_PostRecord(t *testing.T) {
	internal.TestPrelude(t)
	eastest.SkipInCI(t)

	ctx := context.Background()
	user := easclient.NewUserClaims("test@dexpro.de")
	ctx = user.SetOnContext(ctx)

	res, err := eastest.DefaultClient().PostRecord(ctx, &easclient.RecordRequest{
		Fields: map[string]string{
			"Creditor": "DE123456789",
			"Debitor":  "DE987654321",
		},
	})

	require.NoError(t, err)
	require.NotNil(t, res)

	require.Len(t, res.Records, 1)
	require.NoError(t, err)
	require.NotEqual(t, uuid.Nil, res.Records[0].Id)
}
