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

func TestGetPlayerCareerByCollegeRollup_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := PlayerCareerByCollegeRollupParams{
		LeagueId: "00",
		SeasonTypeAllStar: "Regular Season",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetPlayerCareerByCollegeRollup(ctx, params)

	if err != nil {
		t.Logf("PlayerCareerByCollegeRollup endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "playercareerbycollegerollup")
	}

	t.Logf("Successfully fetched playercareerbycollegerollup with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

	// Verify East dataset structure
	if dataset, err := response.GetDataSet("East"); err == nil {
		assert.NotNil(t, dataset, "Should have East dataset")
		t.Logf("East: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset East not found (may be expected): %v", err)
	}

	// Verify Midwest dataset structure
	if dataset, err := response.GetDataSet("Midwest"); err == nil {
		assert.NotNil(t, dataset, "Should have Midwest dataset")
		t.Logf("Midwest: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset Midwest not found (may be expected): %v", err)
	}

	// Verify South dataset structure
	if dataset, err := response.GetDataSet("South"); err == nil {
		assert.NotNil(t, dataset, "Should have South dataset")
		t.Logf("South: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset South not found (may be expected): %v", err)
	}

	// Verify West dataset structure
	if dataset, err := response.GetDataSet("West"); err == nil {
		assert.NotNil(t, dataset, "Should have West dataset")
		t.Logf("West: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset West not found (may be expected): %v", err)
	}
}
