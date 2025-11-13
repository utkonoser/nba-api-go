package player

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayerCompareParams holds parameters for the PlayerCompare endpoint.
type PlayerCompareParams struct {
	VsPlayerIdList string
	PlayerIdList string
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
	SeasonType string
	ConferenceNullable string
	DivisionSimpleNullable string
	GameSegmentNullable string
	LeagueIdNullable string
	LocationNullable string
	OutcomeNullable string
	PoRoundNullable string
	SeasonSegmentNullable string
	ShotClockRangeNullable string
	VsDivisionNullable string
}

// GetPlayerCompare fetches data from the playercompare endpoint.
func (c *Client) GetPlayerCompare(ctx context.Context, params PlayerCompareParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playercompare")

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
	seasonType := params.SeasonType
	if seasonType == "" {
		seasonType = "Regular Season"
	}

	reqParams := map[string]string{
		"VsPlayerIDList": params.VsPlayerIdList,
		"PlayerIDList": params.PlayerIdList,
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
	
	if params.ConferenceNullable != "" {
		reqParams["Conference"] = params.ConferenceNullable
	}
	if params.DivisionSimpleNullable != "" {
		reqParams["Division"] = params.DivisionSimpleNullable
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
	if params.VsDivisionNullable != "" {
		reqParams["VsDivision"] = params.VsDivisionNullable
	}

	resp, err := c.httpClient.SendRequest(ctx, "playercompare", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch playercompare",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch playercompare: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from playercompare endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal playercompare response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched playercompare",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}

