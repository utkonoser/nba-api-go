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

	// Set defaults for required parameters
	leagueId := params.LeagueId
	if leagueId == "" {
		leagueId = "00"
	}
	season := params.Season
	if season == "" {
		season = "2023-24"
	}
	seasonType := params.SeasonTypeAllStar
	if seasonType == "" {
		seasonType = "Regular Season"
	}

	reqParams := map[string]string{
		"TeamID": params.TeamId,
		"LeagueID": leagueId,
		"Season": season,
		"SeasonType": seasonType,
	}
	
	// Add nullable parameters only if they are not empty
	if params.LocationNullable != "" {
		reqParams["Location"] = params.LocationNullable
	}
	if params.OutcomeNullable != "" {
		reqParams["Outcome"] = params.OutcomeNullable
	}
	if params.SeasonIdNullable != "" {
		reqParams["SeasonID"] = params.SeasonIdNullable
	}
	if params.VsConferenceNullable != "" {
		reqParams["VsConference"] = params.VsConferenceNullable
	}
	if params.VsDivisionNullable != "" {
		reqParams["VsDivision"] = params.VsDivisionNullable
	}
	if params.VsTeamIdNullable != "" {
		reqParams["VsTeamID"] = params.VsTeamIdNullable
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
