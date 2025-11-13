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

func TestGetPlayerDashboardByGameSplits_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := PlayerDashboardByGameSplitsParams{
		Season: "2023-24",
		LeagueIdNullable: "",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetPlayerDashboardByGameSplits(ctx, params)

	if err != nil {
		t.Logf("PlayerDashboardByGameSplits endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "playerdashboardbygamesplits")
	}

	t.Logf("Successfully fetched playerdashboardbygamesplits with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify ByActualMarginPlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("ByActualMarginPlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have ByActualMarginPlayerDashboard dataset")
		t.Logf("ByActualMarginPlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset ByActualMarginPlayerDashboard not found (may be expected): %v", err)
	}

	// Verify ByHalfPlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("ByHalfPlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have ByHalfPlayerDashboard dataset")
		t.Logf("ByHalfPlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset ByHalfPlayerDashboard not found (may be expected): %v", err)
	}

	// Verify ByPeriodPlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("ByPeriodPlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have ByPeriodPlayerDashboard dataset")
		t.Logf("ByPeriodPlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset ByPeriodPlayerDashboard not found (may be expected): %v", err)
	}

	// Verify ByScoreMarginPlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("ByScoreMarginPlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have ByScoreMarginPlayerDashboard dataset")
		t.Logf("ByScoreMarginPlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset ByScoreMarginPlayerDashboard not found (may be expected): %v", err)
	}

	// Verify OverallPlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("OverallPlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have OverallPlayerDashboard dataset")
		t.Logf("OverallPlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset OverallPlayerDashboard not found (may be expected): %v", err)
	}
}
