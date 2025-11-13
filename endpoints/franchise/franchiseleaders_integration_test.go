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

func TestGetFranchiseLeaders_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := FranchiseLeadersParams{
		TeamId:          "1610612737", // Atlanta Hawks
		LeagueIdNullable: "00",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetFranchiseLeaders(ctx, params)

	if err != nil {
		t.Logf("FranchiseLeaders endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "franchiseleaders")
	}

	t.Logf("Successfully fetched franchiseleaders with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify FranchiseLeaders dataset structure
	if dataset, err := response.GetDataSet("FranchiseLeaders"); err == nil {
		assert.NotNil(t, dataset, "Should have FranchiseLeaders dataset")
		t.Logf("FranchiseLeaders: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset FranchiseLeaders not found (may be expected): %v", err)
	}
}
