package draft

import (
	"context"
	"fmt"
	"log/slog"
)

// DraftBoardParams holds parameters for the DraftBoard endpoint.
type DraftBoardParams struct {
	LeagueId string
	SeasonYear string
	CollegeNullable string
	OverallPickNullable string
	RoundNumNullable string
	RoundPickNullable string
	TeamIdNullable string
	TopxNullable string
}

// GetDraftBoard fetches data from the draftboard endpoint.
func (c *Client) GetDraftBoard(ctx context.Context, params DraftBoardParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching draftboard")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
		"Season": params.SeasonYear,
		"College": params.CollegeNullable,
		"OverallPick": params.OverallPickNullable,
		"RoundNum": params.RoundNumNullable,
		"RoundPick": params.RoundPickNullable,
		"TeamID": params.TeamIdNullable,
		"TopX": params.TopxNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "draftboard", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch draftboard",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch draftboard: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from draftboard endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal draftboard response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched draftboard",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
