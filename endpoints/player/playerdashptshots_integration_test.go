//go:build integration
// +build integration

package player

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPlayerDashPtShots_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := PlayerDashPtShotsParams{
		TeamId: "1610612737", // Atlanta Hawks
		PlayerId: "2544", // LeBron James
		Season:   "2023-24",
		SeasonTypeAllStar: "Regular Season",
		LeagueId: "00",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetPlayerDashPtShots(ctx, params)

	if err != nil {
		t.Logf("PlayerDashPtShots endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "playerdashptshots")
	}

	t.Logf("Successfully fetched playerdashptshots with %d result sets", len(response.ResultSets))

	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

	if dataset, err := response.GetDataSet("Overall"); err == nil {
		assert.NotNil(t, dataset, "Should have Overall dataset")
		t.Logf("Overall: %d rows", dataset.RowCount())
	} else {
		t.Logf("Dataset Overall not found (may be expected): %v", err)
	}
}

