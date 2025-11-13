package draft

import (
	"context"
	"fmt"
	"log/slog"
)

// DraftCombinePlayerAnthroParams holds parameters for the DraftCombinePlayerAnthro endpoint.
type DraftCombinePlayerAnthroParams struct {
	LeagueId string
	SeasonYear string
}

// GetDraftCombinePlayerAnthro fetches data from the draftcombineplayeranthro endpoint.
func (c *Client) GetDraftCombinePlayerAnthro(ctx context.Context, params DraftCombinePlayerAnthroParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching draftcombineplayeranthro")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
		"SeasonYear": params.SeasonYear,
	}

	resp, err := c.httpClient.SendRequest(ctx, "draftcombineplayeranthro", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch draftcombineplayeranthro",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch draftcombineplayeranthro: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from draftcombineplayeranthro endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal draftcombineplayeranthro response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched draftcombineplayeranthro",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
