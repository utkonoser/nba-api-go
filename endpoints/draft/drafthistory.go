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
	}
	
	// Add optional parameters only if they are not empty
	if params.CollegeNullable != "" {
		reqParams["College"] = params.CollegeNullable
	}
	if params.OverallPickNullable != "" {
		reqParams["OverallPick"] = params.OverallPickNullable
	}
	if params.RoundNumNullable != "" {
		reqParams["RoundNum"] = params.RoundNumNullable
	}
	if params.RoundPickNullable != "" {
		reqParams["RoundPick"] = params.RoundPickNullable
	}
	if params.SeasonYearNullable != "" {
		reqParams["Season"] = params.SeasonYearNullable
	}
	if params.TeamIdNullable != "" {
		reqParams["TeamID"] = params.TeamIdNullable
	}
	if params.TopxNullable != "" {
		reqParams["TopX"] = params.TopxNullable
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
