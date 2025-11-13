package player

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayerGameLogsParams holds parameters for the PlayerGameLogs endpoint.
type PlayerGameLogsParams struct {
	DateFromNullable string
	DateToNullable string
	GameSegmentNullable string
	LastNGamesNullable string
	LeagueIdNullable string
	LocationNullable string
	MeasureTypePlayerGameLogsNullable string
	MonthNullable string
	OppTeamIdNullable string
	OutcomeNullable string
	PoRoundNullable string
	PerModeSimpleNullable string
	PeriodNullable string
	PlayerIdNullable string
	SeasonNullable string
	SeasonSegmentNullable string
	SeasonTypeNullable string
	ShotClockRangeNullable string
	TeamIdNullable string
}

// GetPlayerGameLogs fetches data from the playergamelogs endpoint.
func (c *Client) GetPlayerGameLogs(ctx context.Context, params PlayerGameLogsParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playergamelogs")

	reqParams := map[string]string{}
	
	// Add nullable parameters only if they are not empty
	if params.DateFromNullable != "" {
		reqParams["DateFrom"] = params.DateFromNullable
	}
	if params.DateToNullable != "" {
		reqParams["DateTo"] = params.DateToNullable
	}
	if params.GameSegmentNullable != "" {
		reqParams["GameSegment"] = params.GameSegmentNullable
	}
	if params.LastNGamesNullable != "" {
		reqParams["LastNGames"] = params.LastNGamesNullable
	}
	if params.LeagueIdNullable != "" {
		reqParams["LeagueID"] = params.LeagueIdNullable
	}
	if params.LocationNullable != "" {
		reqParams["Location"] = params.LocationNullable
	}
	if params.MeasureTypePlayerGameLogsNullable != "" {
		reqParams["MeasureType"] = params.MeasureTypePlayerGameLogsNullable
	}
	if params.MonthNullable != "" {
		reqParams["Month"] = params.MonthNullable
	}
	if params.OppTeamIdNullable != "" {
		reqParams["OppTeamID"] = params.OppTeamIdNullable
	}
	if params.OutcomeNullable != "" {
		reqParams["Outcome"] = params.OutcomeNullable
	}
	if params.PoRoundNullable != "" {
		reqParams["PORound"] = params.PoRoundNullable
	}
	if params.PerModeSimpleNullable != "" {
		reqParams["PerMode"] = params.PerModeSimpleNullable
	}
	if params.PeriodNullable != "" {
		reqParams["Period"] = params.PeriodNullable
	}
	if params.PlayerIdNullable != "" {
		reqParams["PlayerID"] = params.PlayerIdNullable
	}
	if params.SeasonNullable != "" {
		reqParams["Season"] = params.SeasonNullable
	}
	if params.SeasonSegmentNullable != "" {
		reqParams["SeasonSegment"] = params.SeasonSegmentNullable
	}
	if params.SeasonTypeNullable != "" {
		reqParams["SeasonType"] = params.SeasonTypeNullable
	}
	if params.ShotClockRangeNullable != "" {
		reqParams["ShotClockRange"] = params.ShotClockRangeNullable
	}
	if params.TeamIdNullable != "" {
		reqParams["TeamID"] = params.TeamIdNullable
	}

	resp, err := c.httpClient.SendRequest(ctx, "playergamelogs", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch playergamelogs",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch playergamelogs: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from playergamelogs endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal playergamelogs response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched playergamelogs",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}

