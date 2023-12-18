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

	t.Run("returns results when using only query", func(t *testing.T) {
		request := &easclient.SearchRequest{
			Query: "Amazo*",
		}

		response, err := eastest.DefaultClient().Search(ctx, request)
		require.NoError(t, err)
		require.NotNil(t, response)

		assert.Equal(t, "Amazo*", response.Query)
		assert.Greater(t, response.TotalHits, 0)
		assert.Greater(t, response.EffectiveResults, 0)
	})

	t.Run("returns results when using pagination details", func(t *testing.T) {
		request := &easclient.SearchRequest{
			Query:        "Amazo*",
			ItemsPerPage: 25,
			StartIndex:   2500,
		}

		response, err := eastest.DefaultClient().Search(ctx, request)
		require.NoError(t, err)
		require.NotNil(t, response)

		assert.Equal(t, "Amazo*", response.Query)
		assert.Greater(t, response.TotalHits, 0)
		assert.Greater(t, response.EffectiveResults, 0)
	})
}

func TestSearchRequestFromURL(t *testing.T) {
	assertValidRequest := func(req *easclient.SearchRequest) {
		assert.Equal(t, "creditor:amaz*", req.Query)
		assert.Equal(t, 25, req.ItemsPerPage)
		assert.Equal(t, 2500, req.StartIndex)
		assert.Equal(t, "creditor", req.Sort)
		assert.Equal(t, "asc", req.SortOrder)
	}

	t.Run("returns proper result for full url", func(t *testing.T) {
		u, err := easclient.SearchRequestFromURL("https://localhost/eas/archives/stores/store42/?query=creditor:amaz*&itemsPerPage=25&startIndex=2500&sort=creditor&sortOrder=asc")
		require.NoError(t, err)
		assertValidRequest(u)
	})

	t.Run("returns proper result for path only input", func(t *testing.T) {
		u, err := easclient.SearchRequestFromURL("/eas/archives/stores/store42/?query=creditor:amaz*&itemsPerPage=25&startIndex=2500&sort=creditor&sortOrder=asc")
		require.NoError(t, err)
		assertValidRequest(u)
	})

	t.Run("returns proper result for query only input", func(t *testing.T) {
		u, err := easclient.SearchRequestFromURL("?query=creditor:amaz*&itemsPerPage=25&startIndex=2500&sort=creditor&sortOrder=asc")
		require.NoError(t, err)
		assertValidRequest(u)
	})
}
