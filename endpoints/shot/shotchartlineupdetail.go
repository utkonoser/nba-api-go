package shot

import (
	"context"
	"fmt"
	"log/slog"
)

// ShotChartLineupDetailParams holds parameters for the ShotChartLineupDetail endpoint.
type ShotChartLineupDetailParams struct {
	ContextMeasureDetailed string
	GroupId string
	LeagueId string
	Period string
	Season string
	SeasonTypeAllStar string
	ContextFilterNullable string
	DateFromNullable string
	DateToNullable string
	GameIdNullable string
	GameSegmentNullable string
	LastNGamesNullable string
	LocationNullable string
	MonthNullable string
	OpponentTeamIdNullable string
	OutcomeNullable string
	SeasonSegmentNullable string
	TeamIdNullable string
	VsConferenceNullable string
	VsDivisionNullable string
}

// GetShotChartLineupDetail fetches data from the shotchartlineupdetail endpoint.
func (c *Client) GetShotChartLineupDetail(ctx context.Context, params ShotChartLineupDetailParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching shotchartlineupdetail")

	reqParams := map[string]string{
		"ContextMeasure": params.ContextMeasureDetailed,
		"GROUP_ID": params.GroupId,
		"LeagueID": params.LeagueId,
		"Period": params.Period,
		"Season": params.Season,
		"SeasonType": params.SeasonTypeAllStar,
		"ContextFilter": params.ContextFilterNullable,
		"DateFrom": params.DateFromNullable,
		"DateTo": params.DateToNullable,
		"GameID": params.GameIdNullable,
		"GameSegment": params.GameSegmentNullable,
		"LastNGames": params.LastNGamesNullable,
		"Location": params.LocationNullable,
		"Month": params.MonthNullable,
		"OpponentTeamID": params.OpponentTeamIdNullable,
		"Outcome": params.OutcomeNullable,
		"SeasonSegment": params.SeasonSegmentNullable,
		"TeamID": params.TeamIdNullable,
		"VsConference": params.VsConferenceNullable,
		"VsDivision": params.VsDivisionNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "shotchartlineupdetail", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch shotchartlineupdetail",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch shotchartlineupdetail: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from shotchartlineupdetail endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal shotchartlineupdetail response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched shotchartlineupdetail",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
