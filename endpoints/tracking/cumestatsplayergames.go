package tracking

import (
	"context"
	"fmt"
	"log/slog"
)

// CumeStatsPlayerGamesParams holds parameters for the CumeStatsPlayerGames endpoint.
type CumeStatsPlayerGamesParams struct {
	PlayerId string
	LeagueId string
	Season string
	SeasonTypeAllStar string
	LocationNullable string
	OutcomeNullable string
	VsConferenceNullable string
	VsDivisionNullable string
	VsTeamIdNullable string
}

// GetCumeStatsPlayerGames fetches data from the cumestatsplayergames endpoint.
func (c *Client) GetCumeStatsPlayerGames(ctx context.Context, params CumeStatsPlayerGamesParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching cumestatsplayergames")

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
		"PlayerID": params.PlayerId,
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
	if params.VsConferenceNullable != "" {
		reqParams["VsConference"] = params.VsConferenceNullable
	}
	if params.VsDivisionNullable != "" {
		reqParams["VsDivision"] = params.VsDivisionNullable
	}
	if params.VsTeamIdNullable != "" {
		reqParams["VsTeamID"] = params.VsTeamIdNullable
	}

	resp, err := c.httpClient.SendRequest(ctx, "cumestatsplayergames", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch cumestatsplayergames",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch cumestatsplayergames: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from cumestatsplayergames endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal cumestatsplayergames response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched cumestatsplayergames",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
