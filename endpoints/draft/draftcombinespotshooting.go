package draft

import (
	"context"
	"fmt"
	"log/slog"
)

// DraftCombineSpotShootingParams holds parameters for the DraftCombineSpotShooting endpoint.
type DraftCombineSpotShootingParams struct {
	LeagueId string
	SeasonYear string
}

// GetDraftCombineSpotShooting fetches data from the draftcombinespotshooting endpoint.
func (c *Client) GetDraftCombineSpotShooting(ctx context.Context, params DraftCombineSpotShootingParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching draftcombinespotshooting")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
	}
	
	// Add optional parameters only if they are not empty
	if params.SeasonYear != "" {
		reqParams["SeasonYear"] = params.SeasonYear
	}

	resp, err := c.httpClient.SendRequest(ctx, "draftcombinespotshooting", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch draftcombinespotshooting",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch draftcombinespotshooting: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from draftcombinespotshooting endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal draftcombinespotshooting response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched draftcombinespotshooting",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
