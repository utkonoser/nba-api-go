//go:build integration

package player

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPlayerFantasyProfileBarGraph_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := PlayerFantasyProfileBarGraphParams{
		PlayerId: "2544", // LeBron James
		Season:   "2023-24",
		LeagueIdNullable: "00",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetPlayerFantasyProfileBarGraph(ctx, params)

	if err != nil {
		t.Logf("PlayerFantasyProfileBarGraph endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "playerfantasyprofilebargraph")
	}

	t.Logf("Successfully fetched playerfantasyprofilebargraph with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

	// Verify LastFiveGamesAvg dataset structure
	if dataset, err := response.GetDataSet("LastFiveGamesAvg"); err == nil {
		assert.NotNil(t, dataset, "Should have LastFiveGamesAvg dataset")
		t.Logf("LastFiveGamesAvg: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset LastFiveGamesAvg not found (may be expected): %v", err)
	}

	// Verify SeasonAvg dataset structure
	if dataset, err := response.GetDataSet("SeasonAvg"); err == nil {
		assert.NotNil(t, dataset, "Should have SeasonAvg dataset")
		t.Logf("SeasonAvg: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset SeasonAvg not found (may be expected): %v", err)
	}
}

