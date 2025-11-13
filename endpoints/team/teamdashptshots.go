package team

import (
	"context"
	"fmt"
	"log/slog"
)

// TeamDashPtShotsParams holds parameters for the TeamDashPtShots endpoint.
type TeamDashPtShotsParams struct {
	TeamId string
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

// GetTeamDashPtShots fetches data from the teamdashptshots endpoint.
func (c *Client) GetTeamDashPtShots(ctx context.Context, params TeamDashPtShotsParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching teamdashptshots")

	reqParams := map[string]string{
		"TeamID": params.TeamId,
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

	resp, err := c.httpClient.SendRequest(ctx, "teamdashptshots", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch teamdashptshots",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch teamdashptshots: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from teamdashptshots endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal teamdashptshots response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched teamdashptshots",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
