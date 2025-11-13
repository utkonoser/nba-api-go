package game

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayByPlayV3Params holds parameters for the PlayByPlayV3 endpoint.
type PlayByPlayV3Params struct {
	GameId string
	EndPeriod string
	StartPeriod string
}

// GetPlayByPlayV3 fetches data from the playbyplayv3 endpoint.
func (c *Client) GetPlayByPlayV3(ctx context.Context, params PlayByPlayV3Params) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playbyplayv3")

	reqParams := map[string]string{
		"GameID": params.GameId,
		"EndPeriod": params.EndPeriod,
		"StartPeriod": params.StartPeriod,
	}

	resp, err := c.httpClient.SendRequest(ctx, "playbyplayv3", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch playbyplayv3",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch playbyplayv3: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from playbyplayv3 endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal playbyplayv3 response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched playbyplayv3",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
