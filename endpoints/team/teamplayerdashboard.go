package team

import (
	"context"
	"fmt"
	"log/slog"
)

// TeamPlayerDashboardParams holds parameters for the TeamPlayerDashboard endpoint.
type TeamPlayerDashboardParams struct {
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

// GetTeamPlayerDashboard fetches data from the teamplayerdashboard endpoint.
func (c *Client) GetTeamPlayerDashboard(ctx context.Context, params TeamPlayerDashboardParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching teamplayerdashboard")

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

	resp, err := c.httpClient.SendRequest(ctx, "teamplayerdashboard", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch teamplayerdashboard",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch teamplayerdashboard: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from teamplayerdashboard endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal teamplayerdashboard response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched teamplayerdashboard",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
