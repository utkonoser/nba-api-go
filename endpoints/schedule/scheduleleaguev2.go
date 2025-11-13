package schedule

import (
	"context"
	"fmt"
	"log/slog"
)

// ScheduleLeagueV2Params holds parameters for the ScheduleLeagueV2 endpoint.
type ScheduleLeagueV2Params struct {
	LeagueId string
	Season string
}

// GetScheduleLeagueV2 fetches data from the scheduleleaguev2 endpoint.
func (c *Client) GetScheduleLeagueV2(ctx context.Context, params ScheduleLeagueV2Params) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching scheduleleaguev2")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
		"Season": params.Season,
	}

	resp, err := c.httpClient.SendRequest(ctx, "scheduleleaguev2", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch scheduleleaguev2",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch scheduleleaguev2: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from scheduleleaguev2 endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal scheduleleaguev2 response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched scheduleleaguev2",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
