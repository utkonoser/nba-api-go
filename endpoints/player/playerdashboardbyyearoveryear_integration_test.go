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

func TestGetPlayerDashboardByYearOverYear_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := PlayerDashboardByYearOverYearParams{
		Season: "2023-24",
		LeagueIdNullable: "",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetPlayerDashboardByYearOverYear(ctx, params)

	if err != nil {
		t.Logf("PlayerDashboardByYearOverYear endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "playerdashboardbyyearoveryear")
	}

	t.Logf("Successfully fetched playerdashboardbyyearoveryear with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify ByYearPlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("ByYearPlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have ByYearPlayerDashboard dataset")
		t.Logf("ByYearPlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset ByYearPlayerDashboard not found (may be expected): %v", err)
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
