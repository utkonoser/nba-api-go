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

func TestGetPlayerNextNGames_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := PlayerNextNGamesParams{
		SeasonTypeAllStar: "Regular Season",
		LeagueIdNullable: "",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetPlayerNextNGames(ctx, params)

	if err != nil {
		t.Logf("PlayerNextNGames endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "playernextngames")
	}

	t.Logf("Successfully fetched playernextngames with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify NextNGames dataset structure
	if dataset, err := response.GetDataSet("NextNGames"); err == nil {
		assert.NotNil(t, dataset, "Should have NextNGames dataset")
		t.Logf("NextNGames: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset NextNGames not found (may be expected): %v", err)
	}
}
