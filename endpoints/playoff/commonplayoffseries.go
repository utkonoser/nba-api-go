package playoff

import (
	"context"
	"fmt"
	"log/slog"
)

// CommonPlayoffSeriesParams holds parameters for the CommonPlayoffSeries endpoint.
type CommonPlayoffSeriesParams struct {
	LeagueId string
	Season string
	SeriesIdNullable string
}

// GetCommonPlayoffSeries fetches data from the commonplayoffseries endpoint.
func (c *Client) GetCommonPlayoffSeries(ctx context.Context, params CommonPlayoffSeriesParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching commonplayoffseries")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
		"Season": params.Season,
		"SeriesID": params.SeriesIdNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "commonplayoffseries", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch commonplayoffseries",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch commonplayoffseries: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from commonplayoffseries endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal commonplayoffseries response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched commonplayoffseries",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
