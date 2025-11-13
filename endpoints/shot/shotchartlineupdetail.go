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

	// Set defaults for required parameters
	contextMeasure := params.ContextMeasureDetailed
	if contextMeasure == "" {
		contextMeasure = "FGM"
	}
	groupId := params.GroupId
	if groupId == "" {
		groupId = "0"
	}
	leagueId := params.LeagueId
	if leagueId == "" {
		leagueId = "00"
	}
	period := params.Period
	if period == "" {
		period = "0"
	}
	season := params.Season
	if season == "" {
		season = "2023-24"
	}
	seasonType := params.SeasonTypeAllStar
	if seasonType == "" {
		seasonType = "Regular Season"
	}

	reqParams := map[string]string{
		"ContextMeasure": contextMeasure,
		"GROUP_ID": groupId,
		"LeagueID": leagueId,
		"Period": period,
		"Season": season,
		"SeasonType": seasonType,
	}
	
	// Add nullable parameters only if they are not empty
	if params.ContextFilterNullable != "" {
		reqParams["ContextFilter"] = params.ContextFilterNullable
	}
	if params.DateFromNullable != "" {
		reqParams["DateFrom"] = params.DateFromNullable
	}
	if params.DateToNullable != "" {
		reqParams["DateTo"] = params.DateToNullable
	}
	if params.GameIdNullable != "" {
		reqParams["GameID"] = params.GameIdNullable
	}
	if params.GameSegmentNullable != "" {
		reqParams["GameSegment"] = params.GameSegmentNullable
	}
	if params.LastNGamesNullable != "" {
		reqParams["LastNGames"] = params.LastNGamesNullable
	}
	if params.LocationNullable != "" {
		reqParams["Location"] = params.LocationNullable
	}
	if params.MonthNullable != "" {
		reqParams["Month"] = params.MonthNullable
	}
	if params.OpponentTeamIdNullable != "" {
		reqParams["OpponentTeamID"] = params.OpponentTeamIdNullable
	}
	if params.OutcomeNullable != "" {
		reqParams["Outcome"] = params.OutcomeNullable
	}
	if params.SeasonSegmentNullable != "" {
		reqParams["SeasonSegment"] = params.SeasonSegmentNullable
	}
	if params.TeamIdNullable != "" {
		reqParams["TeamID"] = params.TeamIdNullable
	}
	if params.VsConferenceNullable != "" {
		reqParams["VsConference"] = params.VsConferenceNullable
	}
	if params.VsDivisionNullable != "" {
		reqParams["VsDivision"] = params.VsDivisionNullable
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
