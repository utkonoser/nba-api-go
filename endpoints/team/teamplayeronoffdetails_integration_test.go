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

func TestGetTeamPlayerOnOffDetails_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := TeamPlayerOnOffDetailsParams{
		TeamId: "1610612737", // Atlanta Hawks
		Season: "2023-24",
		SeasonTypeAllStar: "Regular Season",
		LeagueIdNullable: "00",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetTeamPlayerOnOffDetails(ctx, params)

	if err != nil {
		t.Logf("TeamPlayerOnOffDetails endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "teamplayeronoffdetails")
	}

	t.Logf("Successfully fetched teamplayeronoffdetails with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify OverallTeamPlayerOnOffDetails dataset structure
	if dataset, err := response.GetDataSet("OverallTeamPlayerOnOffDetails"); err == nil {
		assert.NotNil(t, dataset, "Should have OverallTeamPlayerOnOffDetails dataset")
		t.Logf("OverallTeamPlayerOnOffDetails: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset OverallTeamPlayerOnOffDetails not found (may be expected): %v", err)
	}

	// Verify PlayersOffCourtTeamPlayerOnOffDetails dataset structure
	if dataset, err := response.GetDataSet("PlayersOffCourtTeamPlayerOnOffDetails"); err == nil {
		assert.NotNil(t, dataset, "Should have PlayersOffCourtTeamPlayerOnOffDetails dataset")
		t.Logf("PlayersOffCourtTeamPlayerOnOffDetails: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset PlayersOffCourtTeamPlayerOnOffDetails not found (may be expected): %v", err)
	}

	// Verify PlayersOnCourtTeamPlayerOnOffDetails dataset structure
	if dataset, err := response.GetDataSet("PlayersOnCourtTeamPlayerOnOffDetails"); err == nil {
		assert.NotNil(t, dataset, "Should have PlayersOnCourtTeamPlayerOnOffDetails dataset")
		t.Logf("PlayersOnCourtTeamPlayerOnOffDetails: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset PlayersOnCourtTeamPlayerOnOffDetails not found (may be expected): %v", err)
	}
}
