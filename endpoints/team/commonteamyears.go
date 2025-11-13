package team

import (
	"context"
	"fmt"
	"log/slog"
)

// CommonTeamYearsParams holds parameters for the CommonTeamYears endpoint.
type CommonTeamYearsParams struct {
	LeagueId string
}

// GetCommonTeamYears fetches data from the commonteamyears endpoint.
func (c *Client) GetCommonTeamYears(ctx context.Context, params CommonTeamYearsParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching commonteamyears")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
	}

	resp, err := c.httpClient.SendRequest(ctx, "commonteamyears", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch commonteamyears",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch commonteamyears: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from commonteamyears endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal commonteamyears response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched commonteamyears",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
