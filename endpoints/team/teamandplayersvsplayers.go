package team

import (
	"context"
	"fmt"
	"log/slog"
)

// TeamAndPlayersVsPlayersParams holds parameters for the TeamAndPlayersVsPlayers endpoint.
type TeamAndPlayersVsPlayersParams struct {
	VsTeamId string
	VsPlayerId5 string
	VsPlayerId4 string
	VsPlayerId3 string
	VsPlayerId2 string
	VsPlayerId1 string
	TeamId string
	PlayerId5 string
	PlayerId4 string
	PlayerId3 string
	PlayerId2 string
	PlayerId1 string
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
	ConferenceNullable string
	DateFromNullable string
	DateToNullable string
	DivisionSimpleNullable string
	GameSegmentNullable string
	LeagueIdNullable string
	LocationNullable string
	OutcomeNullable string
	SeasonSegmentNullable string
	ShotClockRangeNullable string
	VsConferenceNullable string
	VsDivisionNullable string
}

// GetTeamAndPlayersVsPlayers fetches data from the teamandplayersvsplayers endpoint.
func (c *Client) GetTeamAndPlayersVsPlayers(ctx context.Context, params TeamAndPlayersVsPlayersParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching teamandplayersvsplayers")

	reqParams := map[string]string{
		"VsTeamID": params.VsTeamId,
		"VsPlayerID5": params.VsPlayerId5,
		"VsPlayerID4": params.VsPlayerId4,
		"VsPlayerID3": params.VsPlayerId3,
		"VsPlayerID2": params.VsPlayerId2,
		"VsPlayerID1": params.VsPlayerId1,
		"TeamID": params.TeamId,
		"PlayerID5": params.PlayerId5,
		"PlayerID4": params.PlayerId4,
		"PlayerID3": params.PlayerId3,
		"PlayerID2": params.PlayerId2,
		"PlayerID1": params.PlayerId1,
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
		"Conference": params.ConferenceNullable,
		"DateFrom": params.DateFromNullable,
		"DateTo": params.DateToNullable,
		"Division": params.DivisionSimpleNullable,
		"GameSegment": params.GameSegmentNullable,
		"LeagueID": params.LeagueIdNullable,
		"Location": params.LocationNullable,
		"Outcome": params.OutcomeNullable,
		"SeasonSegment": params.SeasonSegmentNullable,
		"ShotClockRange": params.ShotClockRangeNullable,
		"VsConference": params.VsConferenceNullable,
		"VsDivision": params.VsDivisionNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "teamandplayersvsplayers", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch teamandplayersvsplayers",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch teamandplayersvsplayers: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from teamandplayersvsplayers endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal teamandplayersvsplayers response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched teamandplayersvsplayers",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
