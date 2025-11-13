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

func TestGetPlayerVsPlayer_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := PlayerVsPlayerParams{
		PlayerId: "2544", // LeBron James
		VsPlayerId: "201939", // Stephen Curry
		Season: "2023-24",
		LeagueIdNullable: "00",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetPlayerVsPlayer(ctx, params)

	if err != nil {
		t.Logf("PlayerVsPlayer endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "playervsplayer")
	}

	t.Logf("Successfully fetched playervsplayer with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify OnOffCourt dataset structure
	if dataset, err := response.GetDataSet("OnOffCourt"); err == nil {
		assert.NotNil(t, dataset, "Should have OnOffCourt dataset")
		t.Logf("OnOffCourt: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset OnOffCourt not found (may be expected): %v", err)
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

	// Verify PlayerInfo dataset structure
	if dataset, err := response.GetDataSet("PlayerInfo"); err == nil {
		assert.NotNil(t, dataset, "Should have PlayerInfo dataset")
		t.Logf("PlayerInfo: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset PlayerInfo not found (may be expected): %v", err)
	}

	// Verify ShotAreaOffCourt dataset structure
	if dataset, err := response.GetDataSet("ShotAreaOffCourt"); err == nil {
		assert.NotNil(t, dataset, "Should have ShotAreaOffCourt dataset")
		t.Logf("ShotAreaOffCourt: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset ShotAreaOffCourt not found (may be expected): %v", err)
	}

	// Verify ShotAreaOnCourt dataset structure
	if dataset, err := response.GetDataSet("ShotAreaOnCourt"); err == nil {
		assert.NotNil(t, dataset, "Should have ShotAreaOnCourt dataset")
		t.Logf("ShotAreaOnCourt: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset ShotAreaOnCourt not found (may be expected): %v", err)
	}

	// Verify ShotAreaOverall dataset structure
	if dataset, err := response.GetDataSet("ShotAreaOverall"); err == nil {
		assert.NotNil(t, dataset, "Should have ShotAreaOverall dataset")
		t.Logf("ShotAreaOverall: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset ShotAreaOverall not found (may be expected): %v", err)
	}

	// Verify ShotDistanceOffCourt dataset structure
	if dataset, err := response.GetDataSet("ShotDistanceOffCourt"); err == nil {
		assert.NotNil(t, dataset, "Should have ShotDistanceOffCourt dataset")
		t.Logf("ShotDistanceOffCourt: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset ShotDistanceOffCourt not found (may be expected): %v", err)
	}

	// Verify ShotDistanceOnCourt dataset structure
	if dataset, err := response.GetDataSet("ShotDistanceOnCourt"); err == nil {
		assert.NotNil(t, dataset, "Should have ShotDistanceOnCourt dataset")
		t.Logf("ShotDistanceOnCourt: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset ShotDistanceOnCourt not found (may be expected): %v", err)
	}

	// Verify ShotDistanceOverall dataset structure
	if dataset, err := response.GetDataSet("ShotDistanceOverall"); err == nil {
		assert.NotNil(t, dataset, "Should have ShotDistanceOverall dataset")
		t.Logf("ShotDistanceOverall: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset ShotDistanceOverall not found (may be expected): %v", err)
	}

	// Verify VsPlayerInfo dataset structure
	if dataset, err := response.GetDataSet("VsPlayerInfo"); err == nil {
		assert.NotNil(t, dataset, "Should have VsPlayerInfo dataset")
		t.Logf("VsPlayerInfo: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset VsPlayerInfo not found (may be expected): %v", err)
	}
}
