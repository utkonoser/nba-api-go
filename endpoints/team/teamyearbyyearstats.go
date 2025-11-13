package team

import (
	"context"
	"fmt"
	"log/slog"
)

// TeamYearByYearStatsParams holds parameters for the TeamYearByYearStats endpoint.
type TeamYearByYearStatsParams struct {
	TeamId string
	LeagueId string
	PerModeSimple string
	SeasonTypeAllStar string
}

// GetTeamYearByYearStats fetches data from the teamyearbyyearstats endpoint.
func (c *Client) GetTeamYearByYearStats(ctx context.Context, params TeamYearByYearStatsParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching teamyearbyyearstats")

	reqParams := map[string]string{
		"TeamID": params.TeamId,
		"LeagueID": params.LeagueId,
		"PerMode": params.PerModeSimple,
		"SeasonType": params.SeasonTypeAllStar,
	}

	resp, err := c.httpClient.SendRequest(ctx, "teamyearbyyearstats", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch teamyearbyyearstats",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch teamyearbyyearstats: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from teamyearbyyearstats endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal teamyearbyyearstats response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched teamyearbyyearstats",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
