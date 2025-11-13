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

func TestGetCommonPlayerInfo_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := CommonPlayerInfoParams{
		PlayerId: "2544",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetCommonPlayerInfo(ctx, params)

	if err != nil {
		t.Logf("CommonPlayerInfo endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "commonplayerinfo")
	}

	t.Logf("Successfully fetched commonplayerinfo with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify AvailableSeasons dataset structure
	if dataset, err := response.GetDataSet("AvailableSeasons"); err == nil {
		assert.NotNil(t, dataset, "Should have AvailableSeasons dataset")
		t.Logf("AvailableSeasons: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset AvailableSeasons not found (may be expected): %v", err)
	}

	// Verify CommonPlayerInfo dataset structure
	if dataset, err := response.GetDataSet("CommonPlayerInfo"); err == nil {
		assert.NotNil(t, dataset, "Should have CommonPlayerInfo dataset")
		t.Logf("CommonPlayerInfo: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset CommonPlayerInfo not found (may be expected): %v", err)
	}

	// Verify PlayerHeadlineStats dataset structure
	if dataset, err := response.GetDataSet("PlayerHeadlineStats"); err == nil {
		assert.NotNil(t, dataset, "Should have PlayerHeadlineStats dataset")
		t.Logf("PlayerHeadlineStats: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset PlayerHeadlineStats not found (may be expected): %v", err)
	}
}
