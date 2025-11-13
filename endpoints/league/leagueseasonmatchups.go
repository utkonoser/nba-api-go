package league

import (
	"context"
	"fmt"
	"log/slog"
)

// LeagueSeasonMatchupsParams holds parameters for the LeagueSeasonMatchups endpoint.
type LeagueSeasonMatchupsParams struct {
	LeagueId string
	PerModeSimple string
	Season string
	SeasonTypePlayoffs string
	DefPlayerIdNullable string
	DefTeamIdNullable string
	OffPlayerIdNullable string
	OffTeamIdNullable string
}

// GetLeagueSeasonMatchups fetches data from the leagueseasonmatchups endpoint.
func (c *Client) GetLeagueSeasonMatchups(ctx context.Context, params LeagueSeasonMatchupsParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching leagueseasonmatchups")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
		"PerMode": params.PerModeSimple,
		"Season": params.Season,
		"SeasonType": params.SeasonTypePlayoffs,
		"DefPlayerID": params.DefPlayerIdNullable,
		"DefTeamID": params.DefTeamIdNullable,
		"OffPlayerID": params.OffPlayerIdNullable,
		"OffTeamID": params.OffTeamIdNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "leagueseasonmatchups", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch leagueseasonmatchups",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch leagueseasonmatchups: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from leagueseasonmatchups endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal leagueseasonmatchups response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched leagueseasonmatchups",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
