package league

import (
	"context"
	"fmt"
	"log/slog"
)

// LeagueDashPlayerPtShotParams holds parameters for the LeagueDashPlayerPtShot endpoint.
type LeagueDashPlayerPtShotParams struct {
	LeagueId string
	PerModeSimple string
	Season string
	SeasonTypeAllStar string
	CloseDefDistRangeNullable string
	CollegeNullable string
	ConferenceNullable string
	CountryNullable string
	DateFromNullable string
	DateToNullable string
	DivisionNullable string
	DraftPickNullable string
	DraftYearNullable string
	DribbleRangeNullable string
	GameSegmentNullable string
	GeneralRangeNullable string
	HeightNullable string
	LastNGamesNullable string
	LocationNullable string
	MonthNullable string
	OpponentTeamIdNullable string
	OutcomeNullable string
	PoRoundNullable string
	PeriodNullable string
	PlayerExperienceNullable string
	PlayerPositionNullable string
	SeasonSegmentNullable string
	ShotClockRangeNullable string
	ShotDistRangeNullable string
	StarterBenchNullable string
	TeamIdNullable string
	TouchTimeRangeNullable string
	VsConferenceNullable string
	VsDivisionNullable string
	WeightNullable string
}

// GetLeagueDashPlayerPtShot fetches data from the leaguedashplayerptshot endpoint.
func (c *Client) GetLeagueDashPlayerPtShot(ctx context.Context, params LeagueDashPlayerPtShotParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching leaguedashplayerptshot")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
		"PerMode": params.PerModeSimple,
		"Season": params.Season,
		"SeasonType": params.SeasonTypeAllStar,
		"CloseDefDistRange": params.CloseDefDistRangeNullable,
		"College": params.CollegeNullable,
		"Conference": params.ConferenceNullable,
		"Country": params.CountryNullable,
		"DateFrom": params.DateFromNullable,
		"DateTo": params.DateToNullable,
		"Division": params.DivisionNullable,
		"DraftPick": params.DraftPickNullable,
		"DraftYear": params.DraftYearNullable,
		"DribbleRange": params.DribbleRangeNullable,
		"GameSegment": params.GameSegmentNullable,
		"GeneralRange": params.GeneralRangeNullable,
		"Height": params.HeightNullable,
		"LastNGames": params.LastNGamesNullable,
		"Location": params.LocationNullable,
		"Month": params.MonthNullable,
		"OpponentTeamID": params.OpponentTeamIdNullable,
		"Outcome": params.OutcomeNullable,
		"PORound": params.PoRoundNullable,
		"Period": params.PeriodNullable,
		"PlayerExperience": params.PlayerExperienceNullable,
		"PlayerPosition": params.PlayerPositionNullable,
		"SeasonSegment": params.SeasonSegmentNullable,
		"ShotClockRange": params.ShotClockRangeNullable,
		"ShotDistRange": params.ShotDistRangeNullable,
		"StarterBench": params.StarterBenchNullable,
		"TeamID": params.TeamIdNullable,
		"TouchTimeRange": params.TouchTimeRangeNullable,
		"VsConference": params.VsConferenceNullable,
		"VsDivision": params.VsDivisionNullable,
		"Weight": params.WeightNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "leaguedashplayerptshot", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch leaguedashplayerptshot",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch leaguedashplayerptshot: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from leaguedashplayerptshot endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal leaguedashplayerptshot response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched leaguedashplayerptshot",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
