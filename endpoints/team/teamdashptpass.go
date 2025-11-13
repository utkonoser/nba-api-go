package team

import (
	"context"
	"fmt"
	"log/slog"
)

// TeamDashPtPassParams holds parameters for the TeamDashPtPass endpoint.
type TeamDashPtPassParams struct {
	TeamId string
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

// GetTeamDashPtPass fetches data from the teamdashptpass endpoint.
func (c *Client) GetTeamDashPtPass(ctx context.Context, params TeamDashPtPassParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching teamdashptpass")

	reqParams := map[string]string{
		"TeamID": params.TeamId,
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

	resp, err := c.httpClient.SendRequest(ctx, "teamdashptpass", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch teamdashptpass",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch teamdashptpass: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from teamdashptpass endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal teamdashptpass response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched teamdashptpass",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
