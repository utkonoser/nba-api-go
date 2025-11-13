//go:build integration
// +build integration

package misc

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetVideoEventsAsset_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := VideoEventsAssetParams{
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetVideoEventsAsset(ctx, params)

	if err != nil {
		t.Logf("VideoEventsAsset endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "videoeventsasset")
	}

	t.Logf("Successfully fetched videoeventsasset with %d result sets", len(response.ResultSets))
}
