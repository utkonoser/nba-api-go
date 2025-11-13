package league

import (
	"context"
	"fmt"
	"log/slog"
)

// LeagueDashTeamPtShotParams holds parameters for the LeagueDashTeamPtShot endpoint.
type LeagueDashTeamPtShotParams struct {
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

// GetLeagueDashTeamPtShot fetches data from the leaguedashteamptshot endpoint.
func (c *Client) GetLeagueDashTeamPtShot(ctx context.Context, params LeagueDashTeamPtShotParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching leaguedashteamptshot")

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

	resp, err := c.httpClient.SendRequest(ctx, "leaguedashteamptshot", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch leaguedashteamptshot",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch leaguedashteamptshot: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from leaguedashteamptshot endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal leaguedashteamptshot response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched leaguedashteamptshot",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
