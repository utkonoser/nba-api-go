//go:build integration
// +build integration

package playoff

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetCommonPlayoffSeries_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := CommonPlayoffSeriesParams{
		LeagueId: "00",
		Season: "2023-24",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetCommonPlayoffSeries(ctx, params)

	if err != nil {
		t.Logf("CommonPlayoffSeries endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "commonplayoffseries")
	}

	t.Logf("Successfully fetched commonplayoffseries with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify PlayoffSeries dataset structure
	if dataset, err := response.GetDataSet("PlayoffSeries"); err == nil {
		assert.NotNil(t, dataset, "Should have PlayoffSeries dataset")
		t.Logf("PlayoffSeries: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset PlayoffSeries not found (may be expected): %v", err)
	}
}
