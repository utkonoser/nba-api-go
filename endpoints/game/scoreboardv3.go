package game

import (
	"context"
	"fmt"
	"log/slog"
)

// ScoreboardV3Params holds parameters for the ScoreboardV3 endpoint.
type ScoreboardV3Params struct {
	GameDate string
	LeagueId string
}

// GetScoreboardV3 fetches data from the scoreboardv3 endpoint.
func (c *Client) GetScoreboardV3(ctx context.Context, params ScoreboardV3Params) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching scoreboardv3")

	reqParams := map[string]string{
		"GameDate": params.GameDate,
		"LeagueID": params.LeagueId,
	}

	resp, err := c.httpClient.SendRequest(ctx, "scoreboardv3", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch scoreboardv3",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch scoreboardv3: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from scoreboardv3 endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal scoreboardv3 response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched scoreboardv3",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
