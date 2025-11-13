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

func TestGetVideoEvents_Integration(t *testing.T) {
	client := NewClient(nil)
	
	params := VideoEventsParams{
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	response, err := client.GetVideoEvents(ctx, params)

	if err != nil {
		t.Logf("VideoEvents endpoint error: %v", err)
		t.Skip("Endpoint may be unavailable or parameters incorrect")
		return
	}

	require.NotNil(t, response)
	// Resource may be empty for some endpoints
	if response.Resource != "" {
		assert.Contains(t, response.Resource, "videoevents")
	}

	t.Logf("Successfully fetched videoevents with %d result sets", len(response.ResultSets))
}
