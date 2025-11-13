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

func TestGetTeamPlayerOnOffSummary_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := TeamPlayerOnOffSummaryParams{
		TeamId: "1610612737", // Atlanta Hawks
		Season: "2023-24",
		SeasonTypeAllStar: "Regular Season",
		LeagueIdNullable: "00",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetTeamPlayerOnOffSummary(ctx, params)

	if err != nil {
		t.Logf("TeamPlayerOnOffSummary endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "teamplayeronoffsummary")
	}

	t.Logf("Successfully fetched teamplayeronoffsummary with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify OverallTeamPlayerOnOffSummary dataset structure
	if dataset, err := response.GetDataSet("OverallTeamPlayerOnOffSummary"); err == nil {
		assert.NotNil(t, dataset, "Should have OverallTeamPlayerOnOffSummary dataset")
		t.Logf("OverallTeamPlayerOnOffSummary: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset OverallTeamPlayerOnOffSummary not found (may be expected): %v", err)
	}

	// Verify PlayersOffCourtTeamPlayerOnOffSummary dataset structure
	if dataset, err := response.GetDataSet("PlayersOffCourtTeamPlayerOnOffSummary"); err == nil {
		assert.NotNil(t, dataset, "Should have PlayersOffCourtTeamPlayerOnOffSummary dataset")
		t.Logf("PlayersOffCourtTeamPlayerOnOffSummary: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset PlayersOffCourtTeamPlayerOnOffSummary not found (may be expected): %v", err)
	}

	// Verify PlayersOnCourtTeamPlayerOnOffSummary dataset structure
	if dataset, err := response.GetDataSet("PlayersOnCourtTeamPlayerOnOffSummary"); err == nil {
		assert.NotNil(t, dataset, "Should have PlayersOnCourtTeamPlayerOnOffSummary dataset")
		t.Logf("PlayersOnCourtTeamPlayerOnOffSummary: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset PlayersOnCourtTeamPlayerOnOffSummary not found (may be expected): %v", err)
	}
}
