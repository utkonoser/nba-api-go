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

func TestGetWinProbabilityPBP_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := WinProbabilityPBPParams{
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetWinProbabilityPBP(ctx, params)

	if err != nil {
		t.Logf("WinProbabilityPBP endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "winprobabilitypbp")
	}

	t.Logf("Successfully fetched winprobabilitypbp with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify GameInfo dataset structure
	if dataset, err := response.GetDataSet("GameInfo"); err == nil {
		assert.NotNil(t, dataset, "Should have GameInfo dataset")
		t.Logf("GameInfo: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset GameInfo not found (may be expected): %v", err)
	}

	// Verify WinProbPBP dataset structure
	if dataset, err := response.GetDataSet("WinProbPBP"); err == nil {
		assert.NotNil(t, dataset, "Should have WinProbPBP dataset")
		t.Logf("WinProbPBP: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset WinProbPBP not found (may be expected): %v", err)
	}
}
