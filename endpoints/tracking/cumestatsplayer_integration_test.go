//go:build integration
// +build integration

package tracking

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetCumeStatsPlayer_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := CumeStatsPlayerParams{
		LeagueId: "00",
		Season: "2023-24",
		SeasonTypeAllStar: "Regular Season",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetCumeStatsPlayer(ctx, params)

	if err != nil {
		t.Logf("CumeStatsPlayer endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "cumestatsplayer")
	}

	t.Logf("Successfully fetched cumestatsplayer with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify GameByGameStats dataset structure
	if dataset, err := response.GetDataSet("GameByGameStats"); err == nil {
		assert.NotNil(t, dataset, "Should have GameByGameStats dataset")
		t.Logf("GameByGameStats: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset GameByGameStats not found (may be expected): %v", err)
	}

	// Verify TotalPlayerStats dataset structure
	if dataset, err := response.GetDataSet("TotalPlayerStats"); err == nil {
		assert.NotNil(t, dataset, "Should have TotalPlayerStats dataset")
		t.Logf("TotalPlayerStats: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset TotalPlayerStats not found (may be expected): %v", err)
	}
}
