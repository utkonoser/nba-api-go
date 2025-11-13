package draft

import (
	"context"
	"fmt"
	"log/slog"
)

// DraftHistoryParams holds parameters for the DraftHistory endpoint.
type DraftHistoryParams struct {
	LeagueId string
	CollegeNullable string
	OverallPickNullable string
	RoundNumNullable string
	RoundPickNullable string
	SeasonYearNullable string
	TeamIdNullable string
	TopxNullable string
}

// GetDraftHistory fetches data from the drafthistory endpoint.
func (c *Client) GetDraftHistory(ctx context.Context, params DraftHistoryParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching drafthistory")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
		"College": params.CollegeNullable,
		"OverallPick": params.OverallPickNullable,
		"RoundNum": params.RoundNumNullable,
		"RoundPick": params.RoundPickNullable,
		"Season": params.SeasonYearNullable,
		"TeamID": params.TeamIdNullable,
		"TopX": params.TopxNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "drafthistory", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch drafthistory",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch drafthistory: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from drafthistory endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal drafthistory response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched drafthistory",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
