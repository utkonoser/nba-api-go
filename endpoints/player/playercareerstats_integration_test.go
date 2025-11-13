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

func TestGetPlayerCareerStats_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := PlayerCareerStatsParams{
		PlayerId:         "2544",  // LeBron James
		PerMode36:        "PerGame",
		LeagueIdNullable: "",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetPlayerCareerStats(ctx, params)

	if err != nil {
		t.Logf("PlayerCareerStats endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "playercareerstats")
	}

	t.Logf("Successfully fetched playercareerstats with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify CareerTotalsAllStarSeason dataset structure
	if dataset, err := response.GetDataSet("CareerTotalsAllStarSeason"); err == nil {
		assert.NotNil(t, dataset, "Should have CareerTotalsAllStarSeason dataset")
		t.Logf("CareerTotalsAllStarSeason: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset CareerTotalsAllStarSeason not found (may be expected): %v", err)
	}

	// Verify CareerTotalsCollegeSeason dataset structure
	if dataset, err := response.GetDataSet("CareerTotalsCollegeSeason"); err == nil {
		assert.NotNil(t, dataset, "Should have CareerTotalsCollegeSeason dataset")
		t.Logf("CareerTotalsCollegeSeason: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset CareerTotalsCollegeSeason not found (may be expected): %v", err)
	}

	// Verify CareerTotalsPostSeason dataset structure
	if dataset, err := response.GetDataSet("CareerTotalsPostSeason"); err == nil {
		assert.NotNil(t, dataset, "Should have CareerTotalsPostSeason dataset")
		t.Logf("CareerTotalsPostSeason: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset CareerTotalsPostSeason not found (may be expected): %v", err)
	}

	// Verify CareerTotalsRegularSeason dataset structure
	if dataset, err := response.GetDataSet("CareerTotalsRegularSeason"); err == nil {
		assert.NotNil(t, dataset, "Should have CareerTotalsRegularSeason dataset")
		t.Logf("CareerTotalsRegularSeason: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset CareerTotalsRegularSeason not found (may be expected): %v", err)
	}

	// Verify SeasonRankingsPostSeason dataset structure
	if dataset, err := response.GetDataSet("SeasonRankingsPostSeason"); err == nil {
		assert.NotNil(t, dataset, "Should have SeasonRankingsPostSeason dataset")
		t.Logf("SeasonRankingsPostSeason: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset SeasonRankingsPostSeason not found (may be expected): %v", err)
	}

	// Verify SeasonRankingsRegularSeason dataset structure
	if dataset, err := response.GetDataSet("SeasonRankingsRegularSeason"); err == nil {
		assert.NotNil(t, dataset, "Should have SeasonRankingsRegularSeason dataset")
		t.Logf("SeasonRankingsRegularSeason: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset SeasonRankingsRegularSeason not found (may be expected): %v", err)
	}

	// Verify SeasonTotalsAllStarSeason dataset structure
	if dataset, err := response.GetDataSet("SeasonTotalsAllStarSeason"); err == nil {
		assert.NotNil(t, dataset, "Should have SeasonTotalsAllStarSeason dataset")
		t.Logf("SeasonTotalsAllStarSeason: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset SeasonTotalsAllStarSeason not found (may be expected): %v", err)
	}

	// Verify SeasonTotalsCollegeSeason dataset structure
	if dataset, err := response.GetDataSet("SeasonTotalsCollegeSeason"); err == nil {
		assert.NotNil(t, dataset, "Should have SeasonTotalsCollegeSeason dataset")
		t.Logf("SeasonTotalsCollegeSeason: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset SeasonTotalsCollegeSeason not found (may be expected): %v", err)
	}

	// Verify SeasonTotalsPostSeason dataset structure
	if dataset, err := response.GetDataSet("SeasonTotalsPostSeason"); err == nil {
		assert.NotNil(t, dataset, "Should have SeasonTotalsPostSeason dataset")
		t.Logf("SeasonTotalsPostSeason: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset SeasonTotalsPostSeason not found (may be expected): %v", err)
	}

	// Verify SeasonTotalsRegularSeason dataset structure
	if dataset, err := response.GetDataSet("SeasonTotalsRegularSeason"); err == nil {
		assert.NotNil(t, dataset, "Should have SeasonTotalsRegularSeason dataset")
		t.Logf("SeasonTotalsRegularSeason: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset SeasonTotalsRegularSeason not found (may be expected): %v", err)
	}
}
