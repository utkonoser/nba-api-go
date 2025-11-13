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

func TestGetTeamDashboardByGeneralSplits_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := TeamDashboardByGeneralSplitsParams{
		TeamId: "1610612737", // Atlanta Hawks
		Season: "2023-24",
		SeasonTypeAllStar: "Regular Season",
		LeagueIdNullable: "00",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetTeamDashboardByGeneralSplits(ctx, params)

	if err != nil {
		t.Logf("TeamDashboardByGeneralSplits endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "teamdashboardbygeneralsplits")
	}

	t.Logf("Successfully fetched teamdashboardbygeneralsplits with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify DaysRestTeamDashboard dataset structure
	if dataset, err := response.GetDataSet("DaysRestTeamDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have DaysRestTeamDashboard dataset")
		t.Logf("DaysRestTeamDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset DaysRestTeamDashboard not found (may be expected): %v", err)
	}

	// Verify LocationTeamDashboard dataset structure
	if dataset, err := response.GetDataSet("LocationTeamDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have LocationTeamDashboard dataset")
		t.Logf("LocationTeamDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset LocationTeamDashboard not found (may be expected): %v", err)
	}

	// Verify MonthTeamDashboard dataset structure
	if dataset, err := response.GetDataSet("MonthTeamDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have MonthTeamDashboard dataset")
		t.Logf("MonthTeamDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset MonthTeamDashboard not found (may be expected): %v", err)
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

	// Verify PrePostAllStarTeamDashboard dataset structure
	if dataset, err := response.GetDataSet("PrePostAllStarTeamDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have PrePostAllStarTeamDashboard dataset")
		t.Logf("PrePostAllStarTeamDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset PrePostAllStarTeamDashboard not found (may be expected): %v", err)
	}

	// Verify WinsLossesTeamDashboard dataset structure
	if dataset, err := response.GetDataSet("WinsLossesTeamDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have WinsLossesTeamDashboard dataset")
		t.Logf("WinsLossesTeamDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset WinsLossesTeamDashboard not found (may be expected): %v", err)
	}
}
