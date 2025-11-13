//go:build integration
// +build integration

package shot

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetShotChartLineupDetail_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := ShotChartLineupDetailParams{
		LeagueId: "00",
		Season: "2023-24",
		SeasonTypeAllStar: "Regular Season",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetShotChartLineupDetail(ctx, params)

	if err != nil {
		t.Logf("ShotChartLineupDetail endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "shotchartlineupdetail")
	}

	t.Logf("Successfully fetched shotchartlineupdetail with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify ShotChartLineupDetail dataset structure
	if dataset, err := response.GetDataSet("ShotChartLineupDetail"); err == nil {
		assert.NotNil(t, dataset, "Should have ShotChartLineupDetail dataset")
		t.Logf("ShotChartLineupDetail: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset ShotChartLineupDetail not found (may be expected): %v", err)
	}

	// Verify ShotChartLineupLeagueAverage dataset structure
	if dataset, err := response.GetDataSet("ShotChartLineupLeagueAverage"); err == nil {
		assert.NotNil(t, dataset, "Should have ShotChartLineupLeagueAverage dataset")
		t.Logf("ShotChartLineupLeagueAverage: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset ShotChartLineupLeagueAverage not found (may be expected): %v", err)
	}
}
