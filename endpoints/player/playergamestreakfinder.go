package player

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayerGameStreakFinderParams holds parameters for the PlayerGameStreakFinder endpoint.
type PlayerGameStreakFinderParams struct {
	ActiveStreaksOnlyNullable string
	ConferenceNullable string
	DateFromNullable string
	DateToNullable string
	DivisionSimpleNullable string
	DraftNumberNullable string
	DraftRoundNullable string
	DraftTeamIdNullable string
	DraftYearNullable string
	EqAstNullable string
	EqBlkNullable string
	EqDdNullable string
	EqDrebNullable string
	EqFg3aNullable string
	EqFg3mNullable string
	EqFg3PctNullable string
	EqFgaNullable string
	EqFgmNullable string
	EqFgPctNullable string
	EqFtaNullable string
	EqFtmNullable string
	EqFtPctNullable string
	EqMinutesNullable string
	EqOrebNullable string
	EqPfNullable string
	EqPtsNullable string
	EqRebNullable string
	EqStlNullable string
	EqTdNullable string
	EqTovNullable string
	GameIdNullable string
	GtAstNullable string
	GtBlkNullable string
	GtDdNullable string
	GtDrebNullable string
	GtFg3aNullable string
	GtFg3mNullable string
	GtFg3PctNullable string
	GtFgaNullable string
	GtFgmNullable string
	GtFgPctNullable string
	GtFtaNullable string
	GtFtmNullable string
	GtFtPctNullable string
	GtMinutesNullable string
	GtOrebNullable string
	GtPfNullable string
	GtPtsNullable string
	GtRebNullable string
	GtStlNullable string
	GtTdNullable string
	GtTovNullable string
	LeagueIdNullable string
	LocationNullable string
	LtAstNullable string
	LtBlkNullable string
	LtDdNullable string
	LtDrebNullable string
	LtFg3aNullable string
	LtFg3mNullable string
	LtFg3PctNullable string
	LtFgaNullable string
	LtFgmNullable string
	LtFgPctNullable string
	LtFtaNullable string
	LtFtmNullable string
	LtFtPctNullable string
	LtMinutesNullable string
	LtOrebNullable string
	LtPfNullable string
	LtPtsNullable string
	LtRebNullable string
	LtStlNullable string
	LtTdNullable string
	LtTovNullable string
	PlayerPositionAbbreviationNullable string
	SeasonNullable string
	SeasonTypeNullable string
	StarterBenchNullable string
	TeamIdNullable string
	VsConferenceNullable string
	VsDivisionNullable string
}

