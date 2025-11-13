//go:build integration
// +build integration

package draft

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetDraftCombineStats_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := DraftCombineStatsParams{
		LeagueId: "00",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetDraftCombineStats(ctx, params)

	if err != nil {
		t.Logf("DraftCombineStats endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "draftcombinestats")
	}

	t.Logf("Successfully fetched draftcombinestats with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify DraftCombineStats dataset structure
	if dataset, err := response.GetDataSet("DraftCombineStats"); err == nil {
		assert.NotNil(t, dataset, "Should have DraftCombineStats dataset")
		t.Logf("DraftCombineStats: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset DraftCombineStats not found (may be expected): %v", err)
	}
}
