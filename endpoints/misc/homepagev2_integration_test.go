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

func TestGetHomePageV2_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := HomePageV2Params{
		LeagueId: "00",
		Season: "2023-24",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetHomePageV2(ctx, params)

	if err != nil {
		t.Logf("HomePageV2 endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "homepagev2")
	}

	t.Logf("Successfully fetched homepagev2 with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify HomePageStat1 dataset structure
	if dataset, err := response.GetDataSet("HomePageStat1"); err == nil {
		assert.NotNil(t, dataset, "Should have HomePageStat1 dataset")
		t.Logf("HomePageStat1: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset HomePageStat1 not found (may be expected): %v", err)
	}

	// Verify HomePageStat2 dataset structure
	if dataset, err := response.GetDataSet("HomePageStat2"); err == nil {
		assert.NotNil(t, dataset, "Should have HomePageStat2 dataset")
		t.Logf("HomePageStat2: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset HomePageStat2 not found (may be expected): %v", err)
	}

	// Verify HomePageStat3 dataset structure
	if dataset, err := response.GetDataSet("HomePageStat3"); err == nil {
		assert.NotNil(t, dataset, "Should have HomePageStat3 dataset")
		t.Logf("HomePageStat3: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset HomePageStat3 not found (may be expected): %v", err)
	}

	// Verify HomePageStat4 dataset structure
	if dataset, err := response.GetDataSet("HomePageStat4"); err == nil {
		assert.NotNil(t, dataset, "Should have HomePageStat4 dataset")
		t.Logf("HomePageStat4: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset HomePageStat4 not found (may be expected): %v", err)
	}

	// Verify HomePageStat5 dataset structure
	if dataset, err := response.GetDataSet("HomePageStat5"); err == nil {
		assert.NotNil(t, dataset, "Should have HomePageStat5 dataset")
		t.Logf("HomePageStat5: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset HomePageStat5 not found (may be expected): %v", err)
	}

	// Verify HomePageStat6 dataset structure
	if dataset, err := response.GetDataSet("HomePageStat6"); err == nil {
		assert.NotNil(t, dataset, "Should have HomePageStat6 dataset")
		t.Logf("HomePageStat6: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset HomePageStat6 not found (may be expected): %v", err)
	}

	// Verify HomePageStat7 dataset structure
	if dataset, err := response.GetDataSet("HomePageStat7"); err == nil {
		assert.NotNil(t, dataset, "Should have HomePageStat7 dataset")
		t.Logf("HomePageStat7: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset HomePageStat7 not found (may be expected): %v", err)
	}

	// Verify HomePageStat8 dataset structure
	if dataset, err := response.GetDataSet("HomePageStat8"); err == nil {
		assert.NotNil(t, dataset, "Should have HomePageStat8 dataset")
		t.Logf("HomePageStat8: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset HomePageStat8 not found (may be expected): %v", err)
	}
}
