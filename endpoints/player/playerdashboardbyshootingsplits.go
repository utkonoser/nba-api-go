package player

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayerDashboardByShootingSplitsParams holds parameters for the PlayerDashboardByShootingSplits endpoint.
type PlayerDashboardByShootingSplitsParams struct {
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

// GetPlayerDashboardByShootingSplits fetches data from the playerdashboardbyshootingsplits endpoint.
func (c *Client) GetPlayerDashboardByShootingSplits(ctx context.Context, params PlayerDashboardByShootingSplitsParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playerdashboardbyshootingsplits")

	reqParams := map[string]string{
		"PlayerID": params.PlayerId,
		"LastNGames": params.LastNGames,
		"MeasureType": params.MeasureTypeDetailed,
		"Month": params.Month,
		"OpponentTeamID": params.OpponentTeamId,
		"PaceAdjust": params.PaceAdjust,
		"PerMode": params.PerModeDetailed,
		"Period": params.Period,
		"PlusMinus": params.PlusMinus,
		"Rank": params.Rank,
		"Season": params.Season,
		"SeasonType": params.SeasonTypePlayoffs,
		"DateFrom": params.DateFromNullable,
		"DateTo": params.DateToNullable,
		"GameSegment": params.GameSegmentNullable,
		"LeagueID": params.LeagueIdNullable,
		"Location": params.LocationNullable,
		"Outcome": params.OutcomeNullable,
		"PORound": params.PoRoundNullable,
		"SeasonSegment": params.SeasonSegmentNullable,
		"ShotClockRange": params.ShotClockRangeNullable,
		"VsConference": params.VsConferenceNullable,
		"VsDivision": params.VsDivisionNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "playerdashboardbyshootingsplits", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch playerdashboardbyshootingsplits",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch playerdashboardbyshootingsplits: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from playerdashboardbyshootingsplits endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal playerdashboardbyshootingsplits response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched playerdashboardbyshootingsplits",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
