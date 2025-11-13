package player

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayerVsPlayerParams holds parameters for the PlayerVsPlayer endpoint.
type PlayerVsPlayerParams struct {
	VsPlayerId string
	PlayerId string
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
	SeasonSegmentNullable string
	VsConferenceNullable string
	VsDivisionNullable string
}

// GetPlayerVsPlayer fetches data from the playervsplayer endpoint.
func (c *Client) GetPlayerVsPlayer(ctx context.Context, params PlayerVsPlayerParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playervsplayer")

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
		"PlayerID": params.PlayerId,
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
	if params.SeasonSegmentNullable != "" {
		reqParams["SeasonSegment"] = params.SeasonSegmentNullable
	}
	if params.VsConferenceNullable != "" {
		reqParams["VsConference"] = params.VsConferenceNullable
	}
	if params.VsDivisionNullable != "" {
		reqParams["VsDivision"] = params.VsDivisionNullable
	}

	resp, err := c.httpClient.SendRequest(ctx, "playervsplayer", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch playervsplayer",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch playervsplayer: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from playervsplayer endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal playervsplayer response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched playervsplayer",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
