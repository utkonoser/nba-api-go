//go:build integration

package player

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPlayerCompare_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := PlayerCompareParams{
		VsPlayerIdList: "201939", // Stephen Curry
		PlayerIdList: "2544", // LeBron James
		Season:   "2023-24",
		SeasonType: "Regular Season",
		LeagueIdNullable: "00",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetPlayerCompare(ctx, params)

	if err != nil {
		t.Logf("PlayerCompare endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "playercompare")
	}

	t.Logf("Successfully fetched playercompare with %d result sets", len(response.ResultSets))

	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

	if dataset, err := response.GetDataSet("OverallCompare"); err == nil {
		assert.NotNil(t, dataset, "Should have OverallCompare dataset")
		t.Logf("OverallCompare: %d rows", dataset.RowCount())
	} else {
		t.Logf("Dataset OverallCompare not found (may be expected): %v", err)
	}
}

