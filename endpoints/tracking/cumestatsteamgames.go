package tracking

import (
	"context"
	"fmt"
	"log/slog"
)

// CumeStatsTeamGamesParams holds parameters for the CumeStatsTeamGames endpoint.
type CumeStatsTeamGamesParams struct {
	TeamId string
	LeagueId string
	Season string
	SeasonTypeAllStar string
	LocationNullable string
	OutcomeNullable string
	SeasonIdNullable string
	VsConferenceNullable string
	VsDivisionNullable string
	VsTeamIdNullable string
}

// GetCumeStatsTeamGames fetches data from the cumestatsteamgames endpoint.
func (c *Client) GetCumeStatsTeamGames(ctx context.Context, params CumeStatsTeamGamesParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching cumestatsteamgames")

	reqParams := map[string]string{
		"TeamID": params.TeamId,
		"LeagueID": params.LeagueId,
		"Season": params.Season,
		"SeasonType": params.SeasonTypeAllStar,
		"Location": params.LocationNullable,
		"Outcome": params.OutcomeNullable,
		"SeasonID": params.SeasonIdNullable,
		"VsConference": params.VsConferenceNullable,
		"VsDivision": params.VsDivisionNullable,
		"VsTeamID": params.VsTeamIdNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "cumestatsteamgames", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch cumestatsteamgames",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch cumestatsteamgames: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from cumestatsteamgames endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal cumestatsteamgames response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched cumestatsteamgames",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
