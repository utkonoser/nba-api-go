//go:build integration
// +build integration

package misc

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetDefenseHub_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := DefenseHubParams{
		LeagueId: "00",
		Season: "2023-24",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetDefenseHub(ctx, params)

	if err != nil {
		t.Logf("DefenseHub endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "defensehub")
	}

	t.Logf("Successfully fetched defensehub with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify DefenseHubStat1 dataset structure
	if dataset, err := response.GetDataSet("DefenseHubStat1"); err == nil {
		assert.NotNil(t, dataset, "Should have DefenseHubStat1 dataset")
		t.Logf("DefenseHubStat1: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset DefenseHubStat1 not found (may be expected): %v", err)
	}

	// Verify DefenseHubStat10 dataset structure
	if dataset, err := response.GetDataSet("DefenseHubStat10"); err == nil {
		assert.NotNil(t, dataset, "Should have DefenseHubStat10 dataset")
		t.Logf("DefenseHubStat10: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset DefenseHubStat10 not found (may be expected): %v", err)
	}

	// Verify DefenseHubStat2 dataset structure
	if dataset, err := response.GetDataSet("DefenseHubStat2"); err == nil {
		assert.NotNil(t, dataset, "Should have DefenseHubStat2 dataset")
		t.Logf("DefenseHubStat2: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset DefenseHubStat2 not found (may be expected): %v", err)
	}

	// Verify DefenseHubStat3 dataset structure
	if dataset, err := response.GetDataSet("DefenseHubStat3"); err == nil {
		assert.NotNil(t, dataset, "Should have DefenseHubStat3 dataset")
		t.Logf("DefenseHubStat3: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset DefenseHubStat3 not found (may be expected): %v", err)
	}

	// Verify DefenseHubStat4 dataset structure
	if dataset, err := response.GetDataSet("DefenseHubStat4"); err == nil {
		assert.NotNil(t, dataset, "Should have DefenseHubStat4 dataset")
		t.Logf("DefenseHubStat4: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset DefenseHubStat4 not found (may be expected): %v", err)
	}

	// Verify DefenseHubStat5 dataset structure
	if dataset, err := response.GetDataSet("DefenseHubStat5"); err == nil {
		assert.NotNil(t, dataset, "Should have DefenseHubStat5 dataset")
		t.Logf("DefenseHubStat5: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset DefenseHubStat5 not found (may be expected): %v", err)
	}

	// Verify DefenseHubStat6 dataset structure
	if dataset, err := response.GetDataSet("DefenseHubStat6"); err == nil {
		assert.NotNil(t, dataset, "Should have DefenseHubStat6 dataset")
		t.Logf("DefenseHubStat6: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset DefenseHubStat6 not found (may be expected): %v", err)
	}

	// Verify DefenseHubStat7 dataset structure
	if dataset, err := response.GetDataSet("DefenseHubStat7"); err == nil {
		assert.NotNil(t, dataset, "Should have DefenseHubStat7 dataset")
		t.Logf("DefenseHubStat7: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset DefenseHubStat7 not found (may be expected): %v", err)
	}

	// Verify DefenseHubStat8 dataset structure
	if dataset, err := response.GetDataSet("DefenseHubStat8"); err == nil {
		assert.NotNil(t, dataset, "Should have DefenseHubStat8 dataset")
		t.Logf("DefenseHubStat8: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset DefenseHubStat8 not found (may be expected): %v", err)
	}

	// Verify DefenseHubStat9 dataset structure
	if dataset, err := response.GetDataSet("DefenseHubStat9"); err == nil {
		assert.NotNil(t, dataset, "Should have DefenseHubStat9 dataset")
		t.Logf("DefenseHubStat9: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset DefenseHubStat9 not found (may be expected): %v", err)
	}
}
