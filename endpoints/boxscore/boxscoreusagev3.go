package boxscore

import (
	"context"
	"fmt"
	"log/slog"
)

// BoxScoreUsageV3Params holds parameters for the BoxScoreUsageV3 endpoint.
type BoxScoreUsageV3Params struct {
	GameId string
	EndPeriod string
	EndRange string
	RangeType string
	StartPeriod string
	StartRange string
}

// GetBoxScoreUsageV3 fetches data from the boxscoreusagev3 endpoint.
func (c *Client) GetBoxScoreUsageV3(ctx context.Context, params BoxScoreUsageV3Params) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching boxscoreusagev3")

	reqParams := map[string]string{
		"GameID": params.GameId,
		"EndPeriod": params.EndPeriod,
		"EndRange": params.EndRange,
		"RangeType": params.RangeType,
		"StartPeriod": params.StartPeriod,
		"StartRange": params.StartRange,
	}

	resp, err := c.httpClient.SendRequest(ctx, "boxscoreusagev3", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch boxscoreusagev3",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch boxscoreusagev3: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from boxscoreusagev3 endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal boxscoreusagev3 response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched boxscoreusagev3",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
