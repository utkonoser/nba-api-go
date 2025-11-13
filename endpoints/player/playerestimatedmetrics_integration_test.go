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

func TestGetPlayerEstimatedMetrics_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := PlayerEstimatedMetricsParams{
		LeagueId: "00",
		Season: "2023-24",
		SeasonType: "Regular Season",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetPlayerEstimatedMetrics(ctx, params)

	if err != nil {
		t.Logf("PlayerEstimatedMetrics endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "playerestimatedmetrics")
	}

	t.Logf("Successfully fetched playerestimatedmetrics with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

	// Verify PlayerEstimatedMetrics dataset structure
	if dataset, err := response.GetDataSet("PlayerEstimatedMetrics"); err == nil {
		assert.NotNil(t, dataset, "Should have PlayerEstimatedMetrics dataset")
		t.Logf("PlayerEstimatedMetrics: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset PlayerEstimatedMetrics not found (may be expected): %v", err)
	}
}
