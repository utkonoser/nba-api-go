package draft

import (
	"context"
	"fmt"
	"log/slog"
)

// DraftCombineNonStationaryShootingParams holds parameters for the DraftCombineNonStationaryShooting endpoint.
type DraftCombineNonStationaryShootingParams struct {
	LeagueId string
	SeasonYear string
}

// GetDraftCombineNonStationaryShooting fetches data from the draftcombinenonstationaryshooting endpoint.
func (c *Client) GetDraftCombineNonStationaryShooting(ctx context.Context, params DraftCombineNonStationaryShootingParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching draftcombinenonstationaryshooting")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
		"SeasonYear": params.SeasonYear,
	}

	resp, err := c.httpClient.SendRequest(ctx, "draftcombinenonstationaryshooting", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch draftcombinenonstationaryshooting",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch draftcombinenonstationaryshooting: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from draftcombinenonstationaryshooting endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal draftcombinenonstationaryshooting response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched draftcombinenonstationaryshooting",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
