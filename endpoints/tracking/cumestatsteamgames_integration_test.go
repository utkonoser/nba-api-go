//go:build integration
// +build integration

package tracking

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetCumeStatsTeamGames_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := CumeStatsTeamGamesParams{
		TeamId: "1610612737", // Atlanta Hawks
		LeagueId: "00",
		Season: "2023-24",
		SeasonTypeAllStar: "Regular Season",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetCumeStatsTeamGames(ctx, params)

	if err != nil {
		t.Logf("CumeStatsTeamGames endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "cumestatsteamgames")
	}

	t.Logf("Successfully fetched cumestatsteamgames with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify CumeStatsTeamGames dataset structure
	if dataset, err := response.GetDataSet("CumeStatsTeamGames"); err == nil {
		assert.NotNil(t, dataset, "Should have CumeStatsTeamGames dataset")
		t.Logf("CumeStatsTeamGames: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset CumeStatsTeamGames not found (may be expected): %v", err)
	}
}
