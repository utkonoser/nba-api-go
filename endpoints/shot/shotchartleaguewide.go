package shot

import (
	"context"
	"fmt"
	"log/slog"
)

// ShotChartLeagueWideParams holds parameters for the ShotChartLeagueWide endpoint.
type ShotChartLeagueWideParams struct {
	LeagueId string
	Season string
}

// GetShotChartLeagueWide fetches data from the shotchartleaguewide endpoint.
func (c *Client) GetShotChartLeagueWide(ctx context.Context, params ShotChartLeagueWideParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching shotchartleaguewide")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
		"Season": params.Season,
	}

	resp, err := c.httpClient.SendRequest(ctx, "shotchartleaguewide", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch shotchartleaguewide",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch shotchartleaguewide: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from shotchartleaguewide endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal shotchartleaguewide response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched shotchartleaguewide",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
