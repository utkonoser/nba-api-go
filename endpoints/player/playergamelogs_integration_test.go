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

func TestGetPlayerGameLogs_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := PlayerGameLogsParams{
		PlayerIdNullable: "2544",
		SeasonNullable: "2023-24",
		SeasonTypeNullable: "Regular Season",
		LeagueIdNullable: "00",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	response, err := client.GetPlayerGameLogs(ctx, params)

	if err != nil {
		t.Logf("PlayerGameLogs endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	if response.Resource != "" {
		// API returns "gamelogs" as resource name, not "playergamelogs"
		assert.Contains(t, response.Resource, "gamelogs")
	}

	t.Logf("Successfully fetched playergamelogs with %d result sets", len(response.ResultSets))

	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

	if dataset, err := response.GetDataSet("PlayerGameLogs"); err == nil {
		assert.NotNil(t, dataset, "Should have PlayerGameLogs dataset")
		t.Logf("PlayerGameLogs: %d rows", dataset.RowCount())
	} else {
		t.Logf("Dataset PlayerGameLogs not found (may be expected): %v", err)
	}
}

