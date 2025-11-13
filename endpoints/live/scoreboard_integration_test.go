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

func TestGetScoreboard_Integration(t *testing.T) {
	client := NewClient(nil)
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	scoreboard, err := client.GetScoreboard(ctx)

	require.NoError(t, err, "Should successfully fetch real scoreboard data")
	require.NotNil(t, scoreboard, "Scoreboard should not be nil")
	
	// Verify response structure
	assert.NotZero(t, scoreboard.Meta.Code, "Meta code should be set")
	assert.NotEmpty(t, scoreboard.Scoreboard.GameDate, "Game date should be set")
	assert.NotEmpty(t, scoreboard.Scoreboard.LeagueID, "League ID should be set")
	assert.NotEmpty(t, scoreboard.Scoreboard.LeagueName, "League name should be set")
	
	// Games can be empty (no games today), but structure should be valid
	t.Logf("Found %d games on %s", len(scoreboard.Scoreboard.Games), scoreboard.Scoreboard.GameDate)
	
	// If there are games, verify their structure
	for i, game := range scoreboard.Scoreboard.Games {
		assert.NotEmpty(t, game.GameID, "Game %d should have GameID", i)
		assert.NotEmpty(t, game.GameCode, "Game %d should have GameCode", i)
		assert.NotEmpty(t, game.GameStatusText, "Game %d should have GameStatusText", i)
		
		// Verify team structures
		assert.NotZero(t, game.HomeTeam.TeamID, "Game %d home team should have ID", i)
		assert.NotEmpty(t, game.HomeTeam.TeamName, "Game %d home team should have name", i)
		assert.NotEmpty(t, game.HomeTeam.TeamTricode, "Game %d home team should have tricode", i)
		
		assert.NotZero(t, game.AwayTeam.TeamID, "Game %d away team should have ID", i)
		assert.NotEmpty(t, game.AwayTeam.TeamName, "Game %d away team should have name", i)
		assert.NotEmpty(t, game.AwayTeam.TeamTricode, "Game %d away team should have tricode", i)
		
		t.Logf("Game %d: %s @ %s (%s)", i+1, game.AwayTeam.TeamTricode, game.HomeTeam.TeamTricode, game.GameStatusText)
	}
}

