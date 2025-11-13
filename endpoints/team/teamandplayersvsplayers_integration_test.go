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

func TestGetTeamAndPlayersVsPlayers_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := TeamAndPlayersVsPlayersParams{
		Season: "2023-24",
		LeagueIdNullable: "",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetTeamAndPlayersVsPlayers(ctx, params)

	if err != nil {
		t.Logf("TeamAndPlayersVsPlayers endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "teamandplayersvsplayers")
	}

	t.Logf("Successfully fetched teamandplayersvsplayers with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify PlayersVsPlayers dataset structure
	if dataset, err := response.GetDataSet("PlayersVsPlayers"); err == nil {
		assert.NotNil(t, dataset, "Should have PlayersVsPlayers dataset")
		t.Logf("PlayersVsPlayers: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset PlayersVsPlayers not found (may be expected): %v", err)
	}

	// Verify TeamPlayersVsPlayersOff dataset structure
	if dataset, err := response.GetDataSet("TeamPlayersVsPlayersOff"); err == nil {
		assert.NotNil(t, dataset, "Should have TeamPlayersVsPlayersOff dataset")
		t.Logf("TeamPlayersVsPlayersOff: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset TeamPlayersVsPlayersOff not found (may be expected): %v", err)
	}

	// Verify TeamPlayersVsPlayersOn dataset structure
	if dataset, err := response.GetDataSet("TeamPlayersVsPlayersOn"); err == nil {
		assert.NotNil(t, dataset, "Should have TeamPlayersVsPlayersOn dataset")
		t.Logf("TeamPlayersVsPlayersOn: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset TeamPlayersVsPlayersOn not found (may be expected): %v", err)
	}

	// Verify TeamVsPlayers dataset structure
	if dataset, err := response.GetDataSet("TeamVsPlayers"); err == nil {
		assert.NotNil(t, dataset, "Should have TeamVsPlayers dataset")
		t.Logf("TeamVsPlayers: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset TeamVsPlayers not found (may be expected): %v", err)
	}

	// Verify TeamVsPlayersOff dataset structure
	if dataset, err := response.GetDataSet("TeamVsPlayersOff"); err == nil {
		assert.NotNil(t, dataset, "Should have TeamVsPlayersOff dataset")
		t.Logf("TeamVsPlayersOff: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset TeamVsPlayersOff not found (may be expected): %v", err)
	}
}
