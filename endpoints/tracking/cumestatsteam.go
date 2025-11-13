package tracking

import (
	"context"
	"fmt"
	"log/slog"
)

// CumeStatsTeamParams holds parameters for the CumeStatsTeam endpoint.
type CumeStatsTeamParams struct {
	TeamId string
	GameIds string
	LeagueId string
	Season string
	SeasonTypeAllStar string
}

// GetCumeStatsTeam fetches data from the cumestatsteam endpoint.
func (c *Client) GetCumeStatsTeam(ctx context.Context, params CumeStatsTeamParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching cumestatsteam")

	reqParams := map[string]string{
		"TeamID": params.TeamId,
		"GameIDs": params.GameIds,
		"LeagueID": params.LeagueId,
		"Season": params.Season,
		"SeasonType": params.SeasonTypeAllStar,
	}

	resp, err := c.httpClient.SendRequest(ctx, "cumestatsteam", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch cumestatsteam",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch cumestatsteam: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from cumestatsteam endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal cumestatsteam response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched cumestatsteam",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
