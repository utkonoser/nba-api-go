package team

import (
	"context"
	"fmt"
	"log/slog"
)

// TeamEstimatedMetricsParams holds parameters for the TeamEstimatedMetrics endpoint.
type TeamEstimatedMetricsParams struct {
	LeagueId string
	Season string
	SeasonType string
}

// GetTeamEstimatedMetrics fetches data from the teamestimatedmetrics endpoint.
func (c *Client) GetTeamEstimatedMetrics(ctx context.Context, params TeamEstimatedMetricsParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching teamestimatedmetrics")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
		"Season": params.Season,
		"SeasonType": params.SeasonType,
	}

	resp, err := c.httpClient.SendRequest(ctx, "teamestimatedmetrics", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch teamestimatedmetrics",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch teamestimatedmetrics: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from teamestimatedmetrics endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal teamestimatedmetrics response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched teamestimatedmetrics",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
