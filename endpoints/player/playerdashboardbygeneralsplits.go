package player

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayerDashboardByGeneralSplitsParams holds parameters for the PlayerDashboardByGeneralSplits endpoint.
type PlayerDashboardByGeneralSplitsParams struct {
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

// GetPlayerDashboardByGeneralSplits fetches data from the playerdashboardbygeneralsplits endpoint.
func (c *Client) GetPlayerDashboardByGeneralSplits(ctx context.Context, params PlayerDashboardByGeneralSplitsParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playerdashboardbygeneralsplits")

	reqParams := map[string]string{
		"PlayerID": params.PlayerId,
	}
	
	// Add optional parameters only if they are not empty
	if params.LastNGames != "" {
		reqParams["LastNGames"] = params.LastNGames
	}
	if params.MeasureTypeDetailed != "" {
		reqParams["MeasureType"] = params.MeasureTypeDetailed
	}
	if params.Month != "" {
		reqParams["Month"] = params.Month
	}
	if params.OpponentTeamId != "" {
		reqParams["OpponentTeamID"] = params.OpponentTeamId
	}
	if params.PaceAdjust != "" {
		reqParams["PaceAdjust"] = params.PaceAdjust
	}
	if params.PerModeDetailed != "" {
		reqParams["PerMode"] = params.PerModeDetailed
	}
	if params.Period != "" {
		reqParams["Period"] = params.Period
	}
	if params.PlusMinus != "" {
		reqParams["PlusMinus"] = params.PlusMinus
	}
	if params.Rank != "" {
		reqParams["Rank"] = params.Rank
	}
	if params.Season != "" {
		reqParams["Season"] = params.Season
	}
	if params.SeasonTypePlayoffs != "" {
		reqParams["SeasonType"] = params.SeasonTypePlayoffs
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

	resp, err := c.httpClient.SendRequest(ctx, "playerdashboardbygeneralsplits", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch playerdashboardbygeneralsplits",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch playerdashboardbygeneralsplits: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from playerdashboardbygeneralsplits endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal playerdashboardbygeneralsplits response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched playerdashboardbygeneralsplits",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
