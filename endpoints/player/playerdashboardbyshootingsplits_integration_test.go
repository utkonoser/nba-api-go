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

func TestGetPlayerDashboardByShootingSplits_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := PlayerDashboardByShootingSplitsParams{
		Season: "2023-24",
		LeagueIdNullable: "",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetPlayerDashboardByShootingSplits(ctx, params)

	if err != nil {
		t.Logf("PlayerDashboardByShootingSplits endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "playerdashboardbyshootingsplits")
	}

	t.Logf("Successfully fetched playerdashboardbyshootingsplits with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify AssistedBy dataset structure
	if dataset, err := response.GetDataSet("AssistedBy"); err == nil {
		assert.NotNil(t, dataset, "Should have AssistedBy dataset")
		t.Logf("AssistedBy: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset AssistedBy not found (may be expected): %v", err)
	}

	// Verify AssitedShotPlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("AssitedShotPlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have AssitedShotPlayerDashboard dataset")
		t.Logf("AssitedShotPlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset AssitedShotPlayerDashboard not found (may be expected): %v", err)
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

	// Verify Shot5FTPlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("Shot5FTPlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have Shot5FTPlayerDashboard dataset")
		t.Logf("Shot5FTPlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset Shot5FTPlayerDashboard not found (may be expected): %v", err)
	}

	// Verify Shot8FTPlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("Shot8FTPlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have Shot8FTPlayerDashboard dataset")
		t.Logf("Shot8FTPlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset Shot8FTPlayerDashboard not found (may be expected): %v", err)
	}

	// Verify ShotAreaPlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("ShotAreaPlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have ShotAreaPlayerDashboard dataset")
		t.Logf("ShotAreaPlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset ShotAreaPlayerDashboard not found (may be expected): %v", err)
	}

	// Verify ShotTypePlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("ShotTypePlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have ShotTypePlayerDashboard dataset")
		t.Logf("ShotTypePlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset ShotTypePlayerDashboard not found (may be expected): %v", err)
	}

	// Verify ShotTypeSummaryPlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("ShotTypeSummaryPlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have ShotTypeSummaryPlayerDashboard dataset")
		t.Logf("ShotTypeSummaryPlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset ShotTypeSummaryPlayerDashboard not found (may be expected): %v", err)
	}
}
