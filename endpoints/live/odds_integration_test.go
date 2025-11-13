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

func TestGetOdds_Integration(t *testing.T) {
	// First get a game ID from today's scoreboard
	client := NewClient(nil)
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	scoreboard, err := client.GetScoreboard(ctx)
	require.NoError(t, err, "Should fetch scoreboard")
	
	if len(scoreboard.Scoreboard.Games) == 0 {
		t.Skip("No games today, skipping odds integration test")
		return
	}

	gameID := scoreboard.Scoreboard.Games[0].GameID
	t.Logf("Testing odds for game: %s", gameID)

	// Now get the odds
	ctx2, cancel2 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel2()

	odds, err := client.GetOdds(ctx2, gameID)

	// Note: Odds endpoint might not always be available for all games
	if err != nil {
		t.Logf("Odds not available for game %s: %v", gameID, err)
		t.Skip("Odds endpoint not available, skipping")
		return
	}

	require.NotNil(t, odds, "Odds should not be nil")
	
	// Verify response structure
	assert.NotZero(t, odds.Meta.Code, "Meta code should be set")
	assert.Equal(t, gameID, odds.Game.GameID, "Game ID should match")
	
	// Verify team structures
	assert.NotZero(t, odds.Game.HomeTeam.TeamID, "Home team ID should be set")
	assert.NotEmpty(t, odds.Game.HomeTeam.TeamName, "Home team name should be set")
	
	assert.NotZero(t, odds.Game.AwayTeam.TeamID, "Away team ID should be set")
	assert.NotEmpty(t, odds.Game.AwayTeam.TeamName, "Away team name should be set")
	
	t.Logf("Found %d odds providers", len(odds.Game.GameOdds))
	
	// If there are odds, verify their structure
	for i, odd := range odds.Game.GameOdds {
		assert.NotEmpty(t, odd.Provider, "Odd %d should have provider", i)
		t.Logf("Provider %s: Home ML %.1f, Away ML %.1f, O/U %.1f",
			odd.Provider,
			odd.HomeTeamOdds.Moneyline,
			odd.AwayTeamOdds.Moneyline,
			odd.OverUnder.Total)
	}
}

