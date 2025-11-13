package team

import (
	"context"
	"fmt"
	"log/slog"
)

// TeamDashPtPassParams holds parameters for the TeamDashPtPass endpoint.
type TeamDashPtPassParams struct {
	TeamId string
	LastNGames string
	LeagueId string
	Month string
	OpponentTeamId string
	PerModeSimple string
	Season string
	SeasonTypeAllStar string
	DateFromNullable string
	DateToNullable string
	LocationNullable string
	OutcomeNullable string
	SeasonSegmentNullable string
	VsConferenceNullable string
	VsDivisionNullable string
}

// GetTeamDashPtPass fetches data from the teamdashptpass endpoint.
func (c *Client) GetTeamDashPtPass(ctx context.Context, params TeamDashPtPassParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching teamdashptpass")

	// Set defaults for required parameters
	lastNGames := params.LastNGames
	if lastNGames == "" {
		lastNGames = "0"
	}
	leagueId := params.LeagueId
	if leagueId == "" {
		leagueId = "00"
	}
	month := params.Month
	if month == "" {
		month = "0"
	}
	opponentTeamId := params.OpponentTeamId
	if opponentTeamId == "" {
		opponentTeamId = "0"
	}
	perMode := params.PerModeSimple
	if perMode == "" {
		perMode = "Totals"
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
		"LastNGames": lastNGames,
		"LeagueID": leagueId,
		"Month": month,
		"OpponentTeamID": opponentTeamId,
		"PerMode": perMode,
		"Season": season,
		"SeasonType": seasonType,
	}
	
	// Add nullable parameters only if they are not empty
	if params.DateFromNullable != "" {
		reqParams["DateFrom"] = params.DateFromNullable
	}
	if params.DateToNullable != "" {
		reqParams["DateTo"] = params.DateToNullable
	}
	if params.LocationNullable != "" {
		reqParams["Location"] = params.LocationNullable
	}
	if params.OutcomeNullable != "" {
		reqParams["Outcome"] = params.OutcomeNullable
	}
	if params.SeasonSegmentNullable != "" {
		reqParams["SeasonSegment"] = params.SeasonSegmentNullable
	}
	if params.VsConferenceNullable != "" {
		reqParams["VsConference"] = params.VsConferenceNullable
	}
	if params.VsDivisionNullable != "" {
		reqParams["VsDivision"] = params.VsDivisionNullable
	}

	resp, err := c.httpClient.SendRequest(ctx, "teamdashptpass", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch teamdashptpass",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch teamdashptpass: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from teamdashptpass endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal teamdashptpass response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched teamdashptpass",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
