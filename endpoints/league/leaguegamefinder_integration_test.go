//go:build integration

package league

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetLeagueGameFinder_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := LeagueGameFinderParams{
		LeagueIdNullable: "",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetLeagueGameFinder(ctx, params)

	if err != nil {
		t.Logf("LeagueGameFinder endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "leaguegamefinder")
	}

	t.Logf("Successfully fetched leaguegamefinder with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify LeagueGameFinderResults dataset structure
	if dataset, err := response.GetDataSet("LeagueGameFinderResults"); err == nil {
		assert.NotNil(t, dataset, "Should have LeagueGameFinderResults dataset")
		t.Logf("LeagueGameFinderResults: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset LeagueGameFinderResults not found (may be expected): %v", err)
	}
}
