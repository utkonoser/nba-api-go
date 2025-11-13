package player

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayerGameLogsParams holds parameters for the PlayerGameLogs endpoint.
type PlayerGameLogsParams struct {
	DateFromNullable string
	DateToNullable string
	GameSegmentNullable string
	LastNGamesNullable string
	LeagueIdNullable string
	LocationNullable string
	MeasureTypePlayerGameLogsNullable string
	MonthNullable string
	OppTeamIdNullable string
	OutcomeNullable string
	PoRoundNullable string
	PerModeSimpleNullable string
	PeriodNullable string
	PlayerIdNullable string
	SeasonNullable string
	SeasonSegmentNullable string
	SeasonTypeNullable string
	ShotClockRangeNullable string
	TeamIdNullable string
	VsConferenceNullable string
	VsDivisionNullable string
}

// GetPlayerGameLogs fetches data from the playergamelogs endpoint.
func (c *Client) GetPlayerGameLogs(ctx context.Context, params PlayerGameLogsParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playergamelogs")

	reqParams := map[string]string{
		"DateFrom": params.DateFromNullable,
		"DateTo": params.DateToNullable,
		"GameSegment": params.GameSegmentNullable,
		"LastNGames": params.LastNGamesNullable,
		"LeagueID": params.LeagueIdNullable,
		"Location": params.LocationNullable,
		"MeasureType": params.MeasureTypePlayerGameLogsNullable,
		"Month": params.MonthNullable,
		"OpponentTeamID": params.OppTeamIdNullable,
		"Outcome": params.OutcomeNullable,
		"PORound": params.PoRoundNullable,
		"PerMode": params.PerModeSimpleNullable,
		"Period": params.PeriodNullable,
		"PlayerID": params.PlayerIdNullable,
		"Season": params.SeasonNullable,
		"SeasonSegment": params.SeasonSegmentNullable,
		"SeasonType": params.SeasonTypeNullable,
		"ShotClockRange": params.ShotClockRangeNullable,
		"TeamID": params.TeamIdNullable,
		"VsConference": params.VsConferenceNullable,
		"VsDivision": params.VsDivisionNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "playergamelogs", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch playergamelogs",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch playergamelogs: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from playergamelogs endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal playergamelogs response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched playergamelogs",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
