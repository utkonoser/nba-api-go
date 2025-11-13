package misc

import (
	"context"
	"fmt"
	"log/slog"
)

// FantasyWidgetParams holds parameters for the FantasyWidget endpoint.
type FantasyWidgetParams struct {
	ActivePlayers string
	LastNGames string
	LeagueId string
	Season string
	SeasonTypeAllStar string
	TodaysOpponent string
	TodaysPlayers string
	DateFromNullable string
	DateToNullable string
	LocationNullable string
	MonthNullable string
	OpponentTeamIdNullable string
	PoRoundNullable string
	PlayerIdNullable string
	PositionNullable string
	SeasonSegmentNullable string
	TeamIdNullable string
	VsConferenceNullable string
	VsDivisionNullable string
}

// GetFantasyWidget fetches data from the fantasywidget endpoint.
func (c *Client) GetFantasyWidget(ctx context.Context, params FantasyWidgetParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching fantasywidget")

	reqParams := map[string]string{
		"ActivePlayers": params.ActivePlayers,
		"LastNGames": params.LastNGames,
		"LeagueID": params.LeagueId,
		"Season": params.Season,
		"SeasonType": params.SeasonTypeAllStar,
		"TodaysOpponent": params.TodaysOpponent,
		"TodaysPlayers": params.TodaysPlayers,
		"DateFrom": params.DateFromNullable,
		"DateTo": params.DateToNullable,
		"Location": params.LocationNullable,
		"Month": params.MonthNullable,
		"OpponentTeamID": params.OpponentTeamIdNullable,
		"PORound": params.PoRoundNullable,
		"PlayerID": params.PlayerIdNullable,
		"Position": params.PositionNullable,
		"SeasonSegment": params.SeasonSegmentNullable,
		"TeamID": params.TeamIdNullable,
		"VsConference": params.VsConferenceNullable,
		"VsDivision": params.VsDivisionNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "fantasywidget", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch fantasywidget",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch fantasywidget: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from fantasywidget endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal fantasywidget response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched fantasywidget",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
