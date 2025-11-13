package live

import (
	"context"
	"fmt"
	"log/slog"
)

// OddsResponse represents the full odds response.
type OddsResponse struct {
	Meta Meta       `json:"meta"`
	Game OddsGame   `json:"game"`
}

// OddsGame represents the game odds data.
type OddsGame struct {
	GameID       string       `json:"gameId"`
	HomeTeam     OddsTeam     `json:"homeTeam"`
	AwayTeam     OddsTeam     `json:"awayTeam"`
	GameOdds     []GameOdd    `json:"gameOdds"`
}

// OddsTeam represents a team in odds data.
type OddsTeam struct {
	TeamID      int    `json:"teamId"`
	TeamName    string `json:"teamName"`
	TeamCity    string `json:"teamCity"`
	TeamTricode string `json:"teamTricode"`
}

// GameOdd represents odds from a provider.
type GameOdd struct {
	Provider        string  `json:"provider"`
	HomeTeamOdds    OddData `json:"homeTeamOdds"`
	AwayTeamOdds    OddData `json:"awayTeamOdds"`
	OverUnder       OverUnderOdd `json:"overUnder"`
	Suspended       int     `json:"suspended"`
}

// OddData represents odds data for a team.
type OddData struct {
	Moneyline   float64 `json:"moneyline"`
	Spread      float64 `json:"spread"`
	SpreadOdds  float64 `json:"spreadOdds"`
}

// OverUnderOdd represents over/under odds.
type OverUnderOdd struct {
	Total     float64 `json:"total"`
	OverOdds  float64 `json:"overOdds"`
	UnderOdds float64 `json:"underOdds"`
}

// GetOdds fetches the odds for a specific game.
func (c *Client) GetOdds(ctx context.Context, gameID string) (*OddsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching NBA odds",
		slog.String("game_id", gameID))

	endpoint := fmt.Sprintf("odds/odds_%s.json", gameID)
	resp, err := c.httpClient.SendRequest(ctx, endpoint, nil)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch odds",
			slog.String("game_id", gameID),
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch odds: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from odds endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var oddsResp OddsResponse
	if err := resp.GetJSON(&oddsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal odds response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched odds",
		slog.String("game_id", gameID),
		slog.Int("providers_count", len(oddsResp.Game.GameOdds)))

	return &oddsResp, nil
}

