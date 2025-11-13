package team

import (
	"context"
	"fmt"
	"log/slog"
)

// CommonTeamRosterParams holds parameters for the CommonTeamRoster endpoint.
type CommonTeamRosterParams struct {
	TeamId string
	Season string
	LeagueIdNullable string
}

// GetCommonTeamRoster fetches data from the commonteamroster endpoint.
func (c *Client) GetCommonTeamRoster(ctx context.Context, params CommonTeamRosterParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching commonteamroster")

	reqParams := map[string]string{
		"TeamID": params.TeamId,
		"Season": params.Season,
		"LeagueID": params.LeagueIdNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "commonteamroster", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch commonteamroster",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch commonteamroster: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from commonteamroster endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal commonteamroster response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched commonteamroster",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
