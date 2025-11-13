package boxscore

import (
	"context"
	"fmt"
	"log/slog"
)

// BoxScoreAdvancedV3Params holds parameters for the BoxScoreAdvancedV3 endpoint.
type BoxScoreAdvancedV3Params struct {
	GameId string
	EndPeriod string
	EndRange string
	RangeType string
	StartPeriod string
	StartRange string
}

// GetBoxScoreAdvancedV3 fetches data from the boxscoreadvancedv3 endpoint.
func (c *Client) GetBoxScoreAdvancedV3(ctx context.Context, params BoxScoreAdvancedV3Params) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching boxscoreadvancedv3")

	reqParams := map[string]string{
		"GameID": params.GameId,
		"EndPeriod": params.EndPeriod,
		"EndRange": params.EndRange,
		"RangeType": params.RangeType,
		"StartPeriod": params.StartPeriod,
		"StartRange": params.StartRange,
	}

	resp, err := c.httpClient.SendRequest(ctx, "boxscoreadvancedv3", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch boxscoreadvancedv3",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch boxscoreadvancedv3: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from boxscoreadvancedv3 endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal boxscoreadvancedv3 response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched boxscoreadvancedv3",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
