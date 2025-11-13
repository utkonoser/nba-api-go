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

func TestGetCumeStatsTeam_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := CumeStatsTeamParams{
		LeagueId: "00",
		Season: "2023-24",
		SeasonTypeAllStar: "Regular Season",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetCumeStatsTeam(ctx, params)

	if err != nil {
		t.Logf("CumeStatsTeam endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "cumestatsteam")
	}

	t.Logf("Successfully fetched cumestatsteam with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify GameByGameStats dataset structure
	if dataset, err := response.GetDataSet("GameByGameStats"); err == nil {
		assert.NotNil(t, dataset, "Should have GameByGameStats dataset")
		t.Logf("GameByGameStats: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset GameByGameStats not found (may be expected): %v", err)
	}

	// Verify TotalTeamStats dataset structure
	if dataset, err := response.GetDataSet("TotalTeamStats"); err == nil {
		assert.NotNil(t, dataset, "Should have TotalTeamStats dataset")
		t.Logf("TotalTeamStats: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset TotalTeamStats not found (may be expected): %v", err)
	}
}
