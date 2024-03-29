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
		Title: "This is my archive file",
		Fields: map[string]string{
			"Creditor": "DE123456789",
			"Debitor":  "DE987654321",
		},
	})

	require.NoError(t, err)
	require.NotNil(t, res)
	require.NotEqual(t, uuid.Nil, res.ID.Value)

	// Retrieve the record
	record, err := eastest.DefaultClient().GetRecord(ctx, res.ID.Value)
	require.NoError(t, err)

	require.Equal(t, res.ID.Value, record.ID)
	require.Equal(t, "This is my archive file", record.Title)

	require.Equal(t, "DE123456789", record.GetHeaderFieldVal("Creditor"))
	require.Equal(t, "DE987654321", record.GetHeaderFieldVal("Debitor"))
}
