package misc

import (
	"context"
	"fmt"
	"log/slog"
)

// VideoEventsParams holds parameters for the VideoEvents endpoint.
type VideoEventsParams struct {
	GameId string
	GameEventId string
}

// GetVideoEvents fetches data from the videoevents endpoint.
func (c *Client) GetVideoEvents(ctx context.Context, params VideoEventsParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching videoevents")

	reqParams := map[string]string{
		"GameID": params.GameId,
		"GameEventID": params.GameEventId,
	}

	resp, err := c.httpClient.SendRequest(ctx, "videoevents", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch videoevents",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch videoevents: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from videoevents endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal videoevents response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched videoevents",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
