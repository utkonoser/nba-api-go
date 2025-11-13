package player

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayerDashPtPassParams holds parameters for the PlayerDashPtPass endpoint.
type PlayerDashPtPassParams struct {
	TeamId string
	PlayerId string
	LastNGames string
	LeagueId string
	Month string
	OpponentTeamId string
	PerModeSimple string
	Season string
	SeasonTypeAllStar string
	DateFromNullable string
	DateToNullable string
	LocationNullable string
	OutcomeNullable string
	SeasonSegmentNullable string
	VsConferenceNullable string
	VsDivisionNullable string
}

// GetPlayerDashPtPass fetches data from the playerdashptpass endpoint.
func (c *Client) GetPlayerDashPtPass(ctx context.Context, params PlayerDashPtPassParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playerdashptpass")

	reqParams := map[string]string{
		"TeamID": params.TeamId,
		"PlayerID": params.PlayerId,
		"LastNGames": params.LastNGames,
		"LeagueID": params.LeagueId,
		"Month": params.Month,
		"OpponentTeamID": params.OpponentTeamId,
		"PerMode": params.PerModeSimple,
		"Season": params.Season,
		"SeasonType": params.SeasonTypeAllStar,
		"DateFrom": params.DateFromNullable,
		"DateTo": params.DateToNullable,
		"Location": params.LocationNullable,
		"Outcome": params.OutcomeNullable,
		"SeasonSegment": params.SeasonSegmentNullable,
		"VsConference": params.VsConferenceNullable,
		"VsDivision": params.VsDivisionNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "playerdashptpass", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch playerdashptpass",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch playerdashptpass: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from playerdashptpass endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal playerdashptpass response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched playerdashptpass",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
