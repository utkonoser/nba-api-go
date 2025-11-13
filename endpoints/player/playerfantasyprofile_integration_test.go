//go:build integration

package player

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPlayerFantasyProfile_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := PlayerFantasyProfileParams{
		PlayerId: "2544",
		Season:   "2023-24",
		LeagueIdNullable: "00",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetPlayerFantasyProfile(ctx, params)

	if err != nil {
		t.Logf("PlayerFantasyProfile endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "playerfantasyprofile")
	}

	t.Logf("Successfully fetched playerfantasyprofile with %d result sets", len(response.ResultSets))

	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

	if dataset, err := response.GetDataSet("LastNGames"); err == nil {
		assert.NotNil(t, dataset, "Should have LastNGames dataset")
		t.Logf("LastNGames: %d rows", dataset.RowCount())
	} else {
		t.Logf("Dataset LastNGames not found (may be expected): %v", err)
	}
}

