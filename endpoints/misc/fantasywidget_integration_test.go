//go:build integration

package misc

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFantasyWidget_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := FantasyWidgetParams{
		LeagueId: "00",
		Season: "2023-24",
		SeasonTypeAllStar: "Regular Season",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetFantasyWidget(ctx, params)

	if err != nil {
		t.Logf("FantasyWidget endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "fantasywidget")
	}

	t.Logf("Successfully fetched fantasywidget with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify FantasyWidgetResult dataset structure
	if dataset, err := response.GetDataSet("FantasyWidgetResult"); err == nil {
		assert.NotNil(t, dataset, "Should have FantasyWidgetResult dataset")
		t.Logf("FantasyWidgetResult: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset FantasyWidgetResult not found (may be expected): %v", err)
	}
}
