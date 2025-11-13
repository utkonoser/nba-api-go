package player

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayerDashboardByTeamPerformanceParams holds parameters for the PlayerDashboardByTeamPerformance endpoint.
type PlayerDashboardByTeamPerformanceParams struct {
	PlayerId string
	LastNGames string
	MeasureTypeDetailed string
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
	PoRoundNullable string
	SeasonSegmentNullable string
	ShotClockRangeNullable string
	VsConferenceNullable string
	VsDivisionNullable string
}

// GetPlayerDashboardByTeamPerformance fetches data from the playerdashboardbyteamperformance endpoint.
func (c *Client) GetPlayerDashboardByTeamPerformance(ctx context.Context, params PlayerDashboardByTeamPerformanceParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playerdashboardbyteamperformance")

	reqParams := map[string]string{
		"PlayerID": params.PlayerId,
		"LastNGames": params.LastNGames,
		"MeasureType": params.MeasureTypeDetailed,
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
		"PORound": params.PoRoundNullable,
		"SeasonSegment": params.SeasonSegmentNullable,
		"ShotClockRange": params.ShotClockRangeNullable,
		"VsConference": params.VsConferenceNullable,
		"VsDivision": params.VsDivisionNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "playerdashboardbyteamperformance", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch playerdashboardbyteamperformance",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch playerdashboardbyteamperformance: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from playerdashboardbyteamperformance endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal playerdashboardbyteamperformance response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched playerdashboardbyteamperformance",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
