package player

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayerDashboardByYearOverYearParams holds parameters for the PlayerDashboardByYearOverYear endpoint.
type PlayerDashboardByYearOverYearParams struct {
	PlayerId string
	LastNGames string
	MeasureTypeDetailed string
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
	PoRoundNullable string
	SeasonSegmentNullable string
	ShotClockRangeNullable string
	VsConferenceNullable string
	VsDivisionNullable string
}

// GetPlayerDashboardByYearOverYear fetches data from the playerdashboardbyyearoveryear endpoint.
func (c *Client) GetPlayerDashboardByYearOverYear(ctx context.Context, params PlayerDashboardByYearOverYearParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playerdashboardbyyearoveryear")

	// Set defaults for required parameters (matching Python nba_api behavior)
	lastNGames := params.LastNGames
	if lastNGames == "" {
		lastNGames = "0"
	}
	measureType := params.MeasureTypeDetailed
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
		season = "2023-24" // Default current season
	}
	seasonType := params.SeasonTypePlayoffs
	if seasonType == "" {
		seasonType = "Regular Season"
	}

	reqParams := map[string]string{
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
	if params.PoRoundNullable != "" {
		reqParams["PORound"] = params.PoRoundNullable
	}
	if params.SeasonSegmentNullable != "" {
		reqParams["SeasonSegment"] = params.SeasonSegmentNullable
	}
	if params.ShotClockRangeNullable != "" {
		reqParams["ShotClockRange"] = params.ShotClockRangeNullable
	}
	if params.VsConferenceNullable != "" {
		reqParams["VsConference"] = params.VsConferenceNullable
	}
	if params.VsDivisionNullable != "" {
		reqParams["VsDivision"] = params.VsDivisionNullable
	}

	resp, err := c.httpClient.SendRequest(ctx, "playerdashboardbyyearoveryear", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch playerdashboardbyyearoveryear",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch playerdashboardbyyearoveryear: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from playerdashboardbyyearoveryear endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal playerdashboardbyyearoveryear response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched playerdashboardbyyearoveryear",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}

