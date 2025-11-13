package league

import (
	"context"
	"fmt"
	"log/slog"
)

// LeagueDashPlayerShotLocationsParams holds parameters for the LeagueDashPlayerShotLocations endpoint.
type LeagueDashPlayerShotLocationsParams struct {
	DistanceRange string
	LastNGames string
	MeasureTypeSimple string
	Month string
	OpponentTeamId string
	PaceAdjust string
	PerModeDetailed string
	Period string
	PlusMinus string
	Rank string
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
	LeagueIdNullable string
	LocationNullable string
	OutcomeNullable string
	PoRoundNullable string
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

// GetLeagueDashPlayerShotLocations fetches data from the leaguedashplayershotlocations endpoint.
func (c *Client) GetLeagueDashPlayerShotLocations(ctx context.Context, params LeagueDashPlayerShotLocationsParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching leaguedashplayershotlocations")

	reqParams := map[string]string{
		"DistanceRange": params.DistanceRange,
		"LastNGames": params.LastNGames,
		"MeasureType": params.MeasureTypeSimple,
		"Month": params.Month,
		"OpponentTeamID": params.OpponentTeamId,
		"PaceAdjust": params.PaceAdjust,
		"PerMode": params.PerModeDetailed,
		"Period": params.Period,
		"PlusMinus": params.PlusMinus,
		"Rank": params.Rank,
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
		"LeagueID": params.LeagueIdNullable,
		"Location": params.LocationNullable,
		"Outcome": params.OutcomeNullable,
		"PORound": params.PoRoundNullable,
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

	resp, err := c.httpClient.SendRequest(ctx, "leaguedashplayershotlocations", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch leaguedashplayershotlocations",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch leaguedashplayershotlocations: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from leaguedashplayershotlocations endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal leaguedashplayershotlocations response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched leaguedashplayershotlocations",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