// GetPlayerGameStreakFinder fetches data from the playergamestreakfinder endpoint.
func (c *Client) GetPlayerGameStreakFinder(ctx context.Context, params PlayerGameStreakFinderParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playergamestreakfinder")

	reqParams := map[string]string{}
	
	// Add nullable parameters only if they are not empty
	if params.ActiveStreaksOnlyNullable != "" {
		reqParams["ActiveStreaksOnly"] = params.ActiveStreaksOnlyNullable
	}
	if params.ConferenceNullable != "" {
		reqParams["Conference"] = params.ConferenceNullable
	}
	if params.DateFromNullable != "" {
		reqParams["DateFrom"] = params.DateFromNullable
	}
	if params.DateToNullable != "" {
		reqParams["DateTo"] = params.DateToNullable
	}
	if params.DivisionSimpleNullable != "" {
		reqParams["Division"] = params.DivisionSimpleNullable
	}
	if params.DraftNumberNullable != "" {
		reqParams["DraftNumber"] = params.DraftNumberNullable
	}
	if params.DraftRoundNullable != "" {
		reqParams["DraftRound"] = params.DraftRoundNullable
	}
	if params.DraftTeamIdNullable != "" {
		reqParams["DraftTeamID"] = params.DraftTeamIdNullable
	}
	if params.DraftYearNullable != "" {
		reqParams["DraftYear"] = params.DraftYearNullable
	}
	if params.EqAstNullable != "" {
		reqParams["EqAST"] = params.EqAstNullable
	}
	if params.EqBlkNullable != "" {
		reqParams["EqBLK"] = params.EqBlkNullable
	}
	if params.EqDdNullable != "" {
		reqParams["EqDD"] = params.EqDdNullable
	}
	if params.EqDrebNullable != "" {
		reqParams["EqDREB"] = params.EqDrebNullable
	}
	if params.EqFg3aNullable != "" {
		reqParams["EqFG3A"] = params.EqFg3aNullable
	}
	if params.EqFg3mNullable != "" {
		reqParams["EqFG3M"] = params.EqFg3mNullable
	}
	if params.EqFg3PctNullable != "" {
		reqParams["EqFG3_PCT"] = params.EqFg3PctNullable
	}
	if params.EqFgaNullable != "" {
		reqParams["EqFGA"] = params.EqFgaNullable
	}
	if params.EqFgmNullable != "" {
		reqParams["EqFGM"] = params.EqFgmNullable
	}
	if params.EqFgPctNullable != "" {
		reqParams["EqFG_PCT"] = params.EqFgPctNullable
	}
	if params.EqFtaNullable != "" {
		reqParams["EqFTA"] = params.EqFtaNullable
	}
	if params.EqFtmNullable != "" {
		reqParams["EqFTM"] = params.EqFtmNullable
	}
	if params.EqFtPctNullable != "" {
		reqParams["EqFT_PCT"] = params.EqFtPctNullable
	}
	if params.EqMinutesNullable != "" {
		reqParams["EqMINUTES"] = params.EqMinutesNullable
	}
	if params.EqOrebNullable != "" {
		reqParams["EqOREB"] = params.EqOrebNullable
	}
	if params.EqPfNullable != "" {
		reqParams["EqPF"] = params.EqPfNullable
	}
	if params.EqPtsNullable != "" {
		reqParams["EqPTS"] = params.EqPtsNullable
	}
	if params.EqRebNullable != "" {
		reqParams["EqREB"] = params.EqRebNullable
	}
	if params.EqStlNullable != "" {
		reqParams["EqSTL"] = params.EqStlNullable
	}
	if params.EqTdNullable != "" {
		reqParams["EqTD"] = params.EqTdNullable
	}
	if params.EqTovNullable != "" {
		reqParams["EqTOV"] = params.EqTovNullable
	}
	if params.GameIdNullable != "" {
		reqParams["GameID"] = params.GameIdNullable
	}
	if params.GtAstNullable != "" {
		reqParams["GtAST"] = params.GtAstNullable
	}
	if params.GtBlkNullable != "" {
		reqParams["GtBLK"] = params.GtBlkNullable
	}
	if params.GtDdNullable != "" {
		reqParams["GtDD"] = params.GtDdNullable
	}
	if params.GtDrebNullable != "" {
		reqParams["GtDREB"] = params.GtDrebNullable
	}
	if params.GtFg3aNullable != "" {
		reqParams["GtFG3A"] = params.GtFg3aNullable
	}
	if params.GtFg3mNullable != "" {
		reqParams["GtFG3M"] = params.GtFg3mNullable
	}
	if params.GtFg3PctNullable != "" {
		reqParams["GtFG3_PCT"] = params.GtFg3PctNullable
	}
	if params.GtFgaNullable != "" {
		reqParams["GtFGA"] = params.GtFgaNullable
	}
	if params.GtFgmNullable != "" {
		reqParams["GtFGM"] = params.GtFgmNullable
	}
	if params.GtFgPctNullable != "" {
		reqParams["GtFG_PCT"] = params.GtFgPctNullable
	}
	if params.GtFtaNullable != "" {
		reqParams["GtFTA"] = params.GtFtaNullable
	}
	if params.GtFtmNullable != "" {
		reqParams["GtFTM"] = params.GtFtmNullable
	}
	if params.GtFtPctNullable != "" {
		reqParams["GtFT_PCT"] = params.GtFtPctNullable
	}
	if params.GtMinutesNullable != "" {
		reqParams["GtMINUTES"] = params.GtMinutesNullable
	}
	if params.GtOrebNullable != "" {
		reqParams["GtOREB"] = params.GtOrebNullable
	}
	if params.GtPfNullable != "" {
		reqParams["GtPF"] = params.GtPfNullable
	}
	if params.GtPtsNullable != "" {
		reqParams["GtPTS"] = params.GtPtsNullable
	}
	if params.GtRebNullable != "" {
		reqParams["GtREB"] = params.GtRebNullable
	}
	if params.GtStlNullable != "" {
		reqParams["GtSTL"] = params.GtStlNullable
	}
	if params.GtTdNullable != "" {
		reqParams["GtTD"] = params.GtTdNullable
	}
	if params.GtTovNullable != "" {
		reqParams["GtTOV"] = params.GtTovNullable
	}
	if params.LeagueIdNullable != "" {
		reqParams["LeagueID"] = params.LeagueIdNullable
	}
	if params.LocationNullable != "" {
		reqParams["Location"] = params.LocationNullable
	}
	if params.LtAstNullable != "" {
		reqParams["LtAST"] = params.LtAstNullable
	}
	if params.LtBlkNullable != "" {
		reqParams["LtBLK"] = params.LtBlkNullable
	}
	if params.LtDdNullable != "" {
		reqParams["LtDD"] = params.LtDdNullable
	}
	if params.LtDrebNullable != "" {
		reqParams["LtDREB"] = params.LtDrebNullable
	}
	if params.LtFg3aNullable != "" {
		reqParams["LtFG3A"] = params.LtFg3aNullable
	}
	if params.LtFg3mNullable != "" {
		reqParams["LtFG3M"] = params.LtFg3mNullable
	}
	if params.LtFg3PctNullable != "" {
		reqParams["LtFG3_PCT"] = params.LtFg3PctNullable
	}
	if params.LtFgaNullable != "" {
		reqParams["LtFGA"] = params.LtFgaNullable
	}
	if params.LtFgmNullable != "" {
		reqParams["LtFGM"] = params.LtFgmNullable
	}
	if params.LtFgPctNullable != "" {
		reqParams["LtFG_PCT"] = params.LtFgPctNullable
	}
	if params.LtFtaNullable != "" {
		reqParams["LtFTA"] = params.LtFtaNullable
	}
	if params.LtFtmNullable != "" {
		reqParams["LtFTM"] = params.LtFtmNullable
	}
	if params.LtFtPctNullable != "" {
		reqParams["LtFT_PCT"] = params.LtFtPctNullable
	}
	if params.LtMinutesNullable != "" {
		reqParams["LtMINUTES"] = params.LtMinutesNullable
	}
	if params.LtOrebNullable != "" {
		reqParams["LtOREB"] = params.LtOrebNullable
	}
	if params.LtPfNullable != "" {
		reqParams["LtPF"] = params.LtPfNullable
	}
	if params.LtPtsNullable != "" {
		reqParams["LtPTS"] = params.LtPtsNullable
	}
	if params.LtRebNullable != "" {
		reqParams["LtREB"] = params.LtRebNullable
	}
	if params.LtStlNullable != "" {
		reqParams["LtSTL"] = params.LtStlNullable
	}
	if params.LtTdNullable != "" {
		reqParams["LtTD"] = params.LtTdNullable
	}
	if params.LtTovNullable != "" {
		reqParams["LtTOV"] = params.LtTovNullable
	}
	if params.PlayerPositionAbbreviationNullable != "" {
		reqParams["PlayerPosition"] = params.PlayerPositionAbbreviationNullable
	}
	if params.SeasonNullable != "" {
		reqParams["Season"] = params.SeasonNullable
	}
	if params.SeasonTypeNullable != "" {
		reqParams["SeasonType"] = params.SeasonTypeNullable
	}
	if params.StarterBenchNullable != "" {
		reqParams["StarterBench"] = params.StarterBenchNullable
	}
	if params.TeamIdNullable != "" {
		reqParams["TeamID"] = params.TeamIdNullable
	}
	if params.VsConferenceNullable != "" {
		reqParams["VsConference"] = params.VsConferenceNullable
	}
	if params.VsDivisionNullable != "" {
		reqParams["VsDivision"] = params.VsDivisionNullable
	}

	resp, err := c.httpClient.SendRequest(ctx, "playergamestreakfinder", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch playergamestreakfinder",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch playergamestreakfinder: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from playergamestreakfinder endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal playergamestreakfinder response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched playergamestreakfinder",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}

