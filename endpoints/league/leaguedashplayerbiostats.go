package league

import (
	"context"
	"fmt"
	"log/slog"
)

// LeagueDashPlayerBioStatsParams holds parameters for the LeagueDashPlayerBioStats endpoint.
type LeagueDashPlayerBioStatsParams struct {
	LeagueId string
	PerModeSimple string
	Season string
	SeasonTypeAllStar string
	CollegeNullable string
	ConferenceNullable string
	CountryNullable string
	DateFromNullable string
	DateToNullable string
	DivisionSimpleNullable string
	DraftPickNullable string
	DraftYearNullable string
	GameScopeSimpleNullable string
	GameSegmentNullable string
	HeightNullable string
	LastNGamesNullable string
	LocationNullable string
	MonthNullable string
	OpponentTeamIdNullable string
	OutcomeNullable string
	PoRoundNullable string
	PeriodNullable string
	PlayerExperienceNullable string
	PlayerPositionAbbreviationNullable string
	SeasonSegmentNullable string
	ShotClockRangeNullable string
	StarterBenchNullable string
	TeamIdNullable string
	VsConferenceNullable string
	VsDivisionNullable string
	WeightNullable string
}

// GetLeagueDashPlayerBioStats fetches data from the leaguedashplayerbiostats endpoint.
func (c *Client) GetLeagueDashPlayerBioStats(ctx context.Context, params LeagueDashPlayerBioStatsParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching leaguedashplayerbiostats")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
		"PerMode": params.PerModeSimple,
		"Season": params.Season,
		"SeasonType": params.SeasonTypeAllStar,
		"College": params.CollegeNullable,
		"Conference": params.ConferenceNullable,
		"Country": params.CountryNullable,
		"DateFrom": params.DateFromNullable,
		"DateTo": params.DateToNullable,
		"Division": params.DivisionSimpleNullable,
		"DraftPick": params.DraftPickNullable,
		"DraftYear": params.DraftYearNullable,
		"GameScope": params.GameScopeSimpleNullable,
		"GameSegment": params.GameSegmentNullable,
		"Height": params.HeightNullable,
		"LastNGames": params.LastNGamesNullable,
		"Location": params.LocationNullable,
		"Month": params.MonthNullable,
		"OpponentTeamID": params.OpponentTeamIdNullable,
		"Outcome": params.OutcomeNullable,
		"PORound": params.PoRoundNullable,
		"Period": params.PeriodNullable,
		"PlayerExperience": params.PlayerExperienceNullable,
		"PlayerPosition": params.PlayerPositionAbbreviationNullable,
		"SeasonSegment": params.SeasonSegmentNullable,
		"ShotClockRange": params.ShotClockRangeNullable,
		"StarterBench": params.StarterBenchNullable,
		"TeamID": params.TeamIdNullable,
		"VsConference": params.VsConferenceNullable,
		"VsDivision": params.VsDivisionNullable,
		"Weight": params.WeightNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "leaguedashplayerbiostats", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch leaguedashplayerbiostats",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch leaguedashplayerbiostats: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from leaguedashplayerbiostats endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal leaguedashplayerbiostats response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched leaguedashplayerbiostats",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
