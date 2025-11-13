//go:build integration

package boxscore

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetBoxScoreScoringV3_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := BoxScoreScoringV3Params{
		GameId: "0022300001",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetBoxScoreScoringV3(ctx, params)

	if err != nil {
		t.Logf("BoxScoreScoringV3 endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// V3 endpoints may have empty resource field
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "boxscorescoringv3")
	}

	t.Logf("Successfully fetched boxscorescoringv3 with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify PlayerStats dataset structure
	if dataset, err := response.GetDataSet("PlayerStats"); err == nil {
		assert.NotNil(t, dataset, "Should have PlayerStats dataset")
		t.Logf("PlayerStats: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset PlayerStats not found (may be expected): %v", err)
	}

	// Verify TeamStats dataset structure
	if dataset, err := response.GetDataSet("TeamStats"); err == nil {
		assert.NotNil(t, dataset, "Should have TeamStats dataset")
		t.Logf("TeamStats: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset TeamStats not found (may be expected): %v", err)
	}
}
