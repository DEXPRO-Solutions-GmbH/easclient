package easclient_test

import (
	"context"
	"encoding/xml"
	"testing"

	"github.com/DEXPRO-Solutions-GmbH/easclient"
	"github.com/DEXPRO-Solutions-GmbH/easclient/eastest"
	"github.com/DEXPRO-Solutions-GmbH/easclient/internal"
	"github.com/stretchr/testify/require"
)

func TestStoreClient_GetStoreStatus(t *testing.T) {
	internal.TestPrelude(t)
	eastest.SkipInCI(t)

	ctx := context.Background()
	user := easclient.NewUserClaims("test@dexpro.de")
	ctx = user.SetOnContext(ctx)

	status, err := eastest.DefaultClient().GetStoreStatus(ctx)
	require.NoError(t, err)
	require.NotNil(t, status)
}

func Test_UnmarshalStoreStatus(t *testing.T) {
	respBody := `<?xml version="1.0" encoding="UTF-8"?>
<status xmlns="http://namespace.otris.de/2010/09/archive" xmlns:xlink="http://www.w3.org/1999/xlink">
    <registry>
        <records>
            <all>34</all>
            <indexed>34</indexed>
        </records>
        <attachments>
            <all>4</all>
            <indexed>4</indexed>
        </attachments>
    </registry>
    <index>
        <documents>38</documents>
        <isCurrent>true</isCurrent>
        <hasDeletions>false</hasDeletions>
    </index>
</status>`

	var resp easclient.StoreStatus

	require.NoError(t, xml.Unmarshal([]byte(respBody), &resp))
	require.Equal(t, 34, resp.Registry.Records.All)
	require.Equal(t, 34, resp.Registry.Records.Indexed)
	require.Equal(t, 38, resp.Index.Documents)
	require.Equal(t, true, resp.Index.IsCurrent)
	require.Equal(t, false, resp.Index.HasDeletions)
}
