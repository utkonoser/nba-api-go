package player

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayerAwardsParams holds parameters for the PlayerAwards endpoint.
type PlayerAwardsParams struct {
	PlayerId string
}

// GetPlayerAwards fetches data from the playerawards endpoint.
func (c *Client) GetPlayerAwards(ctx context.Context, params PlayerAwardsParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playerawards")

	reqParams := map[string]string{
		"PlayerID": params.PlayerId,
	}

	resp, err := c.httpClient.SendRequest(ctx, "playerawards", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch playerawards",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch playerawards: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from playerawards endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal playerawards response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched playerawards",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}

