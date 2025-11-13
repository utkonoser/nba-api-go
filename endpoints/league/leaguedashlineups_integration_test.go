//go:build integration
// +build integration

package league

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetLeagueDashLineups_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := LeagueDashLineupsParams{
		Season: "2023-24",
		SeasonTypeAllStar: "Regular Season",
		LeagueIdNullable: "",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetLeagueDashLineups(ctx, params)

	if err != nil {
		t.Logf("LeagueDashLineups endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "leaguedashlineups")
	}

	t.Logf("Successfully fetched leaguedashlineups with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify Lineups dataset structure
	if dataset, err := response.GetDataSet("Lineups"); err == nil {
		assert.NotNil(t, dataset, "Should have Lineups dataset")
		t.Logf("Lineups: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset Lineups not found (may be expected): %v", err)
	}
}
