package easclient_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/DEXPRO-Solutions-GmbH/easclient"
	"github.com/DEXPRO-Solutions-GmbH/easclient/eastest"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestStoreClient_GetRecordAttachment(t *testing.T) {
	eastest.SkipInCI(t)

	ctx := context.Background()
	user := easclient.NewUserClaims("test@dexpro.de")
	ctx = user.SetOnContext(ctx)

	buffer := new(bytes.Buffer)

	record, err := eastest.DefaultClient.GetRecordAttachment(
		ctx,
		buffer,
		uuid.MustParse("a65efcf9-8c74-4b84-8106-233c1c64a07c"),
		uuid.MustParse("8cc83908-0590-4e43-9e6e-d676e00ce41f"),
	)
	require.NoError(t, err)
	require.NotNil(t, record)
	require.Greater(t, buffer.Len(), 1000)
}
