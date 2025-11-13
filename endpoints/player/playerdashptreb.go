package player

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayerDashPtRebParams holds parameters for the PlayerDashPtReb endpoint.
type PlayerDashPtRebParams struct {
	TeamId string
	PlayerId string
	LastNGames string
	LeagueId string
	Month string
	OpponentTeamId string
	PerModeSimple string
	Period string
	Season string
	SeasonTypeAllStar string
	DateFromNullable string
	DateToNullable string
	GameSegmentNullable string
	LocationNullable string
	OutcomeNullable string
	SeasonSegmentNullable string
	VsConferenceNullable string
	VsDivisionNullable string
}

// GetPlayerDashPtReb fetches data from the playerdashptreb endpoint.
func (c *Client) GetPlayerDashPtReb(ctx context.Context, params PlayerDashPtRebParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playerdashptreb")

	reqParams := map[string]string{
		"TeamID": params.TeamId,
		"PlayerID": params.PlayerId,
		"LastNGames": params.LastNGames,
		"LeagueID": params.LeagueId,
		"Month": params.Month,
		"OpponentTeamID": params.OpponentTeamId,
		"PerMode": params.PerModeSimple,
		"Period": params.Period,
		"Season": params.Season,
		"SeasonType": params.SeasonTypeAllStar,
		"DateFrom": params.DateFromNullable,
		"DateTo": params.DateToNullable,
		"GameSegment": params.GameSegmentNullable,
		"Location": params.LocationNullable,
		"Outcome": params.OutcomeNullable,
		"SeasonSegment": params.SeasonSegmentNullable,
		"VsConference": params.VsConferenceNullable,
		"VsDivision": params.VsDivisionNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "playerdashptreb", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch playerdashptreb",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch playerdashptreb: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from playerdashptreb endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal playerdashptreb response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched playerdashptreb",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
