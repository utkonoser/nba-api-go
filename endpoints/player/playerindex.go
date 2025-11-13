package player

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayerIndexParams holds parameters for the PlayerIndex endpoint.
type PlayerIndexParams struct {
	ActiveNullable string
	AllstarNullable string
	CollegeNullable string
	CountryNullable string
	DraftPickNullable string
	DraftYearNullable string
	HeightNullable string
	PlayerPositionAbbreviationNullable string
	HistoricalNullable string
	LeagueId string
	Season string
	TeamIdNullable string
	WeightNullable string
}

// GetPlayerIndex fetches data from the playerindex endpoint.
func (c *Client) GetPlayerIndex(ctx context.Context, params PlayerIndexParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playerindex")

	reqParams := map[string]string{
		"Active": params.ActiveNullable,
		"AllStar": params.AllstarNullable,
		"College": params.CollegeNullable,
		"Country": params.CountryNullable,
		"DraftPick": params.DraftPickNullable,
		"DraftYear": params.DraftYearNullable,
		"Height": params.HeightNullable,
		"PlayerPosition": params.PlayerPositionAbbreviationNullable,
		"Historical": params.HistoricalNullable,
		"LeagueID": params.LeagueId,
		"Season": params.Season,
		"TeamID": params.TeamIdNullable,
		"Weight": params.WeightNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "playerindex", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch playerindex",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch playerindex: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from playerindex endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal playerindex response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched playerindex",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
