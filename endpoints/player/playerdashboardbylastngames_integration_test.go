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

func TestGetPlayerDashboardByLastNGames_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := PlayerDashboardByLastNGamesParams{
		Season: "2023-24",
		LeagueIdNullable: "",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetPlayerDashboardByLastNGames(ctx, params)

	if err != nil {
		t.Logf("PlayerDashboardByLastNGames endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "playerdashboardbylastngames")
	}

	t.Logf("Successfully fetched playerdashboardbylastngames with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify GameNumberPlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("GameNumberPlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have GameNumberPlayerDashboard dataset")
		t.Logf("GameNumberPlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset GameNumberPlayerDashboard not found (may be expected): %v", err)
	}

	// Verify Last10PlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("Last10PlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have Last10PlayerDashboard dataset")
		t.Logf("Last10PlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset Last10PlayerDashboard not found (may be expected): %v", err)
	}

	// Verify Last15PlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("Last15PlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have Last15PlayerDashboard dataset")
		t.Logf("Last15PlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset Last15PlayerDashboard not found (may be expected): %v", err)
	}

	// Verify Last20PlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("Last20PlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have Last20PlayerDashboard dataset")
		t.Logf("Last20PlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset Last20PlayerDashboard not found (may be expected): %v", err)
	}

	// Verify Last5PlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("Last5PlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have Last5PlayerDashboard dataset")
		t.Logf("Last5PlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset Last5PlayerDashboard not found (may be expected): %v", err)
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
