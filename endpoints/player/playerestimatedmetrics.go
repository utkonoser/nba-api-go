package player

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayerEstimatedMetricsParams holds parameters for the PlayerEstimatedMetrics endpoint.
type PlayerEstimatedMetricsParams struct {
	LeagueId string
	Season string
	SeasonType string
}

// GetPlayerEstimatedMetrics fetches data from the playerestimatedmetrics endpoint.
func (c *Client) GetPlayerEstimatedMetrics(ctx context.Context, params PlayerEstimatedMetricsParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playerestimatedmetrics")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
		"Season": params.Season,
		"SeasonType": params.SeasonType,
	}

	resp, err := c.httpClient.SendRequest(ctx, "playerestimatedmetrics", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch playerestimatedmetrics",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch playerestimatedmetrics: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from playerestimatedmetrics endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal playerestimatedmetrics response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched playerestimatedmetrics",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
