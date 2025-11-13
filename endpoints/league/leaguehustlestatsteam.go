package league

import (
	"context"
	"fmt"
	"log/slog"
)

// LeagueHustleStatsTeamParams holds parameters for the LeagueHustleStatsTeam endpoint.
type LeagueHustleStatsTeamParams struct {
	PerModeTime string
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
	HeightNullable string
	LeagueIdNullable string
	LocationNullable string
	MonthNullable string
	OpponentTeamIdNullable string
	OutcomeNullable string
	PoRoundNullable string
	PlayerExperienceNullable string
	PlayerPositionNullable string
	SeasonSegmentNullable string
	TeamIdNullable string
	VsConferenceNullable string
	VsDivisionNullable string
	WeightNullable string
}

// GetLeagueHustleStatsTeam fetches data from the leaguehustlestatsteam endpoint.
func (c *Client) GetLeagueHustleStatsTeam(ctx context.Context, params LeagueHustleStatsTeamParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching leaguehustlestatsteam")

	reqParams := map[string]string{
		"PerMode": params.PerModeTime,
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
		"Height": params.HeightNullable,
		"LeagueID": params.LeagueIdNullable,
		"Location": params.LocationNullable,
		"Month": params.MonthNullable,
		"OpponentTeamID": params.OpponentTeamIdNullable,
		"Outcome": params.OutcomeNullable,
		"PORound": params.PoRoundNullable,
		"PlayerExperience": params.PlayerExperienceNullable,
		"PlayerPosition": params.PlayerPositionNullable,
		"SeasonSegment": params.SeasonSegmentNullable,
		"TeamID": params.TeamIdNullable,
		"VsConference": params.VsConferenceNullable,
		"VsDivision": params.VsDivisionNullable,
		"Weight": params.WeightNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "leaguehustlestatsteam", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch leaguehustlestatsteam",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch leaguehustlestatsteam: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from leaguehustlestatsteam endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal leaguehustlestatsteam response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched leaguehustlestatsteam",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
