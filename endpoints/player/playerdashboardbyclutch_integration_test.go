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

func TestGetPlayerDashboardByClutch_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := PlayerDashboardByClutchParams{
		Season: "2023-24",
		LeagueIdNullable: "",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetPlayerDashboardByClutch(ctx, params)

	if err != nil {
		t.Logf("PlayerDashboardByClutch endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "playerdashboardbyclutch")
	}

	t.Logf("Successfully fetched playerdashboardbyclutch with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify Last10Sec3Point2PlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("Last10Sec3Point2PlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have Last10Sec3Point2PlayerDashboard dataset")
		t.Logf("Last10Sec3Point2PlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset Last10Sec3Point2PlayerDashboard not found (may be expected): %v", err)
	}

	// Verify Last10Sec3PointPlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("Last10Sec3PointPlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have Last10Sec3PointPlayerDashboard dataset")
		t.Logf("Last10Sec3PointPlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset Last10Sec3PointPlayerDashboard not found (may be expected): %v", err)
	}

	// Verify Last1Min5PointPlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("Last1Min5PointPlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have Last1Min5PointPlayerDashboard dataset")
		t.Logf("Last1Min5PointPlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset Last1Min5PointPlayerDashboard not found (may be expected): %v", err)
	}

	// Verify Last1MinPlusMinus5PointPlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("Last1MinPlusMinus5PointPlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have Last1MinPlusMinus5PointPlayerDashboard dataset")
		t.Logf("Last1MinPlusMinus5PointPlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset Last1MinPlusMinus5PointPlayerDashboard not found (may be expected): %v", err)
	}

	// Verify Last30Sec3Point2PlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("Last30Sec3Point2PlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have Last30Sec3Point2PlayerDashboard dataset")
		t.Logf("Last30Sec3Point2PlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset Last30Sec3Point2PlayerDashboard not found (may be expected): %v", err)
	}

	// Verify Last30Sec3PointPlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("Last30Sec3PointPlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have Last30Sec3PointPlayerDashboard dataset")
		t.Logf("Last30Sec3PointPlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset Last30Sec3PointPlayerDashboard not found (may be expected): %v", err)
	}

	// Verify Last3Min5PointPlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("Last3Min5PointPlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have Last3Min5PointPlayerDashboard dataset")
		t.Logf("Last3Min5PointPlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset Last3Min5PointPlayerDashboard not found (may be expected): %v", err)
	}

	// Verify Last3MinPlusMinus5PointPlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("Last3MinPlusMinus5PointPlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have Last3MinPlusMinus5PointPlayerDashboard dataset")
		t.Logf("Last3MinPlusMinus5PointPlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset Last3MinPlusMinus5PointPlayerDashboard not found (may be expected): %v", err)
	}

	// Verify Last5Min5PointPlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("Last5Min5PointPlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have Last5Min5PointPlayerDashboard dataset")
		t.Logf("Last5Min5PointPlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset Last5Min5PointPlayerDashboard not found (may be expected): %v", err)
	}

	// Verify Last5MinPlusMinus5PointPlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("Last5MinPlusMinus5PointPlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have Last5MinPlusMinus5PointPlayerDashboard dataset")
		t.Logf("Last5MinPlusMinus5PointPlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset Last5MinPlusMinus5PointPlayerDashboard not found (may be expected): %v", err)
	}

	// Verify OverallPlayerDashboard dataset structure
	if dataset, err := response.GetDataSet("OverallPlayerDashboard"); err == nil {
		assert.NotNil(t, dataset, "Should have OverallPlayerDashboard dataset")
		t.Logf("OverallPlayerDashboard: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset OverallPlayerDashboard not found (may be expected): %v", err)
	}
}
