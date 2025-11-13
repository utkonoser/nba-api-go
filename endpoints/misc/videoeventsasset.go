package misc

import (
	"context"
	"fmt"
	"log/slog"
)

// VideoEventsAssetParams holds parameters for the VideoEventsAsset endpoint.
type VideoEventsAssetParams struct {
	GameId string
	GameEventId string
}

// GetVideoEventsAsset fetches data from the videoeventsasset endpoint.
func (c *Client) GetVideoEventsAsset(ctx context.Context, params VideoEventsAssetParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching videoeventsasset")

	reqParams := map[string]string{
		"GameID": params.GameId,
		"GameEventID": params.GameEventId,
	}

	resp, err := c.httpClient.SendRequest(ctx, "videoeventsasset", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch videoeventsasset",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch videoeventsasset: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from videoeventsasset endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal videoeventsasset response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched videoeventsasset",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
