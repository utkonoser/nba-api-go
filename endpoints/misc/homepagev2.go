package misc

import (
	"context"
	"fmt"
	"log/slog"
)

// HomePageV2Params holds parameters for the HomePageV2 endpoint.
type HomePageV2Params struct {
	GameScopeDetailed string
	LeagueId string
	PlayerOrTeam string
	PlayerScope string
	Season string
	SeasonTypePlayoffs string
	StatType string
}

// GetHomePageV2 fetches data from the homepagev2 endpoint.
func (c *Client) GetHomePageV2(ctx context.Context, params HomePageV2Params) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching homepagev2")

	reqParams := map[string]string{
		"GameScope": params.GameScopeDetailed,
		"LeagueID": params.LeagueId,
		"PlayerOrTeam": params.PlayerOrTeam,
		"PlayerScope": params.PlayerScope,
		"Season": params.Season,
		"SeasonType": params.SeasonTypePlayoffs,
		"StatType": params.StatType,
	}

	resp, err := c.httpClient.SendRequest(ctx, "homepagev2", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch homepagev2",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch homepagev2: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from homepagev2 endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal homepagev2 response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched homepagev2",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
