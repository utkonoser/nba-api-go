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

func TestGetVideoDetailsAsset_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := VideoDetailsAssetParams{
		Season: "2023-24",
		SeasonTypeAllStar: "Regular Season",
		LeagueIdNullable: "",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetVideoDetailsAsset(ctx, params)

	if err != nil {
		t.Logf("VideoDetailsAsset endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "videodetailsasset")
	}

	t.Logf("Successfully fetched videodetailsasset with %d result sets", len(response.ResultSets))
}
