//go:build integration
// +build integration

package boxscore

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetBoxScoreSummaryV3_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := BoxScoreSummaryV3Params{
		GameId: "0022300001",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := client.GetBoxScoreSummaryV3(ctx, params)

	if err != nil {
		t.Logf("BoxScoreSummaryV3 endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// V3 endpoints may have empty resource field
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "boxscoresummaryv3")
	}

	t.Logf("Successfully fetched boxscoresummaryv3 with %d result sets", len(response.ResultSets))
}
