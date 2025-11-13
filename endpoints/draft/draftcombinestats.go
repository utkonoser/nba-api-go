package draft

import (
	"context"
	"fmt"
	"log/slog"
)

// DraftCombineStatsParams holds parameters for the DraftCombineStats endpoint.
type DraftCombineStatsParams struct {
	LeagueId string
	SeasonAllTime string
}

// GetDraftCombineStats fetches data from the draftcombinestats endpoint.
func (c *Client) GetDraftCombineStats(ctx context.Context, params DraftCombineStatsParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching draftcombinestats")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
	}
	
	// Add optional parameters only if they are not empty
	if params.SeasonAllTime != "" {
		reqParams["SeasonYear"] = params.SeasonAllTime
	}

	resp, err := c.httpClient.SendRequest(ctx, "draftcombinestats", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch draftcombinestats",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch draftcombinestats: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from draftcombinestats endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal draftcombinestats response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched draftcombinestats",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
