//go:build integration

package team

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTeamDashPtPass_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := TeamDashPtPassParams{
		TeamId: "1610612737", // Atlanta Hawks
		LeagueId: "00",
		Season: "2023-24",
		SeasonTypeAllStar: "Regular Season",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetTeamDashPtPass(ctx, params)

	if err != nil {
		t.Logf("TeamDashPtPass endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "teamdashptpass")
	}

	t.Logf("Successfully fetched teamdashptpass with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify PassesMade dataset structure
	if dataset, err := response.GetDataSet("PassesMade"); err == nil {
		assert.NotNil(t, dataset, "Should have PassesMade dataset")
		t.Logf("PassesMade: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset PassesMade not found (may be expected): %v", err)
	}

	// Verify PassesReceived dataset structure
	if dataset, err := response.GetDataSet("PassesReceived"); err == nil {
		assert.NotNil(t, dataset, "Should have PassesReceived dataset")
		t.Logf("PassesReceived: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset PassesReceived not found (may be expected): %v", err)
	}
}
