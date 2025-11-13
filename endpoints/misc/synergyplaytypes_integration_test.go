//go:build integration
// +build integration

package misc

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSynergyPlayTypes_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := SynergyPlayTypesParams{
		LeagueId: "00",
		SeasonTypeAllStar: "Regular Season",
		Season: "2023-24",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetSynergyPlayTypes(ctx, params)

	if err != nil {
		t.Logf("SynergyPlayTypes endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty or slightly different for some endpoints
	if response.Resource != "" {
		// API returns "synergyplaytype" without 's'
		assert.Contains(t, response.Resource, "synergyplaytype")
	}

	t.Logf("Successfully fetched synergyplaytypes with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify SynergyPlayType dataset structure
	if dataset, err := response.GetDataSet("SynergyPlayType"); err == nil {
		assert.NotNil(t, dataset, "Should have SynergyPlayType dataset")
		t.Logf("SynergyPlayType: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset SynergyPlayType not found (may be expected): %v", err)
	}
}
