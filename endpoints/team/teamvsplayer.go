package team

import (
	"context"
	"fmt"
	"log/slog"
)

// TeamVsPlayerParams holds parameters for the TeamVsPlayer endpoint.
type TeamVsPlayerParams struct {
	VsPlayerId string
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
	SeasonTypePlayoffs string
	DateFromNullable string
	DateToNullable string
	GameSegmentNullable string
	LeagueIdNullable string
	LocationNullable string
	OutcomeNullable string
	PlayerIdNullable string
	SeasonSegmentNullable string
	VsConferenceNullable string
	VsDivisionNullable string
}

// GetTeamVsPlayer fetches data from the teamvsplayer endpoint.
func (c *Client) GetTeamVsPlayer(ctx context.Context, params TeamVsPlayerParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching teamvsplayer")

	reqParams := map[string]string{
		"VsPlayerID": params.VsPlayerId,
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
		"SeasonType": params.SeasonTypePlayoffs,
		"DateFrom": params.DateFromNullable,
		"DateTo": params.DateToNullable,
		"GameSegment": params.GameSegmentNullable,
		"LeagueID": params.LeagueIdNullable,
		"Location": params.LocationNullable,
		"Outcome": params.OutcomeNullable,
		"PlayerID": params.PlayerIdNullable,
		"SeasonSegment": params.SeasonSegmentNullable,
		"VsConference": params.VsConferenceNullable,
		"VsDivision": params.VsDivisionNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "teamvsplayer", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch teamvsplayer",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch teamvsplayer: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from teamvsplayer endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal teamvsplayer response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched teamvsplayer",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
