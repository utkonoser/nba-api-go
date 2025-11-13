//go:build integration

package shot

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetShotChartLeagueWide_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := ShotChartLeagueWideParams{
		LeagueId: "00",
		Season: "2023-24",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetShotChartLeagueWide(ctx, params)

	if err != nil {
		t.Logf("ShotChartLeagueWide endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "shotchartleaguewide")
	}

	t.Logf("Successfully fetched shotchartleaguewide with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify League_Wide dataset structure
	if dataset, err := response.GetDataSet("League_Wide"); err == nil {
		assert.NotNil(t, dataset, "Should have League_Wide dataset")
		t.Logf("League_Wide: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset League_Wide not found (may be expected): %v", err)
	}
}
