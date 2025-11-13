package team

import (
	"context"
	"fmt"
	"log/slog"
)

// TeamPlayerOnOffSummaryParams holds parameters for the TeamPlayerOnOffSummary endpoint.
type TeamPlayerOnOffSummaryParams struct {
	TeamId string
	LastNGames string
	MeasureTypeDetailedDefense string
	Month string
	OpponentTeamId string
	PaceAdjust string
	PerModeDetailed string
	Period string
	PlusMinus string
	Rank string
	Season string
	SeasonTypeAllStar string
	DateFromNullable string
	DateToNullable string
	GameSegmentNullable string
	LeagueIdNullable string
	LocationNullable string
	OutcomeNullable string
	SeasonSegmentNullable string
	VsConferenceNullable string
	VsDivisionNullable string
}

// GetTeamPlayerOnOffSummary fetches data from the teamplayeronoffsummary endpoint.
func (c *Client) GetTeamPlayerOnOffSummary(ctx context.Context, params TeamPlayerOnOffSummaryParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching teamplayeronoffsummary")

	reqParams := map[string]string{
		"TeamID": params.TeamId,
		"LastNGames": params.LastNGames,
		"MeasureType": params.MeasureTypeDetailedDefense,
		"Month": params.Month,
		"OpponentTeamID": params.OpponentTeamId,
		"PaceAdjust": params.PaceAdjust,
		"PerMode": params.PerModeDetailed,
		"Period": params.Period,
		"PlusMinus": params.PlusMinus,
		"Rank": params.Rank,
		"Season": params.Season,
		"SeasonType": params.SeasonTypeAllStar,
		"DateFrom": params.DateFromNullable,
		"DateTo": params.DateToNullable,
		"GameSegment": params.GameSegmentNullable,
		"LeagueID": params.LeagueIdNullable,
		"Location": params.LocationNullable,
		"Outcome": params.OutcomeNullable,
		"SeasonSegment": params.SeasonSegmentNullable,
		"VsConference": params.VsConferenceNullable,
		"VsDivision": params.VsDivisionNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "teamplayeronoffsummary", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch teamplayeronoffsummary",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch teamplayeronoffsummary: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from teamplayeronoffsummary endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal teamplayeronoffsummary response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched teamplayeronoffsummary",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
