package team

import (
	"context"
	"fmt"
	"log/slog"
)

// TeamDashPtRebParams holds parameters for the TeamDashPtReb endpoint.
type TeamDashPtRebParams struct {
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

// GetTeamDashPtReb fetches data from the teamdashptreb endpoint.
func (c *Client) GetTeamDashPtReb(ctx context.Context, params TeamDashPtRebParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching teamdashptreb")

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

	resp, err := c.httpClient.SendRequest(ctx, "teamdashptreb", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch teamdashptreb",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch teamdashptreb: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from teamdashptreb endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal teamdashptreb response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched teamdashptreb",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
