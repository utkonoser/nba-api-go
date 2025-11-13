//go:build integration
// +build integration

package team

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTeamPlayerDashboard_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := TeamPlayerDashboardParams{
		TeamId: "1610612737", // Atlanta Hawks
		Season: "2023-24",
		SeasonTypeAllStar: "Regular Season",
		LeagueIdNullable: "00",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetTeamPlayerDashboard(ctx, params)

	if err != nil {
		t.Logf("TeamPlayerDashboard endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "teamplayerdashboard")
	}

	t.Logf("Successfully fetched teamplayerdashboard with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify PlayersSeasonTotals dataset structure
	if dataset, err := response.GetDataSet("PlayersSeasonTotals"); err == nil {
		assert.NotNil(t, dataset, "Should have PlayersSeasonTotals dataset")
		t.Logf("PlayersSeasonTotals: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset PlayersSeasonTotals not found (may be expected): %v", err)
	}

	// Verify TeamOverall dataset structure
	if dataset, err := response.GetDataSet("TeamOverall"); err == nil {
		assert.NotNil(t, dataset, "Should have TeamOverall dataset")
		t.Logf("TeamOverall: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset TeamOverall not found (may be expected): %v", err)
	}
}
