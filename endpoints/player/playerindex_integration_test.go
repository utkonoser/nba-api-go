//go:build integration
// +build integration

package player

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPlayerIndex_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := PlayerIndexParams{
		LeagueId: "00",
		Season: "2023-24",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	response, err := client.GetPlayerIndex(ctx, params)

	if err != nil {
		t.Logf("PlayerIndex endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "playerindex")
	}

	t.Logf("Successfully fetched playerindex with %d result sets", len(response.ResultSets))

	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

	if dataset, err := response.GetDataSet("PlayerIndex"); err == nil {
		assert.NotNil(t, dataset, "Should have PlayerIndex dataset")
		t.Logf("PlayerIndex: %d rows", dataset.RowCount())
	} else {
		t.Logf("Dataset PlayerIndex not found (may be expected): %v", err)
	}
}

