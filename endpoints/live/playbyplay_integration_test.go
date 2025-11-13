//go:build integration
// +build integration

package live

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPlayByPlay_Integration(t *testing.T) {
	// First get a game ID from today's scoreboard
	client := NewClient(nil)
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	scoreboard, err := client.GetScoreboard(ctx)
	require.NoError(t, err, "Should fetch scoreboard")
	
	if len(scoreboard.Scoreboard.Games) == 0 {
		t.Skip("No games today, skipping play-by-play integration test")
		return
	}

	gameID := scoreboard.Scoreboard.Games[0].GameID
	t.Logf("Testing play-by-play for game: %s", gameID)

	// Now get the play-by-play
	ctx2, cancel2 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel2()

	pbp, err := client.GetPlayByPlay(ctx2, gameID)

	require.NoError(t, err, "Should successfully fetch real play-by-play data")
	require.NotNil(t, pbp, "Play-by-play should not be nil")
	
	// Verify response structure
	assert.NotZero(t, pbp.Meta.Code, "Meta code should be set")
	assert.Equal(t, gameID, pbp.Game.GameID, "Game ID should match")
	
	t.Logf("Found %d actions in play-by-play", len(pbp.Game.Actions))
	
	// If there are actions, verify their structure
	if len(pbp.Game.Actions) > 0 {
		action := pbp.Game.Actions[0]
		assert.NotZero(t, action.ActionNumber, "Action should have number")
		assert.NotEmpty(t, action.Clock, "Action should have clock")
		assert.NotZero(t, action.Period, "Action should have period")
		assert.NotEmpty(t, action.ActionType, "Action should have type")
		
		t.Logf("First action: %s - %s (Period %d, %s)", 
			action.ActionType, action.Description, action.Period, action.Clock)
	}
}

