package team

import (
	"context"
	"fmt"
	"log/slog"
)

// TeamDashLineupsParams holds parameters for the TeamDashLineups endpoint.
type TeamDashLineupsParams struct {
	TeamId string
	GroupQuantity string
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
	GameIdNullable string
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

// GetTeamDashLineups fetches data from the teamdashlineups endpoint.
func (c *Client) GetTeamDashLineups(ctx context.Context, params TeamDashLineupsParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching teamdashlineups")

	reqParams := map[string]string{
		"TeamID": params.TeamId,
		"GroupQuantity": params.GroupQuantity,
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
		"GameID": params.GameIdNullable,
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

	resp, err := c.httpClient.SendRequest(ctx, "teamdashlineups", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch teamdashlineups",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch teamdashlineups: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from teamdashlineups endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal teamdashlineups response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched teamdashlineups",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
