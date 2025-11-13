package league

import (
	"context"
	"fmt"
	"log/slog"
)

// LeagueLeadersParams holds parameters for the LeagueLeaders endpoint.
type LeagueLeadersParams struct {
	LeagueId string
	PerMode48 string
	Scope string
	Season string
	SeasonTypeAllStar string
	StatCategoryAbbreviation string
	ActiveFlagNullable string
}

// GetLeagueLeaders fetches data from the leagueleaders endpoint.
func (c *Client) GetLeagueLeaders(ctx context.Context, params LeagueLeadersParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching leagueleaders")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
		"PerMode": params.PerMode48,
		"Scope": params.Scope,
		"Season": params.Season,
		"SeasonType": params.SeasonTypeAllStar,
		"StatCategory": params.StatCategoryAbbreviation,
		"ActiveFlag": params.ActiveFlagNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "leagueleaders", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch leagueleaders",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch leagueleaders: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from leagueleaders endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal leagueleaders response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched leagueleaders",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
