package league

import (
	"context"
	"fmt"
	"log/slog"
)

// LeagueDashTeamClutchParams holds parameters for the LeagueDashTeamClutch endpoint.
type LeagueDashTeamClutchParams struct {
	AheadBehind string
	ClutchTime string
	LastNGames string
	MeasureTypeDetailedDefense string
	Month string
	OpponentTeamId string
	PaceAdjust string
	PerModeDetailed string
	Period string
	PlusMinus string
	PointDiff string
	Rank string
	Season string
	SeasonTypeAllStar string
	ConferenceNullable string
	DateFromNullable string
	DateToNullable string
	DivisionSimpleNullable string
	GameScopeSimpleNullable string
	GameSegmentNullable string
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
}

// GetLeagueDashTeamClutch fetches data from the leaguedashteamclutch endpoint.
func (c *Client) GetLeagueDashTeamClutch(ctx context.Context, params LeagueDashTeamClutchParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching leaguedashteamclutch")

	reqParams := map[string]string{
		"AheadBehind": params.AheadBehind,
		"ClutchTime": params.ClutchTime,
		"LastNGames": params.LastNGames,
		"MeasureType": params.MeasureTypeDetailedDefense,
		"Month": params.Month,
		"OpponentTeamID": params.OpponentTeamId,
		"PaceAdjust": params.PaceAdjust,
		"PerMode": params.PerModeDetailed,
		"Period": params.Period,
		"PlusMinus": params.PlusMinus,
		"PointDiff": params.PointDiff,
		"Rank": params.Rank,
		"Season": params.Season,
		"SeasonType": params.SeasonTypeAllStar,
		"Conference": params.ConferenceNullable,
		"DateFrom": params.DateFromNullable,
		"DateTo": params.DateToNullable,
		"Division": params.DivisionSimpleNullable,
		"GameScope": params.GameScopeSimpleNullable,
		"GameSegment": params.GameSegmentNullable,
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
	}

	resp, err := c.httpClient.SendRequest(ctx, "leaguedashteamclutch", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch leaguedashteamclutch",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch leaguedashteamclutch: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from leaguedashteamclutch endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal leaguedashteamclutch response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched leaguedashteamclutch",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
