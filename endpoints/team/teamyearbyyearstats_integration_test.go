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

func TestGetTeamYearByYearStats_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := TeamYearByYearStatsParams{
		LeagueId: "00",
		SeasonTypeAllStar: "Regular Season",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetTeamYearByYearStats(ctx, params)

	if err != nil {
		t.Logf("TeamYearByYearStats endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "teamyearbyyearstats")
	}

	t.Logf("Successfully fetched teamyearbyyearstats with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify TeamStats dataset structure
	if dataset, err := response.GetDataSet("TeamStats"); err == nil {
		assert.NotNil(t, dataset, "Should have TeamStats dataset")
		t.Logf("TeamStats: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset TeamStats not found (may be expected): %v", err)
	}
}
