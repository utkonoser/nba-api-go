package league

import (
	"context"
	"fmt"
	"log/slog"
)

// LeagueDashPtStatsParams holds parameters for the LeagueDashPtStats endpoint.
type LeagueDashPtStatsParams struct {
	LastNGames string
	Month string
	OpponentTeamId string
	PerModeSimple string
	PlayerOrTeam string
	PtMeasureType string
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
	HeightNullable string
	LeagueIdNullable string
	LocationNullable string
	OutcomeNullable string
	PoRoundNullable string
	PlayerExperienceNullable string
	PlayerPositionAbbreviationNullable string
	SeasonSegmentNullable string
	StarterBenchNullable string
	TeamIdNullable string
	VsConferenceNullable string
	VsDivisionNullable string
	WeightNullable string
}

// GetLeagueDashPtStats fetches data from the leaguedashptstats endpoint.
func (c *Client) GetLeagueDashPtStats(ctx context.Context, params LeagueDashPtStatsParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching leaguedashptstats")

	reqParams := map[string]string{
		"LastNGames": params.LastNGames,
		"Month": params.Month,
		"OpponentTeamID": params.OpponentTeamId,
		"PerMode": params.PerModeSimple,
		"PlayerOrTeam": params.PlayerOrTeam,
		"PtMeasureType": params.PtMeasureType,
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
		"Height": params.HeightNullable,
		"LeagueID": params.LeagueIdNullable,
		"Location": params.LocationNullable,
		"Outcome": params.OutcomeNullable,
		"PORound": params.PoRoundNullable,
		"PlayerExperience": params.PlayerExperienceNullable,
		"PlayerPosition": params.PlayerPositionAbbreviationNullable,
		"SeasonSegment": params.SeasonSegmentNullable,
		"StarterBench": params.StarterBenchNullable,
		"TeamID": params.TeamIdNullable,
		"VsConference": params.VsConferenceNullable,
		"VsDivision": params.VsDivisionNullable,
		"Weight": params.WeightNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "leaguedashptstats", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch leaguedashptstats",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch leaguedashptstats: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from leaguedashptstats endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal leaguedashptstats response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched leaguedashptstats",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
