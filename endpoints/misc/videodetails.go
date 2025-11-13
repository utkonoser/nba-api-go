package misc

import (
	"context"
	"fmt"
	"log/slog"
)

// VideoDetailsParams holds parameters for the VideoDetails endpoint.
type VideoDetailsParams struct {
	TeamId string
	PlayerId string
	ContextMeasureDetailed string
	LastNGames string
	Month string
	OpponentTeamId string
	Period string
	Season string
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
	LeagueIdNullable string
	LocationNullable string
	OutcomeNullable string
	PointDiffNullable string
	PositionNullable string
	RangeTypeNullable string
	RookieYearNullable string
	SeasonSegmentNullable string
	StartPeriodNullable string
	StartRangeNullable string
	VsConferenceNullable string
	VsDivisionNullable string
}

// GetVideoDetails fetches data from the videodetails endpoint.
func (c *Client) GetVideoDetails(ctx context.Context, params VideoDetailsParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching videodetails")

	reqParams := map[string]string{
		"TeamID": params.TeamId,
		"PlayerID": params.PlayerId,
		"ContextMeasure": params.ContextMeasureDetailed,
		"LastNGames": params.LastNGames,
		"Month": params.Month,
		"OpponentTeamID": params.OpponentTeamId,
		"Period": params.Period,
		"Season": params.Season,
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
		"LeagueID": params.LeagueIdNullable,
		"Location": params.LocationNullable,
		"Outcome": params.OutcomeNullable,
		"PointDiff": params.PointDiffNullable,
		"Position": params.PositionNullable,
		"RangeType": params.RangeTypeNullable,
		"RookieYear": params.RookieYearNullable,
		"SeasonSegment": params.SeasonSegmentNullable,
		"StartPeriod": params.StartPeriodNullable,
		"StartRange": params.StartRangeNullable,
		"VsConference": params.VsConferenceNullable,
		"VsDivision": params.VsDivisionNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "videodetails", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch videodetails",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch videodetails: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from videodetails endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal videodetails response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched videodetails",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
