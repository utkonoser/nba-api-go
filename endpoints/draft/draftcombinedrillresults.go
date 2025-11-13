package draft

import (
	"context"
	"fmt"
	"log/slog"
)

// DraftCombineDrillResultsParams holds parameters for the DraftCombineDrillResults endpoint.
type DraftCombineDrillResultsParams struct {
	LeagueId string
	SeasonYear string
}

// GetDraftCombineDrillResults fetches data from the draftcombinedrillresults endpoint.
func (c *Client) GetDraftCombineDrillResults(ctx context.Context, params DraftCombineDrillResultsParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching draftcombinedrillresults")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
		"SeasonYear": params.SeasonYear,
	}

	resp, err := c.httpClient.SendRequest(ctx, "draftcombinedrillresults", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch draftcombinedrillresults",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch draftcombinedrillresults: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from draftcombinedrillresults endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal draftcombinedrillresults response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched draftcombinedrillresults",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
