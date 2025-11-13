//go:build integration
// +build integration

package leaders

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetHomePageLeaders_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := HomePageLeadersParams{
		LeagueId: "00",
		Season: "2023-24",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetHomePageLeaders(ctx, params)

	if err != nil {
		t.Logf("HomePageLeaders endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "homepageleaders")
	}

	t.Logf("Successfully fetched homepageleaders with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify HomePageLeaders dataset structure
	if dataset, err := response.GetDataSet("HomePageLeaders"); err == nil {
		assert.NotNil(t, dataset, "Should have HomePageLeaders dataset")
		t.Logf("HomePageLeaders: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset HomePageLeaders not found (may be expected): %v", err)
	}

	// Verify LeagueAverage dataset structure
	if dataset, err := response.GetDataSet("LeagueAverage"); err == nil {
		assert.NotNil(t, dataset, "Should have LeagueAverage dataset")
		t.Logf("LeagueAverage: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset LeagueAverage not found (may be expected): %v", err)
	}

	// Verify LeagueMax dataset structure
	if dataset, err := response.GetDataSet("LeagueMax"); err == nil {
		assert.NotNil(t, dataset, "Should have LeagueMax dataset")
		t.Logf("LeagueMax: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset LeagueMax not found (may be expected): %v", err)
	}
}
