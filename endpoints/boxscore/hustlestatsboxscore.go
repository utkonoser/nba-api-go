package boxscore

import (
	"context"
	"fmt"
	"log/slog"
)

// HustleStatsBoxScoreParams holds parameters for the HustleStatsBoxScore endpoint.
type HustleStatsBoxScoreParams struct {
	GameId string
}

// GetHustleStatsBoxScore fetches data from the hustlestatsboxscore endpoint.
func (c *Client) GetHustleStatsBoxScore(ctx context.Context, params HustleStatsBoxScoreParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching hustlestatsboxscore")

	reqParams := map[string]string{
		"GameID": params.GameId,
	}

	resp, err := c.httpClient.SendRequest(ctx, "hustlestatsboxscore", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch hustlestatsboxscore",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch hustlestatsboxscore: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from hustlestatsboxscore endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal hustlestatsboxscore response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched hustlestatsboxscore",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
