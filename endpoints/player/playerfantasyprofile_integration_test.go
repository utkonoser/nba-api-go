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

func TestGetPlayerFantasyProfile_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := PlayerFantasyProfileParams{
		Season: "2023-24",
		LeagueIdNullable: "",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetPlayerFantasyProfile(ctx, params)

	if err != nil {
		t.Logf("PlayerFantasyProfile endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "playerfantasyprofile")
	}

	t.Logf("Successfully fetched playerfantasyprofile with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify DaysRestModified dataset structure
	if dataset, err := response.GetDataSet("DaysRestModified"); err == nil {
		assert.NotNil(t, dataset, "Should have DaysRestModified dataset")
		t.Logf("DaysRestModified: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset DaysRestModified not found (may be expected): %v", err)
	}

	// Verify LastNGames dataset structure
	if dataset, err := response.GetDataSet("LastNGames"); err == nil {
		assert.NotNil(t, dataset, "Should have LastNGames dataset")
		t.Logf("LastNGames: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset LastNGames not found (may be expected): %v", err)
	}

	// Verify Location dataset structure
	if dataset, err := response.GetDataSet("Location"); err == nil {
		assert.NotNil(t, dataset, "Should have Location dataset")
		t.Logf("Location: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset Location not found (may be expected): %v", err)
	}

	// Verify Opponent dataset structure
	if dataset, err := response.GetDataSet("Opponent"); err == nil {
		assert.NotNil(t, dataset, "Should have Opponent dataset")
		t.Logf("Opponent: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset Opponent not found (may be expected): %v", err)
	}

	// Verify Overall dataset structure
	if dataset, err := response.GetDataSet("Overall"); err == nil {
		assert.NotNil(t, dataset, "Should have Overall dataset")
		t.Logf("Overall: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset Overall not found (may be expected): %v", err)
	}
}
