package team

import (
	"context"
	"fmt"
	"log/slog"
)

// TeamPlayerOnOffDetailsParams holds parameters for the TeamPlayerOnOffDetails endpoint.
type TeamPlayerOnOffDetailsParams struct {
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

// GetTeamPlayerOnOffDetails fetches data from the teamplayeronoffdetails endpoint.
func (c *Client) GetTeamPlayerOnOffDetails(ctx context.Context, params TeamPlayerOnOffDetailsParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching teamplayeronoffdetails")

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

	resp, err := c.httpClient.SendRequest(ctx, "teamplayeronoffdetails", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch teamplayeronoffdetails",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch teamplayeronoffdetails: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from teamplayeronoffdetails endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal teamplayeronoffdetails response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched teamplayeronoffdetails",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
