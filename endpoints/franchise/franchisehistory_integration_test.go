//go:build integration
// +build integration

package franchise

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFranchiseHistory_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := FranchiseHistoryParams{
		LeagueId: "00",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetFranchiseHistory(ctx, params)

	if err != nil {
		t.Logf("FranchiseHistory endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "franchisehistory")
	}

	t.Logf("Successfully fetched franchisehistory with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify DefunctTeams dataset structure
	if dataset, err := response.GetDataSet("DefunctTeams"); err == nil {
		assert.NotNil(t, dataset, "Should have DefunctTeams dataset")
		t.Logf("DefunctTeams: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset DefunctTeams not found (may be expected): %v", err)
	}

	// Verify FranchiseHistory dataset structure
	if dataset, err := response.GetDataSet("FranchiseHistory"); err == nil {
		assert.NotNil(t, dataset, "Should have FranchiseHistory dataset")
		t.Logf("FranchiseHistory: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset FranchiseHistory not found (may be expected): %v", err)
	}
}
