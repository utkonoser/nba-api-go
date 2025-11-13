package team

import (
	"context"
	"fmt"
	"log/slog"
)

// TeamDashboardByShootingSplitsParams holds parameters for the TeamDashboardByShootingSplits endpoint.
type TeamDashboardByShootingSplitsParams struct {
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
	SeasonTypeAllStar string
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

// GetTeamDashboardByShootingSplits fetches data from the teamdashboardbyshootingsplits endpoint.
func (c *Client) GetTeamDashboardByShootingSplits(ctx context.Context, params TeamDashboardByShootingSplitsParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching teamdashboardbyshootingsplits")

	reqParams := map[string]string{
		"TeamID": params.TeamId,
		"LastNGames": params.LastNGames,
		"MeasureType": params.MeasureTypeDetailedDefense,
		"Month": params.Month,
		"OpponentTeamID": params.OpponentTeamId,
		"PaceAdjust": params.PaceAdjust,
		"PerMode": params.PerModeDetailed,
		"Period": params.Period,
		"PlusMinus": params.PlusMinus,
		"Rank": params.Rank,
		"Season": params.Season,
		"SeasonType": params.SeasonTypeAllStar,
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

	resp, err := c.httpClient.SendRequest(ctx, "teamdashboardbyshootingsplits", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch teamdashboardbyshootingsplits",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch teamdashboardbyshootingsplits: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from teamdashboardbyshootingsplits endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal teamdashboardbyshootingsplits response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched teamdashboardbyshootingsplits",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
