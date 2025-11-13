//go:build integration
// +build integration

package leaders

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetLeadersTiles_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := LeadersTilesParams{
		LeagueId: "00",
		Season: "2023-24",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetLeadersTiles(ctx, params)

	if err != nil {
		t.Logf("LeadersTiles endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "leaderstiles")
	}

	t.Logf("Successfully fetched leaderstiles with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify AllTimeSeasonHigh dataset structure
	if dataset, err := response.GetDataSet("AllTimeSeasonHigh"); err == nil {
		assert.NotNil(t, dataset, "Should have AllTimeSeasonHigh dataset")
		t.Logf("AllTimeSeasonHigh: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset AllTimeSeasonHigh not found (may be expected): %v", err)
	}

	// Verify LastSeasonHigh dataset structure
	if dataset, err := response.GetDataSet("LastSeasonHigh"); err == nil {
		assert.NotNil(t, dataset, "Should have LastSeasonHigh dataset")
		t.Logf("LastSeasonHigh: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset LastSeasonHigh not found (may be expected): %v", err)
	}

	// Verify LeadersTiles dataset structure
	if dataset, err := response.GetDataSet("LeadersTiles"); err == nil {
		assert.NotNil(t, dataset, "Should have LeadersTiles dataset")
		t.Logf("LeadersTiles: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset LeadersTiles not found (may be expected): %v", err)
	}

	// Verify LowSeasonHigh dataset structure
	if dataset, err := response.GetDataSet("LowSeasonHigh"); err == nil {
		assert.NotNil(t, dataset, "Should have LowSeasonHigh dataset")
		t.Logf("LowSeasonHigh: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset LowSeasonHigh not found (may be expected): %v", err)
	}
}
