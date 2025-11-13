package team

import (
	"context"
	"fmt"
	"log/slog"
)

// TeamVsPlayerParams holds parameters for the TeamVsPlayer endpoint.
type TeamVsPlayerParams struct {
	VsPlayerId string
	TeamId string
	LastNGames string
	MeasureTypeDetailedDefense string
	Month string
	OpponentTeamId string
	PaceAdjust string
	PerModeDetailed string
	Period string
	PlusMinus string
	Rank string
	Season string
	SeasonTypePlayoffs string
	DateFromNullable string
	DateToNullable string
	GameSegmentNullable string
	LeagueIdNullable string
	LocationNullable string
	OutcomeNullable string
	PlayerIdNullable string
	SeasonSegmentNullable string
	VsConferenceNullable string
	VsDivisionNullable string
}

// GetTeamVsPlayer fetches data from the teamvsplayer endpoint.
func (c *Client) GetTeamVsPlayer(ctx context.Context, params TeamVsPlayerParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching teamvsplayer")

	// Set defaults for required parameters
	lastNGames := params.LastNGames
	if lastNGames == "" {
		lastNGames = "0"
	}
	measureType := params.MeasureTypeDetailedDefense
	if measureType == "" {
		measureType = "Base"
	}
	month := params.Month
	if month == "" {
		month = "0"
	}
	opponentTeamId := params.OpponentTeamId
	if opponentTeamId == "" {
		opponentTeamId = "0"
	}
	paceAdjust := params.PaceAdjust
	if paceAdjust == "" {
		paceAdjust = "N"
	}
	perMode := params.PerModeDetailed
	if perMode == "" {
		perMode = "Totals"
	}
	period := params.Period
	if period == "" {
		period = "0"
	}
	plusMinus := params.PlusMinus
	if plusMinus == "" {
		plusMinus = "N"
	}
	rank := params.Rank
	if rank == "" {
		rank = "N"
	}
	season := params.Season
	if season == "" {
		season = "2023-24"
	}
	seasonType := params.SeasonTypePlayoffs
	if seasonType == "" {
		seasonType = "Regular Season"
	}

	reqParams := map[string]string{
		"VsPlayerID": params.VsPlayerId,
		"TeamID": params.TeamId,
		"LastNGames": lastNGames,
		"MeasureType": measureType,
		"Month": month,
		"OpponentTeamID": opponentTeamId,
		"PaceAdjust": paceAdjust,
		"PerMode": perMode,
		"Period": period,
		"PlusMinus": plusMinus,
		"Rank": rank,
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
	if params.GameSegmentNullable != "" {
		reqParams["GameSegment"] = params.GameSegmentNullable
	}
	if params.LeagueIdNullable != "" {
		reqParams["LeagueID"] = params.LeagueIdNullable
	}
	if params.LocationNullable != "" {
		reqParams["Location"] = params.LocationNullable
	}
	if params.OutcomeNullable != "" {
		reqParams["Outcome"] = params.OutcomeNullable
	}
	if params.PlayerIdNullable != "" {
		reqParams["PlayerID"] = params.PlayerIdNullable
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

	resp, err := c.httpClient.SendRequest(ctx, "teamvsplayer", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch teamvsplayer",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch teamvsplayer: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from teamvsplayer endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal teamvsplayer response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched teamvsplayer",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
