package boxscore

import (
	"context"
	"fmt"
	"log/slog"
)

// BoxScoreDefensiveV2Params holds parameters for the BoxScoreDefensiveV2 endpoint.
type BoxScoreDefensiveV2Params struct {
	GameId string
}

// GetBoxScoreDefensiveV2 fetches data from the boxscoredefensivev2 endpoint.
func (c *Client) GetBoxScoreDefensiveV2(ctx context.Context, params BoxScoreDefensiveV2Params) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching boxscoredefensivev2")

	reqParams := map[string]string{
		"GameID": params.GameId,
	}

	resp, err := c.httpClient.SendRequest(ctx, "boxscoredefensivev2", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch boxscoredefensivev2",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch boxscoredefensivev2: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from boxscoredefensivev2 endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal boxscoredefensivev2 response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched boxscoredefensivev2",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
