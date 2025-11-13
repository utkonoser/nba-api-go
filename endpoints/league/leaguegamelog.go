package league

import (
	"context"
	"fmt"
	"log/slog"
)

// LeagueGameLogParams holds parameters for the LeagueGameLog endpoint.
type LeagueGameLogParams struct {
	Counter string
	Direction string
	LeagueId string
	PlayerOrTeamAbbreviation string
	Season string
	SeasonTypeAllStar string
	Sorter string
	DateFromNullable string
	DateToNullable string
}

// GetLeagueGameLog fetches data from the leaguegamelog endpoint.
func (c *Client) GetLeagueGameLog(ctx context.Context, params LeagueGameLogParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching leaguegamelog")

	reqParams := map[string]string{
		"Counter": params.Counter,
		"Direction": params.Direction,
		"LeagueID": params.LeagueId,
		"PlayerOrTeam": params.PlayerOrTeamAbbreviation,
		"Season": params.Season,
		"SeasonType": params.SeasonTypeAllStar,
		"Sorter": params.Sorter,
		"DateFrom": params.DateFromNullable,
		"DateTo": params.DateToNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "leaguegamelog", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch leaguegamelog",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch leaguegamelog: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from leaguegamelog endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal leaguegamelog response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched leaguegamelog",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
