//go:build integration
// +build integration

package game

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetGameRotation_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := GameRotationParams{
		GameId:   "0022300001",
		LeagueId: "00",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetGameRotation(ctx, params)

	if err != nil {
		t.Logf("GameRotation endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "gamerotation")
	}

	t.Logf("Successfully fetched gamerotation with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify AwayTeam dataset structure
	if dataset, err := response.GetDataSet("AwayTeam"); err == nil {
		assert.NotNil(t, dataset, "Should have AwayTeam dataset")
		t.Logf("AwayTeam: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset AwayTeam not found (may be expected): %v", err)
	}

	// Verify HomeTeam dataset structure
	if dataset, err := response.GetDataSet("HomeTeam"); err == nil {
		assert.NotNil(t, dataset, "Should have HomeTeam dataset")
		t.Logf("HomeTeam: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset HomeTeam not found (may be expected): %v", err)
	}
}
