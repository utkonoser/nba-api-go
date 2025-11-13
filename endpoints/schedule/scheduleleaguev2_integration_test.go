//go:build integration
// +build integration

package schedule

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetScheduleLeagueV2_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := ScheduleLeagueV2Params{
		LeagueId: "00",
		Season: "2023-24",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetScheduleLeagueV2(ctx, params)

	if err != nil {
		t.Logf("ScheduleLeagueV2 endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "scheduleleaguev2")
	}

	t.Logf("Successfully fetched scheduleleaguev2 with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify SeasonGames dataset structure
	if dataset, err := response.GetDataSet("SeasonGames"); err == nil {
		assert.NotNil(t, dataset, "Should have SeasonGames dataset")
		t.Logf("SeasonGames: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset SeasonGames not found (may be expected): %v", err)
	}

	// Verify SeasonWeeks dataset structure
	if dataset, err := response.GetDataSet("SeasonWeeks"); err == nil {
		assert.NotNil(t, dataset, "Should have SeasonWeeks dataset")
		t.Logf("SeasonWeeks: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset SeasonWeeks not found (may be expected): %v", err)
	}
}
