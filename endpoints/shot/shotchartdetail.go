package shot

import (
	"context"
	"fmt"
	"log/slog"
)

// ShotChartDetailParams holds parameters for the ShotChartDetail endpoint.
type ShotChartDetailParams struct {
	TeamId string
	PlayerId string
	ContextMeasureSimple string
	LastNGames string
	LeagueId string
	Month string
	OpponentTeamId string
	Period string
	SeasonTypeAllStar string
	AheadBehindNullable string
	ClutchTimeNullable string
	ContextFilterNullable string
	DateFromNullable string
	DateToNullable string
	EndPeriodNullable string
	EndRangeNullable string
	GameIdNullable string
	GameSegmentNullable string
	LocationNullable string
	OutcomeNullable string
	PlayerPositionNullable string
	PointDiffNullable string
	PositionNullable string
	RangeTypeNullable string
	RookieYearNullable string
	SeasonNullable string
	SeasonSegmentNullable string
	StartPeriodNullable string
	StartRangeNullable string
	VsConferenceNullable string
	VsDivisionNullable string
}

// GetShotChartDetail fetches data from the shotchartdetail endpoint.
func (c *Client) GetShotChartDetail(ctx context.Context, params ShotChartDetailParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching shotchartdetail")

	reqParams := map[string]string{
		"TeamID": params.TeamId,
		"PlayerID": params.PlayerId,
		"ContextMeasure": params.ContextMeasureSimple,
		"LastNGames": params.LastNGames,
		"LeagueID": params.LeagueId,
		"Month": params.Month,
		"OpponentTeamID": params.OpponentTeamId,
		"Period": params.Period,
		"SeasonType": params.SeasonTypeAllStar,
		"AheadBehind": params.AheadBehindNullable,
		"ClutchTime": params.ClutchTimeNullable,
		"ContextFilter": params.ContextFilterNullable,
		"DateFrom": params.DateFromNullable,
		"DateTo": params.DateToNullable,
		"EndPeriod": params.EndPeriodNullable,
		"EndRange": params.EndRangeNullable,
		"GameID": params.GameIdNullable,
		"GameSegment": params.GameSegmentNullable,
		"Location": params.LocationNullable,
		"Outcome": params.OutcomeNullable,
		"PlayerPosition": params.PlayerPositionNullable,
		"PointDiff": params.PointDiffNullable,
		"Position": params.PositionNullable,
		"RangeType": params.RangeTypeNullable,
		"RookieYear": params.RookieYearNullable,
		"Season": params.SeasonNullable,
		"SeasonSegment": params.SeasonSegmentNullable,
		"StartPeriod": params.StartPeriodNullable,
		"StartRange": params.StartRangeNullable,
		"VsConference": params.VsConferenceNullable,
		"VsDivision": params.VsDivisionNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "shotchartdetail", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch shotchartdetail",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch shotchartdetail: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from shotchartdetail endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal shotchartdetail response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched shotchartdetail",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
