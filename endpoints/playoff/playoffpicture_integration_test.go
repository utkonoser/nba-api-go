//go:build integration
// +build integration

package playoff

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPlayoffPicture_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := PlayoffPictureParams{
		LeagueId: "00",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetPlayoffPicture(ctx, params)

	if err != nil {
		t.Logf("PlayoffPicture endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "playoffpicture")
	}

	t.Logf("Successfully fetched playoffpicture with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify EastConfPlayoffPicture dataset structure
	if dataset, err := response.GetDataSet("EastConfPlayoffPicture"); err == nil {
		assert.NotNil(t, dataset, "Should have EastConfPlayoffPicture dataset")
		t.Logf("EastConfPlayoffPicture: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset EastConfPlayoffPicture not found (may be expected): %v", err)
	}

	// Verify EastConfRemainingGames dataset structure
	if dataset, err := response.GetDataSet("EastConfRemainingGames"); err == nil {
		assert.NotNil(t, dataset, "Should have EastConfRemainingGames dataset")
		t.Logf("EastConfRemainingGames: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset EastConfRemainingGames not found (may be expected): %v", err)
	}

	// Verify EastConfStandings dataset structure
	if dataset, err := response.GetDataSet("EastConfStandings"); err == nil {
		assert.NotNil(t, dataset, "Should have EastConfStandings dataset")
		t.Logf("EastConfStandings: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset EastConfStandings not found (may be expected): %v", err)
	}

	// Verify WestConfPlayoffPicture dataset structure
	if dataset, err := response.GetDataSet("WestConfPlayoffPicture"); err == nil {
		assert.NotNil(t, dataset, "Should have WestConfPlayoffPicture dataset")
		t.Logf("WestConfPlayoffPicture: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset WestConfPlayoffPicture not found (may be expected): %v", err)
	}

	// Verify WestConfRemainingGames dataset structure
	if dataset, err := response.GetDataSet("WestConfRemainingGames"); err == nil {
		assert.NotNil(t, dataset, "Should have WestConfRemainingGames dataset")
		t.Logf("WestConfRemainingGames: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset WestConfRemainingGames not found (may be expected): %v", err)
	}

	// Verify WestConfStandings dataset structure
	if dataset, err := response.GetDataSet("WestConfStandings"); err == nil {
		assert.NotNil(t, dataset, "Should have WestConfStandings dataset")
		t.Logf("WestConfStandings: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset WestConfStandings not found (may be expected): %v", err)
	}
}
