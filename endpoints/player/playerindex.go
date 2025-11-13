package player

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayerIndexParams holds parameters for the PlayerIndex endpoint.
type PlayerIndexParams struct {
	ActiveNullable string
	AllStarNullable string
	CollegeNullable string
	CountryNullable string
	DraftPickNullable string
	DraftYearNullable string
	HeightNullable string
	PlayerPositionAbbreviationNullable string
	HistoricalNullable string
	LeagueId string
	Season string
	TeamIdNullable string
	WeightNullable string
}

// GetPlayerIndex fetches data from the playerindex endpoint.
func (c *Client) GetPlayerIndex(ctx context.Context, params PlayerIndexParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playerindex")

	// Set defaults for required parameters
	leagueId := params.LeagueId
	if leagueId == "" {
		leagueId = "00"
	}
	season := params.Season
	if season == "" {
		season = "2023-24"
	}

	reqParams := map[string]string{
		"LeagueID": leagueId,
		"Season": season,
	}
	
	// Add nullable parameters only if they are not empty
	if params.ActiveNullable != "" {
		reqParams["Active"] = params.ActiveNullable
	}
	if params.AllStarNullable != "" {
		reqParams["AllStar"] = params.AllStarNullable
	}
	if params.CollegeNullable != "" {
		reqParams["College"] = params.CollegeNullable
	}
	if params.CountryNullable != "" {
		reqParams["Country"] = params.CountryNullable
	}
	if params.DraftPickNullable != "" {
		reqParams["DraftPick"] = params.DraftPickNullable
	}
	if params.DraftYearNullable != "" {
		reqParams["DraftYear"] = params.DraftYearNullable
	}
	if params.HeightNullable != "" {
		reqParams["Height"] = params.HeightNullable
	}
	if params.PlayerPositionAbbreviationNullable != "" {
		reqParams["PlayerPosition"] = params.PlayerPositionAbbreviationNullable
	}
	if params.HistoricalNullable != "" {
		reqParams["Historical"] = params.HistoricalNullable
	}
	if params.TeamIdNullable != "" {
		reqParams["TeamID"] = params.TeamIdNullable
	}
	if params.WeightNullable != "" {
		reqParams["Weight"] = params.WeightNullable
	}

	resp, err := c.httpClient.SendRequest(ctx, "playerindex", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch playerindex",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch playerindex: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from playerindex endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal playerindex response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched playerindex",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}

