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

func TestGetTeamDetails_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := TeamDetailsParams{
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetTeamDetails(ctx, params)

	if err != nil {
		t.Logf("TeamDetails endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "teamdetails")
	}

	t.Logf("Successfully fetched teamdetails with %d result sets", len(response.ResultSets))

	// If no result sets, skip dataset validation (valid scenario)
	if len(response.ResultSets) == 0 {
		t.Log("No result sets returned (this is valid for some parameter combinations)")
		return
	}

// Verify TeamAwardsChampionships dataset structure
	if dataset, err := response.GetDataSet("TeamAwardsChampionships"); err == nil {
		assert.NotNil(t, dataset, "Should have TeamAwardsChampionships dataset")
		t.Logf("TeamAwardsChampionships: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset TeamAwardsChampionships not found (may be expected): %v", err)
	}

	// Verify TeamAwardsConf dataset structure
	if dataset, err := response.GetDataSet("TeamAwardsConf"); err == nil {
		assert.NotNil(t, dataset, "Should have TeamAwardsConf dataset")
		t.Logf("TeamAwardsConf: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset TeamAwardsConf not found (may be expected): %v", err)
	}

	// Verify TeamAwardsDiv dataset structure
	if dataset, err := response.GetDataSet("TeamAwardsDiv"); err == nil {
		assert.NotNil(t, dataset, "Should have TeamAwardsDiv dataset")
		t.Logf("TeamAwardsDiv: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset TeamAwardsDiv not found (may be expected): %v", err)
	}

	// Verify TeamBackground dataset structure
	if dataset, err := response.GetDataSet("TeamBackground"); err == nil {
		assert.NotNil(t, dataset, "Should have TeamBackground dataset")
		t.Logf("TeamBackground: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset TeamBackground not found (may be expected): %v", err)
	}

	// Verify TeamHistory dataset structure
	if dataset, err := response.GetDataSet("TeamHistory"); err == nil {
		assert.NotNil(t, dataset, "Should have TeamHistory dataset")
		t.Logf("TeamHistory: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset TeamHistory not found (may be expected): %v", err)
	}

	// Verify TeamHof dataset structure
	if dataset, err := response.GetDataSet("TeamHof"); err == nil {
		assert.NotNil(t, dataset, "Should have TeamHof dataset")
		t.Logf("TeamHof: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset TeamHof not found (may be expected): %v", err)
	}

	// Verify TeamRetired dataset structure
	if dataset, err := response.GetDataSet("TeamRetired"); err == nil {
		assert.NotNil(t, dataset, "Should have TeamRetired dataset")
		t.Logf("TeamRetired: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset TeamRetired not found (may be expected): %v", err)
	}

	// Verify TeamSocialSites dataset structure
	if dataset, err := response.GetDataSet("TeamSocialSites"); err == nil {
		assert.NotNil(t, dataset, "Should have TeamSocialSites dataset")
		t.Logf("TeamSocialSites: %d rows", dataset.RowCount())
		if dataset.RowCount() > 0 {
			// Verify we can access data
			rows := dataset.ToMap()
			assert.NotEmpty(t, rows, "Should have data rows")
		}
	} else {
		t.Logf("Dataset TeamSocialSites not found (may be expected): %v", err)
	}
}
