package league

import (
	"context"
	"fmt"
	"log/slog"
)

// LeagueDashLineupsParams holds parameters for the LeagueDashLineups endpoint.
type LeagueDashLineupsParams struct {
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
	ConferenceNullable string
	DateFromNullable string
	DateToNullable string
	DivisionSimpleNullable string
	GameSegmentNullable string
	LeagueIdNullable string
	LocationNullable string
	OutcomeNullable string
	PoRoundNullable string
	SeasonSegmentNullable string
	ShotClockRangeNullable string
	TeamIdNullable string
	VsConferenceNullable string
	VsDivisionNullable string
}

// GetLeagueDashLineups fetches data from the leaguedashlineups endpoint.
func (c *Client) GetLeagueDashLineups(ctx context.Context, params LeagueDashLineupsParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching leaguedashlineups")

	reqParams := map[string]string{
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
		"Conference": params.ConferenceNullable,
		"DateFrom": params.DateFromNullable,
		"DateTo": params.DateToNullable,
		"Division": params.DivisionSimpleNullable,
		"GameSegment": params.GameSegmentNullable,
		"LeagueID": params.LeagueIdNullable,
		"Location": params.LocationNullable,
		"Outcome": params.OutcomeNullable,
		"PORound": params.PoRoundNullable,
		"SeasonSegment": params.SeasonSegmentNullable,
		"ShotClockRange": params.ShotClockRangeNullable,
		"TeamID": params.TeamIdNullable,
		"VsConference": params.VsConferenceNullable,
		"VsDivision": params.VsDivisionNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "leaguedashlineups", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch leaguedashlineups",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch leaguedashlineups: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from leaguedashlineups endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal leaguedashlineups response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched leaguedashlineups",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
