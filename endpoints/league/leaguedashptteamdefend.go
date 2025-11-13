package league

import (
	"context"
	"fmt"
	"log/slog"
)

// LeagueDashPtTeamDefendParams holds parameters for the LeagueDashPtTeamDefend endpoint.
type LeagueDashPtTeamDefendParams struct {
	DefenseCategory string
	LeagueId string
	PerModeSimple string
	Season string
	SeasonTypeAllStar string
	ConferenceNullable string
	DateFromNullable string
	DateToNullable string
	DivisionNullable string
	GameSegmentNullable string
	LastNGamesNullable string
	LocationNullable string
	MonthNullable string
	OpponentTeamIdNullable string
	OutcomeNullable string
	PoRoundNullable string
	PeriodNullable string
	SeasonSegmentNullable string
	TeamIdNullable string
	VsConferenceNullable string
	VsDivisionNullable string
}

// GetLeagueDashPtTeamDefend fetches data from the leaguedashptteamdefend endpoint.
func (c *Client) GetLeagueDashPtTeamDefend(ctx context.Context, params LeagueDashPtTeamDefendParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching leaguedashptteamdefend")

	reqParams := map[string]string{
		"DefenseCategory": params.DefenseCategory,
		"LeagueID": params.LeagueId,
		"PerMode": params.PerModeSimple,
		"Season": params.Season,
		"SeasonType": params.SeasonTypeAllStar,
		"Conference": params.ConferenceNullable,
		"DateFrom": params.DateFromNullable,
		"DateTo": params.DateToNullable,
		"Division": params.DivisionNullable,
		"GameSegment": params.GameSegmentNullable,
		"LastNGames": params.LastNGamesNullable,
		"Location": params.LocationNullable,
		"Month": params.MonthNullable,
		"OpponentTeamID": params.OpponentTeamIdNullable,
		"Outcome": params.OutcomeNullable,
		"PORound": params.PoRoundNullable,
		"Period": params.PeriodNullable,
		"SeasonSegment": params.SeasonSegmentNullable,
		"TeamID": params.TeamIdNullable,
		"VsConference": params.VsConferenceNullable,
		"VsDivision": params.VsDivisionNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "leaguedashptteamdefend", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch leaguedashptteamdefend",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch leaguedashptteamdefend: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from leaguedashptteamdefend endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal leaguedashptteamdefend response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched leaguedashptteamdefend",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
