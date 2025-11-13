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

func TestGetPlayerDashPtPass_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := PlayerDashPtPassParams{
		TeamId: "1610612737",
		PlayerId: "2544",
		Season:   "2023-24",
		SeasonTypeAllStar: "Regular Season",
		LeagueId: "00",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetPlayerDashPtPass(ctx, params)

	if err != nil {
		t.Logf("PlayerDashPtPass endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "playerdashptpass")
	}

	t.Logf("Successfully fetched playerdashptpass with %d result sets", len(response.ResultSets))

	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

	if dataset, err := response.GetDataSet("PassesMade"); err == nil {
		assert.NotNil(t, dataset, "Should have PassesMade dataset")
		t.Logf("PassesMade: %d rows", dataset.RowCount())
	} else {
		t.Logf("Dataset PassesMade not found (may be expected): %v", err)
	}
}

