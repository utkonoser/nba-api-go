package leaders

import (
	"context"
	"fmt"
	"log/slog"
)

// HomePageLeadersParams holds parameters for the HomePageLeaders endpoint.
type HomePageLeadersParams struct {
	GameScopeDetailed string
	LeagueId string
	PlayerOrTeam string
	PlayerScope string
	Season string
	SeasonTypePlayoffs string
	StatCategory string
}

// GetHomePageLeaders fetches data from the homepageleaders endpoint.
func (c *Client) GetHomePageLeaders(ctx context.Context, params HomePageLeadersParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching homepageleaders")

	reqParams := map[string]string{
		"GameScope": params.GameScopeDetailed,
		"LeagueID": params.LeagueId,
		"PlayerOrTeam": params.PlayerOrTeam,
		"PlayerScope": params.PlayerScope,
		"Season": params.Season,
		"SeasonType": params.SeasonTypePlayoffs,
		"StatCategory": params.StatCategory,
	}

	resp, err := c.httpClient.SendRequest(ctx, "homepageleaders", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch homepageleaders",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch homepageleaders: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from homepageleaders endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal homepageleaders response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched homepageleaders",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
