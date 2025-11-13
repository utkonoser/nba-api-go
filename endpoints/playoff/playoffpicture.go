package playoff

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayoffPictureParams holds parameters for the PlayoffPicture endpoint.
type PlayoffPictureParams struct {
	LeagueId string
	SeasonId string
}

// GetPlayoffPicture fetches data from the playoffpicture endpoint.
func (c *Client) GetPlayoffPicture(ctx context.Context, params PlayoffPictureParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playoffpicture")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
	}
	
	if params.SeasonId != "" {
		reqParams["SeasonID"] = params.SeasonId
	}

	resp, err := c.httpClient.SendRequest(ctx, "playoffpicture", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch playoffpicture",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch playoffpicture: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from playoffpicture endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal playoffpicture response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched playoffpicture",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
