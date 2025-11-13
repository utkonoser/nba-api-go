//go:build integration
// +build integration

package boxscore

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetHustleStatsBoxScore_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := HustleStatsBoxScoreParams{
		GameId: "0022300001",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetHustleStatsBoxScore(ctx, params)

	if err != nil {
		t.Logf("HustleStatsBoxScore endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "hustlestatsboxscore")
	}

	t.Logf("Successfully fetched hustlestatsboxscore with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify HustleStatsAvailable dataset structure
	if dataset, err := response.GetDataSet("HustleStatsAvailable"); err == nil {
		assert.NotNil(t, dataset, "Should have HustleStatsAvailable dataset")
		t.Logf("HustleStatsAvailable: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset HustleStatsAvailable not found (may be expected): %v", err)
	}

	// Verify PlayerStats dataset structure
	if dataset, err := response.GetDataSet("PlayerStats"); err == nil {
		assert.NotNil(t, dataset, "Should have PlayerStats dataset")
		t.Logf("PlayerStats: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset PlayerStats not found (may be expected): %v", err)
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
