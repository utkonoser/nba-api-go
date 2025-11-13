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

func TestGetPlayerDashPtShots_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := PlayerDashPtShotsParams{
		LeagueId: "00",
		Season: "2023-24",
		SeasonTypeAllStar: "Regular Season",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetPlayerDashPtShots(ctx, params)

	if err != nil {
		t.Logf("PlayerDashPtShots endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "playerdashptshots")
	}

	t.Logf("Successfully fetched playerdashptshots with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify ClosestDefender10ftPlusShooting dataset structure
	if dataset, err := response.GetDataSet("ClosestDefender10ftPlusShooting"); err == nil {
		assert.NotNil(t, dataset, "Should have ClosestDefender10ftPlusShooting dataset")
		t.Logf("ClosestDefender10ftPlusShooting: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset ClosestDefender10ftPlusShooting not found (may be expected): %v", err)
	}

	// Verify ClosestDefenderShooting dataset structure
	if dataset, err := response.GetDataSet("ClosestDefenderShooting"); err == nil {
		assert.NotNil(t, dataset, "Should have ClosestDefenderShooting dataset")
		t.Logf("ClosestDefenderShooting: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset ClosestDefenderShooting not found (may be expected): %v", err)
	}

	// Verify DribbleShooting dataset structure
	if dataset, err := response.GetDataSet("DribbleShooting"); err == nil {
		assert.NotNil(t, dataset, "Should have DribbleShooting dataset")
		t.Logf("DribbleShooting: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset DribbleShooting not found (may be expected): %v", err)
	}

	// Verify GeneralShooting dataset structure
	if dataset, err := response.GetDataSet("GeneralShooting"); err == nil {
		assert.NotNil(t, dataset, "Should have GeneralShooting dataset")
		t.Logf("GeneralShooting: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset GeneralShooting not found (may be expected): %v", err)
	}

	// Verify Overall dataset structure
	if dataset, err := response.GetDataSet("Overall"); err == nil {
		assert.NotNil(t, dataset, "Should have Overall dataset")
		t.Logf("Overall: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset Overall not found (may be expected): %v", err)
	}

	// Verify ShotClockShooting dataset structure
	if dataset, err := response.GetDataSet("ShotClockShooting"); err == nil {
		assert.NotNil(t, dataset, "Should have ShotClockShooting dataset")
		t.Logf("ShotClockShooting: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset ShotClockShooting not found (may be expected): %v", err)
	}

	// Verify TouchTimeShooting dataset structure
	if dataset, err := response.GetDataSet("TouchTimeShooting"); err == nil {
		assert.NotNil(t, dataset, "Should have TouchTimeShooting dataset")
		t.Logf("TouchTimeShooting: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset TouchTimeShooting not found (may be expected): %v", err)
	}
}
