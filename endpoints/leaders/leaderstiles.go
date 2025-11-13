package leaders

import (
	"context"
	"fmt"
	"log/slog"
)

// LeadersTilesParams holds parameters for the LeadersTiles endpoint.
type LeadersTilesParams struct {
	GameScopeDetailed string
	LeagueId string
	PlayerOrTeam string
	PlayerScope string
	Season string
	SeasonTypePlayoffs string
	Stat string
}

// GetLeadersTiles fetches data from the leaderstiles endpoint.
func (c *Client) GetLeadersTiles(ctx context.Context, params LeadersTilesParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching leaderstiles")

	reqParams := map[string]string{
		"GameScope": params.GameScopeDetailed,
		"LeagueID": params.LeagueId,
		"PlayerOrTeam": params.PlayerOrTeam,
		"PlayerScope": params.PlayerScope,
		"Season": params.Season,
		"SeasonType": params.SeasonTypePlayoffs,
		"Stat": params.Stat,
	}

	resp, err := c.httpClient.SendRequest(ctx, "leaderstiles", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch leaderstiles",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch leaderstiles: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from leaderstiles endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal leaderstiles response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched leaderstiles",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
