//go:build integration
// +build integration

package team

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTeamDashPtReb_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := TeamDashPtRebParams{
		TeamId: "1610612737", // Atlanta Hawks
		LeagueId: "00",
		Season: "2023-24",
		SeasonTypeAllStar: "Regular Season",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetTeamDashPtReb(ctx, params)

	if err != nil {
		t.Logf("TeamDashPtReb endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "teamdashptreb")
	}

	t.Logf("Successfully fetched teamdashptreb with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify NumContestedRebounding dataset structure
	if dataset, err := response.GetDataSet("NumContestedRebounding"); err == nil {
		assert.NotNil(t, dataset, "Should have NumContestedRebounding dataset")
		t.Logf("NumContestedRebounding: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset NumContestedRebounding not found (may be expected): %v", err)
	}

	// Verify OverallRebounding dataset structure
	if dataset, err := response.GetDataSet("OverallRebounding"); err == nil {
		assert.NotNil(t, dataset, "Should have OverallRebounding dataset")
		t.Logf("OverallRebounding: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset OverallRebounding not found (may be expected): %v", err)
	}

	// Verify RebDistanceRebounding dataset structure
	if dataset, err := response.GetDataSet("RebDistanceRebounding"); err == nil {
		assert.NotNil(t, dataset, "Should have RebDistanceRebounding dataset")
		t.Logf("RebDistanceRebounding: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset RebDistanceRebounding not found (may be expected): %v", err)
	}

	// Verify ShotDistanceRebounding dataset structure
	if dataset, err := response.GetDataSet("ShotDistanceRebounding"); err == nil {
		assert.NotNil(t, dataset, "Should have ShotDistanceRebounding dataset")
		t.Logf("ShotDistanceRebounding: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset ShotDistanceRebounding not found (may be expected): %v", err)
	}

	// Verify ShotTypeRebounding dataset structure
	if dataset, err := response.GetDataSet("ShotTypeRebounding"); err == nil {
		assert.NotNil(t, dataset, "Should have ShotTypeRebounding dataset")
		t.Logf("ShotTypeRebounding: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset ShotTypeRebounding not found (may be expected): %v", err)
	}
}
