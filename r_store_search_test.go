package easclient_test

import (
	"context"
	"testing"

	"github.com/DEXPRO-Solutions-GmbH/easclient"
	"github.com/DEXPRO-Solutions-GmbH/easclient/eastest"
	"github.com/DEXPRO-Solutions-GmbH/easclient/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStoreClient_Search(t *testing.T) {
	internal.TestPrelude(t)
	eastest.SkipInCI(t)

	ctx := context.Background()
	user := easclient.NewUserClaims("test@dexpro.de")
	ctx = user.SetOnContext(ctx)

	request := &easclient.SearchRequest{
		Query: "Amazo*",
	}

	response, err := eastest.DefaultClient().Search(ctx, request)
	require.NoError(t, err)
	require.NotNil(t, response)

	assert.Equal(t, "Amazo*", response.Query)
	assert.Greater(t, response.TotalHits, 0)
	assert.Greater(t, response.EffectiveResults, 0)
}
