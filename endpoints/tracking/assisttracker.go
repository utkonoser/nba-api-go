package tracking

import (
	"context"
	"fmt"
	"log/slog"
)

// AssistTrackerParams holds parameters for the AssistTracker endpoint.
type AssistTrackerParams struct {
	CollegeNullable string
	ConferenceNullable string
	CountryNullable string
	DateFromNullable string
	DateToNullable string
	DivisionSimpleNullable string
	DraftPickNullable string
	DraftYearNullable string
	GameScopeSimpleNullable string
	HeightNullable string
	LastNGamesNullable string
	LeagueIdNullable string
	LocationNullable string
	MonthNullable string
	OpponentTeamIdNullable string
	OutcomeNullable string
	PoRoundNullable string
	PerModeSimpleNullable string
	PlayerExperienceNullable string
	PlayerPositionAbbreviationNullable string
	SeasonNullable string
	SeasonSegmentNullable string
	SeasonTypeAllStarNullable string
	StarterBenchNullable string
	TeamIdNullable string
	VsConferenceNullable string
	VsDivisionNullable string
	WeightNullable string
}

// GetAssistTracker fetches data from the assisttracker endpoint.
func (c *Client) GetAssistTracker(ctx context.Context, params AssistTrackerParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching assisttracker")

	reqParams := map[string]string{
		"College": params.CollegeNullable,
		"Conference": params.ConferenceNullable,
		"Country": params.CountryNullable,
		"DateFrom": params.DateFromNullable,
		"DateTo": params.DateToNullable,
		"Division": params.DivisionSimpleNullable,
		"DraftPick": params.DraftPickNullable,
		"DraftYear": params.DraftYearNullable,
		"GameScope": params.GameScopeSimpleNullable,
		"Height": params.HeightNullable,
		"LastNGames": params.LastNGamesNullable,
		"LeagueID": params.LeagueIdNullable,
		"Location": params.LocationNullable,
		"Month": params.MonthNullable,
		"OpponentTeamID": params.OpponentTeamIdNullable,
		"Outcome": params.OutcomeNullable,
		"PORound": params.PoRoundNullable,
		"PerMode": params.PerModeSimpleNullable,
		"PlayerExperience": params.PlayerExperienceNullable,
		"PlayerPosition": params.PlayerPositionAbbreviationNullable,
		"Season": params.SeasonNullable,
		"SeasonSegment": params.SeasonSegmentNullable,
		"SeasonType": params.SeasonTypeAllStarNullable,
		"StarterBench": params.StarterBenchNullable,
		"TeamID": params.TeamIdNullable,
		"VsConference": params.VsConferenceNullable,
		"VsDivision": params.VsDivisionNullable,
		"Weight": params.WeightNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "assisttracker", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch assisttracker",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch assisttracker: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from assisttracker endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal assisttracker response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched assisttracker",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
