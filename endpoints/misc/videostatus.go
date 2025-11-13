package misc

import (
	"context"
	"fmt"
	"log/slog"
)

// VideoStatusParams holds parameters for the VideoStatus endpoint.
type VideoStatusParams struct {
	GameDate string
	LeagueId string
}

// GetVideoStatus fetches data from the videostatus endpoint.
func (c *Client) GetVideoStatus(ctx context.Context, params VideoStatusParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching videostatus")

	reqParams := map[string]string{
		"GameDate": params.GameDate,
		"LeagueID": params.LeagueId,
	}

	resp, err := c.httpClient.SendRequest(ctx, "videostatus", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch videostatus",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch videostatus: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from videostatus endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal videostatus response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched videostatus",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
