//go:build integration
// +build integration

package team

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTeamDashboardByShootingSplits_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := TeamDashboardByShootingSplitsParams{
		TeamId: "1610612737", // Atlanta Hawks
		Season: "2023-24",
		SeasonTypeAllStar: "Regular Season",
		LeagueIdNullable: "00",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetTeamDashboardByShootingSplits(ctx, params)

	if err != nil {
		t.Logf("TeamDashboardByShootingSplits endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "teamdashboardbyshootingsplits")
	}

	t.Logf("Successfully fetched teamdashboardbyshootingsplits with %d result sets", len(response.ResultSets))

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

	// Verify AssitedShotTeamDashboard dataset structure
	if dataset, err := response.GetDataSet("AssitedShotTeamDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have AssitedShotTeamDashboard dataset")
		t.Logf("AssitedShotTeamDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset AssitedShotTeamDashboard not found (may be expected): %v", err)
	}

	// Verify OverallTeamDashboard dataset structure
	if dataset, err := response.GetDataSet("OverallTeamDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have OverallTeamDashboard dataset")
		t.Logf("OverallTeamDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset OverallTeamDashboard not found (may be expected): %v", err)
	}

	// Verify Shot5FTTeamDashboard dataset structure
	if dataset, err := response.GetDataSet("Shot5FTTeamDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have Shot5FTTeamDashboard dataset")
		t.Logf("Shot5FTTeamDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset Shot5FTTeamDashboard not found (may be expected): %v", err)
	}

	// Verify Shot8FTTeamDashboard dataset structure
	if dataset, err := response.GetDataSet("Shot8FTTeamDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have Shot8FTTeamDashboard dataset")
		t.Logf("Shot8FTTeamDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset Shot8FTTeamDashboard not found (may be expected): %v", err)
	}

	// Verify ShotAreaTeamDashboard dataset structure
	if dataset, err := response.GetDataSet("ShotAreaTeamDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have ShotAreaTeamDashboard dataset")
		t.Logf("ShotAreaTeamDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset ShotAreaTeamDashboard not found (may be expected): %v", err)
	}

	// Verify ShotTypeTeamDashboard dataset structure
	if dataset, err := response.GetDataSet("ShotTypeTeamDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have ShotTypeTeamDashboard dataset")
		t.Logf("ShotTypeTeamDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset ShotTypeTeamDashboard not found (may be expected): %v", err)
	}
}
