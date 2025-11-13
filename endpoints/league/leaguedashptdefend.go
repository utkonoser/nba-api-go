package league

import (
	"context"
	"fmt"
	"log/slog"
)

// LeagueDashPtDefendParams holds parameters for the LeagueDashPtDefend endpoint.
type LeagueDashPtDefendParams struct {
	DefenseCategory string
	LeagueId string
	PerModeSimple string
	Season string
	SeasonTypeAllStar string
	CollegeNullable string
	ConferenceNullable string
	CountryNullable string
	DateFromNullable string
	DateToNullable string
	DivisionNullable string
	DraftPickNullable string
	DraftYearNullable string
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
	PlayerIdNullable string
	PlayerPositionNullable string
	SeasonSegmentNullable string
	StarterBenchNullable string
	TeamIdNullable string
	VsConferenceNullable string
	VsDivisionNullable string
	WeightNullable string
}

// GetLeagueDashPtDefend fetches data from the leaguedashptdefend endpoint.
func (c *Client) GetLeagueDashPtDefend(ctx context.Context, params LeagueDashPtDefendParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching leaguedashptdefend")

	reqParams := map[string]string{
		"DefenseCategory": params.DefenseCategory,
		"LeagueID": params.LeagueId,
		"PerMode": params.PerModeSimple,
		"Season": params.Season,
		"SeasonType": params.SeasonTypeAllStar,
		"College": params.CollegeNullable,
		"Conference": params.ConferenceNullable,
		"Country": params.CountryNullable,
		"DateFrom": params.DateFromNullable,
		"DateTo": params.DateToNullable,
		"Division": params.DivisionNullable,
		"DraftPick": params.DraftPickNullable,
		"DraftYear": params.DraftYearNullable,
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
		"PlayerID": params.PlayerIdNullable,
		"PlayerPosition": params.PlayerPositionNullable,
		"SeasonSegment": params.SeasonSegmentNullable,
		"StarterBench": params.StarterBenchNullable,
		"TeamID": params.TeamIdNullable,
		"VsConference": params.VsConferenceNullable,
		"VsDivision": params.VsDivisionNullable,
		"Weight": params.WeightNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "leaguedashptdefend", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch leaguedashptdefend",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch leaguedashptdefend: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from leaguedashptdefend endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal leaguedashptdefend response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched leaguedashptdefend",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
