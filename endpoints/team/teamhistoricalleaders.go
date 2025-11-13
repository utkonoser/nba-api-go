package team

import (
	"context"
	"fmt"
	"log/slog"
)

// TeamHistoricalLeadersParams holds parameters for the TeamHistoricalLeaders endpoint.
type TeamHistoricalLeadersParams struct {
	TeamId string
	LeagueId string
	SeasonId string
}

// GetTeamHistoricalLeaders fetches data from the teamhistoricalleaders endpoint.
func (c *Client) GetTeamHistoricalLeaders(ctx context.Context, params TeamHistoricalLeadersParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching teamhistoricalleaders")

	reqParams := map[string]string{
		"TeamID": params.TeamId,
		"LeagueID": params.LeagueId,
		"SeasonID": params.SeasonId,
	}

	resp, err := c.httpClient.SendRequest(ctx, "teamhistoricalleaders", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch teamhistoricalleaders",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch teamhistoricalleaders: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from teamhistoricalleaders endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal teamhistoricalleaders response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched teamhistoricalleaders",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
