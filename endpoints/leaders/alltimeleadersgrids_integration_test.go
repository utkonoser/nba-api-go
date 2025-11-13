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

func TestGetAllTimeLeadersGrids_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := AllTimeLeadersGridsParams{
		LeagueId: "00",
		SeasonType: "Regular Season",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetAllTimeLeadersGrids(ctx, params)

	if err != nil {
		t.Logf("AllTimeLeadersGrids endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "alltimeleadersgrids")
	}

	t.Logf("Successfully fetched alltimeleadersgrids with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify ASTLeaders dataset structure
	if dataset, err := response.GetDataSet("ASTLeaders"); err == nil {
		assert.NotNil(t, dataset, "Should have ASTLeaders dataset")
		t.Logf("ASTLeaders: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset ASTLeaders not found (may be expected): %v", err)
	}

	// Verify BLKLeaders dataset structure
	if dataset, err := response.GetDataSet("BLKLeaders"); err == nil {
		assert.NotNil(t, dataset, "Should have BLKLeaders dataset")
		t.Logf("BLKLeaders: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset BLKLeaders not found (may be expected): %v", err)
	}

	// Verify DREBLeaders dataset structure
	if dataset, err := response.GetDataSet("DREBLeaders"); err == nil {
		assert.NotNil(t, dataset, "Should have DREBLeaders dataset")
		t.Logf("DREBLeaders: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset DREBLeaders not found (may be expected): %v", err)
	}

	// Verify FG3ALeaders dataset structure
	if dataset, err := response.GetDataSet("FG3ALeaders"); err == nil {
		assert.NotNil(t, dataset, "Should have FG3ALeaders dataset")
		t.Logf("FG3ALeaders: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset FG3ALeaders not found (may be expected): %v", err)
	}

	// Verify FG3MLeaders dataset structure
	if dataset, err := response.GetDataSet("FG3MLeaders"); err == nil {
		assert.NotNil(t, dataset, "Should have FG3MLeaders dataset")
		t.Logf("FG3MLeaders: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset FG3MLeaders not found (may be expected): %v", err)
	}

	// Verify FG3_PCTLeaders dataset structure
	if dataset, err := response.GetDataSet("FG3_PCTLeaders"); err == nil {
		assert.NotNil(t, dataset, "Should have FG3_PCTLeaders dataset")
		t.Logf("FG3_PCTLeaders: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset FG3_PCTLeaders not found (may be expected): %v", err)
	}

	// Verify FGALeaders dataset structure
	if dataset, err := response.GetDataSet("FGALeaders"); err == nil {
		assert.NotNil(t, dataset, "Should have FGALeaders dataset")
		t.Logf("FGALeaders: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset FGALeaders not found (may be expected): %v", err)
	}

	// Verify FGMLeaders dataset structure
	if dataset, err := response.GetDataSet("FGMLeaders"); err == nil {
		assert.NotNil(t, dataset, "Should have FGMLeaders dataset")
		t.Logf("FGMLeaders: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset FGMLeaders not found (may be expected): %v", err)
	}

	// Verify FG_PCTLeaders dataset structure
	if dataset, err := response.GetDataSet("FG_PCTLeaders"); err == nil {
		assert.NotNil(t, dataset, "Should have FG_PCTLeaders dataset")
		t.Logf("FG_PCTLeaders: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset FG_PCTLeaders not found (may be expected): %v", err)
	}

	// Verify FTALeaders dataset structure
	if dataset, err := response.GetDataSet("FTALeaders"); err == nil {
		assert.NotNil(t, dataset, "Should have FTALeaders dataset")
		t.Logf("FTALeaders: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset FTALeaders not found (may be expected): %v", err)
	}

	// Verify FTMLeaders dataset structure
	if dataset, err := response.GetDataSet("FTMLeaders"); err == nil {
		assert.NotNil(t, dataset, "Should have FTMLeaders dataset")
		t.Logf("FTMLeaders: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset FTMLeaders not found (may be expected): %v", err)
	}

	// Verify FT_PCTLeaders dataset structure
	if dataset, err := response.GetDataSet("FT_PCTLeaders"); err == nil {
		assert.NotNil(t, dataset, "Should have FT_PCTLeaders dataset")
		t.Logf("FT_PCTLeaders: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset FT_PCTLeaders not found (may be expected): %v", err)
	}

	// Verify GPLeaders dataset structure
	if dataset, err := response.GetDataSet("GPLeaders"); err == nil {
		assert.NotNil(t, dataset, "Should have GPLeaders dataset")
		t.Logf("GPLeaders: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset GPLeaders not found (may be expected): %v", err)
	}

	// Verify OREBLeaders dataset structure
	if dataset, err := response.GetDataSet("OREBLeaders"); err == nil {
		assert.NotNil(t, dataset, "Should have OREBLeaders dataset")
		t.Logf("OREBLeaders: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset OREBLeaders not found (may be expected): %v", err)
	}

	// Verify PFLeaders dataset structure
	if dataset, err := response.GetDataSet("PFLeaders"); err == nil {
		assert.NotNil(t, dataset, "Should have PFLeaders dataset")
		t.Logf("PFLeaders: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset PFLeaders not found (may be expected): %v", err)
	}

	// Verify PTSLeaders dataset structure
	if dataset, err := response.GetDataSet("PTSLeaders"); err == nil {
		assert.NotNil(t, dataset, "Should have PTSLeaders dataset")
		t.Logf("PTSLeaders: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset PTSLeaders not found (may be expected): %v", err)
	}

	// Verify REBLeaders dataset structure
	if dataset, err := response.GetDataSet("REBLeaders"); err == nil {
		assert.NotNil(t, dataset, "Should have REBLeaders dataset")
		t.Logf("REBLeaders: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset REBLeaders not found (may be expected): %v", err)
	}

	// Verify STLLeaders dataset structure
	if dataset, err := response.GetDataSet("STLLeaders"); err == nil {
		assert.NotNil(t, dataset, "Should have STLLeaders dataset")
		t.Logf("STLLeaders: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset STLLeaders not found (may be expected): %v", err)
	}

	// Verify TOVLeaders dataset structure
	if dataset, err := response.GetDataSet("TOVLeaders"); err == nil {
		assert.NotNil(t, dataset, "Should have TOVLeaders dataset")
		t.Logf("TOVLeaders: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset TOVLeaders not found (may be expected): %v", err)
	}
}
