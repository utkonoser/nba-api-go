package league

import (
	"context"
	"fmt"
	"log/slog"
)

// LeagueGameFinderParams holds parameters for the LeagueGameFinder endpoint.
type LeagueGameFinderParams struct {
	PlayerOrTeamAbbreviation string
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
	OutcomeNullable string
	PoRoundNullable string
	PlayerIdNullable string
	RookieYearNullable string
	SeasonNullable string
	SeasonSegmentNullable string
	SeasonTypeNullable string
	StarterBenchNullable string
	TeamIdNullable string
	VsConferenceNullable string
	VsDivisionNullable string
	VsTeamIdNullable string
	YearsExperienceNullable string
}

// GetLeagueGameFinder fetches data from the leaguegamefinder endpoint.
func (c *Client) GetLeagueGameFinder(ctx context.Context, params LeagueGameFinderParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching leaguegamefinder")

	reqParams := map[string]string{
		"PlayerOrTeam": params.PlayerOrTeamAbbreviation,
		"Conference": params.ConferenceNullable,
		"DateFrom": params.DateFromNullable,
		"DateTo": params.DateToNullable,
		"Division": params.DivisionSimpleNullable,
		"DraftNumber": params.DraftNumberNullable,
		"DraftRound": params.DraftRoundNullable,
		"DraftTeamID": params.DraftTeamIdNullable,
		"DraftYear": params.DraftYearNullable,
		"EqAST": params.EqAstNullable,
		"EqBLK": params.EqBlkNullable,
		"EqDD": params.EqDdNullable,
		"EqDREB": params.EqDrebNullable,
		"EqFG3A": params.EqFg3aNullable,
		"EqFG3M": params.EqFg3mNullable,
		"EqFG3_PCT": params.EqFg3PctNullable,
		"EqFGA": params.EqFgaNullable,
		"EqFGM": params.EqFgmNullable,
		"EqFG_PCT": params.EqFgPctNullable,
		"EqFTA": params.EqFtaNullable,
		"EqFTM": params.EqFtmNullable,
		"EqFT_PCT": params.EqFtPctNullable,
		"EqMINUTES": params.EqMinutesNullable,
		"EqOREB": params.EqOrebNullable,
		"EqPF": params.EqPfNullable,
		"EqPTS": params.EqPtsNullable,
		"EqREB": params.EqRebNullable,
		"EqSTL": params.EqStlNullable,
		"EqTD": params.EqTdNullable,
		"EqTOV": params.EqTovNullable,
		"GameID": params.GameIdNullable,
		"GtAST": params.GtAstNullable,
		"GtBLK": params.GtBlkNullable,
		"GtDD": params.GtDdNullable,
		"GtDREB": params.GtDrebNullable,
		"GtFG3A": params.GtFg3aNullable,
		"GtFG3M": params.GtFg3mNullable,
		"GtFG3_PCT": params.GtFg3PctNullable,
		"GtFGA": params.GtFgaNullable,
		"GtFGM": params.GtFgmNullable,
		"GtFG_PCT": params.GtFgPctNullable,
		"GtFTA": params.GtFtaNullable,
		"GtFTM": params.GtFtmNullable,
		"GtFT_PCT": params.GtFtPctNullable,
		"GtMINUTES": params.GtMinutesNullable,
		"GtOREB": params.GtOrebNullable,
		"GtPF": params.GtPfNullable,
		"GtPTS": params.GtPtsNullable,
		"GtREB": params.GtRebNullable,
		"GtSTL": params.GtStlNullable,
		"GtTD": params.GtTdNullable,
		"GtTOV": params.GtTovNullable,
		"LeagueID": params.LeagueIdNullable,
		"Location": params.LocationNullable,
		"LtAST": params.LtAstNullable,
		"LtBLK": params.LtBlkNullable,
		"LtDD": params.LtDdNullable,
		"LtDREB": params.LtDrebNullable,
		"LtFG3A": params.LtFg3aNullable,
		"LtFG3M": params.LtFg3mNullable,
		"LtFG3_PCT": params.LtFg3PctNullable,
		"LtFGA": params.LtFgaNullable,
		"LtFGM": params.LtFgmNullable,
		"LtFG_PCT": params.LtFgPctNullable,
		"LtFTA": params.LtFtaNullable,
		"LtFTM": params.LtFtmNullable,
		"LtFT_PCT": params.LtFtPctNullable,
		"LtMINUTES": params.LtMinutesNullable,
		"LtOREB": params.LtOrebNullable,
		"LtPF": params.LtPfNullable,
		"LtPTS": params.LtPtsNullable,
		"LtREB": params.LtRebNullable,
		"LtSTL": params.LtStlNullable,
		"LtTD": params.LtTdNullable,
		"LtTOV": params.LtTovNullable,
		"Outcome": params.OutcomeNullable,
		"PORound": params.PoRoundNullable,
		"PlayerID": params.PlayerIdNullable,
		"RookieYear": params.RookieYearNullable,
		"Season": params.SeasonNullable,
		"SeasonSegment": params.SeasonSegmentNullable,
		"SeasonType": params.SeasonTypeNullable,
		"StarterBench": params.StarterBenchNullable,
		"TeamID": params.TeamIdNullable,
		"VsConference": params.VsConferenceNullable,
		"VsDivision": params.VsDivisionNullable,
		"VsTeamID": params.VsTeamIdNullable,
		"YearsExperience": params.YearsExperienceNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "leaguegamefinder", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch leaguegamefinder",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch leaguegamefinder: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from leaguegamefinder endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal leaguegamefinder response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched leaguegamefinder",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
