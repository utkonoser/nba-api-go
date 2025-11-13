package player

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayerCareerStatsParams holds parameters for the PlayerCareerStats endpoint.
type PlayerCareerStatsParams struct {
	PlayerId string
	PerMode36 string
	LeagueIdNullable string
}

// GetPlayerCareerStats fetches data from the playercareerstats endpoint.
func (c *Client) GetPlayerCareerStats(ctx context.Context, params PlayerCareerStatsParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playercareerstats")

	reqParams := map[string]string{
		"PlayerID": params.PlayerId,
		"PerMode": params.PerMode36,
		"LeagueID": params.LeagueIdNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "playercareerstats", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch playercareerstats",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch playercareerstats: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from playercareerstats endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal playercareerstats response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched playercareerstats",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
