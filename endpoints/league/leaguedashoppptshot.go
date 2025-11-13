package league

import (
	"context"
	"fmt"
	"log/slog"
)

// LeagueDashOppPtShotParams holds parameters for the LeagueDashOppPtShot endpoint.
type LeagueDashOppPtShotParams struct {
	LeagueId string
	PerModeSimple string
	Season string
	SeasonTypeAllStar string
	CloseDefDistRangeNullable string
	ConferenceNullable string
	DateFromNullable string
	DateToNullable string
	DivisionNullable string
	DribbleRangeNullable string
	GameSegmentNullable string
	GeneralRangeNullable string
	LastNGamesNullable string
	LocationNullable string
	MonthNullable string
	OpponentTeamIdNullable string
	OutcomeNullable string
	PoRoundNullable string
	PeriodNullable string
	SeasonSegmentNullable string
	ShotClockRangeNullable string
	ShotDistRangeNullable string
	TeamIdNullable string
	TouchTimeRangeNullable string
	VsConferenceNullable string
	VsDivisionNullable string
}

// GetLeagueDashOppPtShot fetches data from the leaguedashoppptshot endpoint.
func (c *Client) GetLeagueDashOppPtShot(ctx context.Context, params LeagueDashOppPtShotParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching leaguedashoppptshot")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
		"PerMode": params.PerModeSimple,
		"Season": params.Season,
		"SeasonType": params.SeasonTypeAllStar,
		"CloseDefDistRange": params.CloseDefDistRangeNullable,
		"Conference": params.ConferenceNullable,
		"DateFrom": params.DateFromNullable,
		"DateTo": params.DateToNullable,
		"Division": params.DivisionNullable,
		"DribbleRange": params.DribbleRangeNullable,
		"GameSegment": params.GameSegmentNullable,
		"GeneralRange": params.GeneralRangeNullable,
		"LastNGames": params.LastNGamesNullable,
		"Location": params.LocationNullable,
		"Month": params.MonthNullable,
		"OpponentTeamID": params.OpponentTeamIdNullable,
		"Outcome": params.OutcomeNullable,
		"PORound": params.PoRoundNullable,
		"Period": params.PeriodNullable,
		"SeasonSegment": params.SeasonSegmentNullable,
		"ShotClockRange": params.ShotClockRangeNullable,
		"ShotDistRange": params.ShotDistRangeNullable,
		"TeamID": params.TeamIdNullable,
		"TouchTimeRange": params.TouchTimeRangeNullable,
		"VsConference": params.VsConferenceNullable,
		"VsDivision": params.VsDivisionNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "leaguedashoppptshot", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch leaguedashoppptshot",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch leaguedashoppptshot: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from leaguedashoppptshot endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal leaguedashoppptshot response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched leaguedashoppptshot",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
