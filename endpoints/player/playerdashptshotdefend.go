package player

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayerDashPtShotDefendParams holds parameters for the PlayerDashPtShotDefend endpoint.
type PlayerDashPtShotDefendParams struct {
	TeamId string
	PlayerId string
	LastNGames string
	LeagueId string
	Month string
	OpponentTeamId string
	PerModeSimple string
	Period string
	Season string
	SeasonTypeAllStar string
	DateFromNullable string
	DateToNullable string
	GameSegmentNullable string
	LocationNullable string
	OutcomeNullable string
	SeasonSegmentNullable string
	VsConferenceNullable string
	VsDivisionNullable string
}

// GetPlayerDashPtShotDefend fetches data from the playerdashptshotdefend endpoint.
func (c *Client) GetPlayerDashPtShotDefend(ctx context.Context, params PlayerDashPtShotDefendParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playerdashptshotdefend")

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
	period := params.Period
	if period == "" {
		period = "0"
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
		"PlayerID": params.PlayerId,
		"LastNGames": lastNGames,
		"LeagueID": leagueId,
		"Month": month,
		"OpponentTeamID": opponentTeamId,
		"PerMode": perMode,
		"Period": period,
		"Season": season,
		"SeasonType": seasonType,
	}
	
	if params.DateFromNullable != "" {
		reqParams["DateFrom"] = params.DateFromNullable
	}
	if params.DateToNullable != "" {
		reqParams["DateTo"] = params.DateToNullable
	}
	if params.GameSegmentNullable != "" {
		reqParams["GameSegment"] = params.GameSegmentNullable
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

	resp, err := c.httpClient.SendRequest(ctx, "playerdashptshotdefend", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch playerdashptshotdefend",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch playerdashptshotdefend: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from playerdashptshotdefend endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal playerdashptshotdefend response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched playerdashptshotdefend",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}

