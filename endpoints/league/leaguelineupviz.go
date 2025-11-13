package league

import (
	"context"
	"fmt"
	"log/slog"
)

// LeagueLineupVizParams holds parameters for the LeagueLineupViz endpoint.
type LeagueLineupVizParams struct {
	MinutesMin string
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

// GetLeagueLineupViz fetches data from the leaguelineupviz endpoint.
func (c *Client) GetLeagueLineupViz(ctx context.Context, params LeagueLineupVizParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching leaguelineupviz")

	reqParams := map[string]string{
		"MinutesMin": params.MinutesMin,
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

	resp, err := c.httpClient.SendRequest(ctx, "leaguelineupviz", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch leaguelineupviz",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch leaguelineupviz: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from leaguelineupviz endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal leaguelineupviz response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched leaguelineupviz",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
