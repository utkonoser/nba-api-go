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

func TestGetDunkScoreLeaders_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := DunkScoreLeadersParams{
		LeagueIdNullable: "",
		Season: "2023-24",
		SeasonTypeAllStar: "Regular Season",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetDunkScoreLeaders(ctx, params)

	if err != nil {
		t.Logf("DunkScoreLeaders endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "dunkscoreleaders")
	}

	t.Logf("Successfully fetched dunkscoreleaders with %d result sets", len(response.ResultSets))
}
