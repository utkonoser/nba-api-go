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

func TestGetBoxScore_Integration(t *testing.T) {
	// First get a game ID from today's scoreboard
	client := NewClient(nil)
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	scoreboard, err := client.GetScoreboard(ctx)
	require.NoError(t, err, "Should fetch scoreboard")
	
	if len(scoreboard.Scoreboard.Games) == 0 {
		t.Skip("No games today, skipping boxscore integration test")
		return
	}

	gameID := scoreboard.Scoreboard.Games[0].GameID
	t.Logf("Testing boxscore for game: %s", gameID)

	// Now get the boxscore
	ctx2, cancel2 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel2()

	boxscore, err := client.GetBoxScore(ctx2, gameID)

	require.NoError(t, err, "Should successfully fetch real boxscore data")
	require.NotNil(t, boxscore, "Boxscore should not be nil")
	
	// Verify response structure
	assert.NotZero(t, boxscore.Meta.Code, "Meta code should be set")
	assert.Equal(t, gameID, boxscore.Game.GameID, "Game ID should match")
	assert.NotEmpty(t, boxscore.Game.GameCode, "Game code should be set")
	assert.NotEmpty(t, boxscore.Game.GameStatusText, "Game status should be set")
	
	// Verify arena
	assert.NotZero(t, boxscore.Game.Arena.ArenaID, "Arena ID should be set")
	assert.NotEmpty(t, boxscore.Game.Arena.ArenaName, "Arena name should be set")
	
	// Verify home team
	assert.NotZero(t, boxscore.Game.HomeTeam.TeamID, "Home team ID should be set")
	assert.NotEmpty(t, boxscore.Game.HomeTeam.TeamName, "Home team name should be set")
	assert.NotEmpty(t, boxscore.Game.HomeTeam.TeamTricode, "Home team tricode should be set")
	
	// Verify away team
	assert.NotZero(t, boxscore.Game.AwayTeam.TeamID, "Away team ID should be set")
	assert.NotEmpty(t, boxscore.Game.AwayTeam.TeamName, "Away team name should be set")
	assert.NotEmpty(t, boxscore.Game.AwayTeam.TeamTricode, "Away team tricode should be set")
	
	t.Logf("Boxscore: %s %d - %s %d (%s)",
		boxscore.Game.AwayTeam.TeamTricode, boxscore.Game.AwayTeam.Score,
		boxscore.Game.HomeTeam.TeamTricode, boxscore.Game.HomeTeam.Score,
		boxscore.Game.GameStatusText)
	
	// If game has started, verify player data structure
	if len(boxscore.Game.HomeTeam.Players) > 0 {
		player := boxscore.Game.HomeTeam.Players[0]
		assert.NotZero(t, player.PersonID, "Player should have PersonID")
		assert.NotEmpty(t, player.Name, "Player should have Name")
		t.Logf("Sample home player: %s", player.Name)
	}
}

