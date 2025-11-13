package boxscore

import (
	"context"
	"fmt"
	"log/slog"
)

// BoxScoreMiscV3Params holds parameters for the BoxScoreMiscV3 endpoint.
type BoxScoreMiscV3Params struct {
	GameId string
	EndPeriod string
	EndRange string
	RangeType string
	StartPeriod string
	StartRange string
}

// GetBoxScoreMiscV3 fetches data from the boxscoremiscv3 endpoint.
func (c *Client) GetBoxScoreMiscV3(ctx context.Context, params BoxScoreMiscV3Params) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching boxscoremiscv3")

	reqParams := map[string]string{
		"GameID": params.GameId,
		"EndPeriod": params.EndPeriod,
		"EndRange": params.EndRange,
		"RangeType": params.RangeType,
		"StartPeriod": params.StartPeriod,
		"StartRange": params.StartRange,
	}

	resp, err := c.httpClient.SendRequest(ctx, "boxscoremiscv3", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch boxscoremiscv3",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch boxscoremiscv3: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from boxscoremiscv3 endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal boxscoremiscv3 response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched boxscoremiscv3",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
