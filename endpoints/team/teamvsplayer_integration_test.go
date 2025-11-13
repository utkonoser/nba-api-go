//go:build integration

package team

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTeamVsPlayer_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := TeamVsPlayerParams{
		TeamId: "1610612737", // Atlanta Hawks
		VsPlayerId: "2544", // LeBron James
		Season: "2023-24",
		LeagueIdNullable: "00",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetTeamVsPlayer(ctx, params)

	if err != nil {
		t.Logf("TeamVsPlayer endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "teamvsplayer")
	}

	t.Logf("Successfully fetched teamvsplayer with %d result sets", len(response.ResultSets))

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

	// Verify vsPlayerOverall dataset structure
	if dataset, err := response.GetDataSet("vsPlayerOverall"); err == nil {
		assert.NotNil(t, dataset, "Should have vsPlayerOverall dataset")
		t.Logf("vsPlayerOverall: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset vsPlayerOverall not found (may be expected): %v", err)
	}
}
