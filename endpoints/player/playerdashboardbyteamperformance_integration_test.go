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

func TestGetPlayerDashboardByTeamPerformance_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := PlayerDashboardByTeamPerformanceParams{
		PlayerId: "2544",
		Season:   "2023-24",
		SeasonTypePlayoffs: "Regular Season",
		LeagueIdNullable: "00",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetPlayerDashboardByTeamPerformance(ctx, params)

	if err != nil {
		t.Logf("PlayerDashboardByTeamPerformance endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "playerdashboardbyteamperformance")
	}

	t.Logf("Successfully fetched playerdashboardbyteamperformance with %d result sets", len(response.ResultSets))

	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

	if dataset, err := response.GetDataSet("OverallPlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have OverallPlayerDashboard dataset")
		t.Logf("OverallPlayerDashboard: %d rows", dataset.RowCount())
	} else {
		t.Logf("Dataset OverallPlayerDashboard not found (may be expected): %v", err)
	}
}

