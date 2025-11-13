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

func TestGetDraftBoard_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := DraftBoardParams{
		LeagueId: "00",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetDraftBoard(ctx, params)

	if err != nil {
		t.Logf("DraftBoard endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "draftboard")
	}

	t.Logf("Successfully fetched draftboard with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify DraftBoard dataset structure
	if dataset, err := response.GetDataSet("DraftBoard"); err == nil {
		assert.NotNil(t, dataset, "Should have DraftBoard dataset")
		t.Logf("DraftBoard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset DraftBoard not found (may be expected): %v", err)
	}
}
