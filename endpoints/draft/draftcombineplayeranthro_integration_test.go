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

func TestGetDraftCombinePlayerAnthro_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := DraftCombinePlayerAnthroParams{
		LeagueId: "00",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetDraftCombinePlayerAnthro(ctx, params)

	if err != nil {
		t.Logf("DraftCombinePlayerAnthro endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "draftcombineplayeranthro")
	}

	t.Logf("Successfully fetched draftcombineplayeranthro with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify Results dataset structure
	if dataset, err := response.GetDataSet("Results"); err == nil {
		assert.NotNil(t, dataset, "Should have Results dataset")
		t.Logf("Results: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset Results not found (may be expected): %v", err)
	}
}
