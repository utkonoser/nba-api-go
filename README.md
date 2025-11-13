# NBA API Go Client

[![Go Version](https://img.shields.io/badge/Go-1.25.1-blue.svg)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Linter](https://github.com/utkonoser/nba-api-go/workflows/CI/badge.svg?branch=main&event=push)](https://github.com/utkonoser/nba-api-go/actions)
[![Unit Tests](https://github.com/utkonoser/nba-api-go/workflows/CI/badge.svg?branch=main&event=push)](https://github.com/utkonoser/nba-api-go/actions)
[![Coverage](https://img.shields.io/badge/coverage-ready-blue)](https://github.com/utkonoser/nba-api-go/actions)

A comprehensive Go client library for accessing NBA.com APIs. This library provides easy access to NBA statistics, live game data, player information, and team data.

## Features

- 🏀 **Live Data API**: Real-time game scores, box scores, play-by-play, and odds
- 📊 **Stats API**: 125 comprehensive endpoints covering all NBA data
- 🔍 **Type-Safe**: Fully typed responses with Go structs
- 🎯 **Easy Search**: Find players and teams through API endpoints
- ✅ **Well-Tested**: 174 tests (91 unit + 83 integration) - all passing
- 🎯 **Google Style Guide**: Follows Go best practices and Google's style guide

## Installation

```bash
go get github.com/utkonoser/nba-api-go
```

## Examples

### Find Players

Search for players by name using the `CommonAllPlayers` endpoint:

```go
package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/utkonoser/nba-api-go/endpoints/player"
)

func main() {
	client := player.NewClient(nil)

	// Get all NBA players
	params := player.CommonAllPlayersParams{
		IsOnlyCurrentSeason: "0",  // 0 = all players, 1 = current season only
		LeagueId:            "00",
		Season:              "2023-24",
	}

	response, err := client.GetCommonAllPlayers(context.Background(), params)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	dataset, _ := response.GetDataSet("CommonAllPlayers")
	players := dataset.ToMap()

	// Search for a player
	for _, playerData := range players {
		fullName := fmt.Sprintf("%v", playerData["DISPLAY_FIRST_LAST"])
		if strings.Contains(strings.ToLower(fullName), "jokic") {
			fmt.Printf("Found: %s (ID: %.0f)\n", fullName, playerData["PERSON_ID"])
			fmt.Printf("Team: %s\n", playerData["TEAM_NAME"])
			break
		}
	}
}
```

See [`examples/find_players`](examples/find_players) for more search examples.

## API Reference

### Live Data Package

The `live` package provides access to real-time NBA data:

- **`GetScoreboard(ctx)`** - Live scoreboard

### Stats Package

The `endpoints` package provides access to NBA statistics with **125 endpoints** organized into 13 logical packages:

#### Popular Endpoints:
- **`GetPlayerCareerStats(ctx, params)`** - Player career statistics
- **`GetCommonPlayerInfo(ctx, params)`** - Player information
- **`GetLeagueDashPlayerStats(ctx, params)`** - League-wide player statistics
- **`GetBoxScoreTraditionalV2(ctx, params)`** - Traditional box score
- **`GetShotChartDetail(ctx, params)`** - Shot chart data
- **`GetPlayByPlayV2(ctx, params)`** - Play-by-play data
- **`GetLeagueStandings(ctx, params)`** - League standings
- **`GetScoreboardV2(ctx, params)`** - Scoreboard data
- **`GetTeamDashLineups(ctx, params)`** - Team lineup statistics
- **`GetPlayerGameLog(ctx, params)`** - Player game log

#### All Available Endpoints (125 total):

<details>
<summary>Click to expand full endpoint list</summary>

**Box Score Endpoints:**
- AllTimeLeadersGrids, AssistLeaders, AssistTracker
- BoxScoreAdvancedV2, BoxScoreAdvancedV3
- BoxScoreDefensiveV2, BoxScoreFourFactorsV2, BoxScoreFourFactorsV3
- BoxScoreHustleV2, BoxScoreMatchupsV3
- BoxScoreMiscV2, BoxScoreMiscV3
- BoxScorePlayerTrackV2, BoxScorePlayerTrackV3
- BoxScoreScoringV2, BoxScoreScoringV3
- BoxScoreSummaryV2, BoxScoreSummaryV3
- BoxScoreTraditionalV2, BoxScoreTraditionalV3
- BoxScoreUsageV2, BoxScoreUsageV3

**Player Endpoints:**
- CommonAllPlayers, CommonPlayerInfo, CommonPlayoffSeries
- CumeStatsPlayer, CumeStatsPlayerGames
- PlayerAwards, PlayerCareerByCollegeRollup, PlayerCareerStats
- PlayerCompare
- PlayerDashboardByClutch, PlayerDashboardByGameSplits
- PlayerDashboardByGeneralSplits, PlayerDashboardByLastNGames
- PlayerDashboardByShootingSplits, PlayerDashboardByTeamPerformance
- PlayerDashboardByYearOverYear
- PlayerDashPtPass, PlayerDashPtReb, PlayerDashPtShotDefend, PlayerDashPtShots
- PlayerEstimatedMetrics, PlayerFantasyProfile, PlayerFantasyProfileBarGraph
- PlayerGameLog, PlayerGameLogs, PlayerGameStreakFinder
- PlayerIndex, PlayerProfileV2, PlayerVsPlayer

**Team Endpoints:**
- CommonTeamRoster, CommonTeamYears
- CumeStatsTeam, CumeStatsTeamGames
- TeamDashboardByGeneralSplits, TeamDashboardByShootingSplits
- TeamDashLineups, TeamDashPtPass, TeamDashPtReb
- TeamDetails, TeamEstimatedMetrics
- TeamHistoricalLeaders, TeamInfoCommon
- TeamPlayerDashboard, TeamPlayerOnOffDetails, TeamPlayerOnOffSummary
- TeamVsPlayer, TeamYearByYearStats

**League Endpoints:**
- LeagueDashOppPtShot
- LeagueDashPlayerBioStats
- LeagueDashPlayerPtShot, LeagueDashPlayerStats
- LeagueDashPtTeamDefend
- LeagueDashTeamPtShot
- LeagueGameFinder, LeagueGameLog
- LeagueLeaders
- LeagueStandings, LeagueStandingsV3

**Draft Endpoints:**
- DraftCombineDrillResults, DraftCombineNonStationaryShooting
- DraftCombinePlayerAnthro, DraftCombineSpotShooting
- DraftCombineStats, DraftHistory

**Game Endpoints:**
- GameRotation, HustleStatsBoxScore
- PlayByPlay, PlayByPlayV2, PlayByPlayV3
- ScoreboardV2, ScoreboardV3

**Shot Chart Endpoints:**
- ShotChartDetail, ShotChartLeagueWide, ShotChartLineupDetail

**Franchise Endpoints:**
- FranchiseHistory, FranchiseLeaders, FranchisePlayers

**Other Endpoints:**
- FantasyWidget
- HomePage, HomePageV2
- InfographicFanDuelPlayer, ISTStandings
- MatchupsRollup, PlayoffPicture
- ScheduleLeagueV2, ScheduleLeagueV2Int
- SynergyPlayTypes
</details>

### Unit Tests ✅
Fast, reliable tests using mock HTTP servers - **all 91 unit tests passing**:

```bash
# Run all unit tests
make test

# With coverage
make test-coverage

# With verbose output
make test-verbose
```

### Integration Tests ✅
Tests against the actual NBA API to verify data structures and compatibility:

```bash
# Run all integration tests
make test-integration

# Run specific package
go test -v -tags=integration ./endpoints/player -timeout 60s

# Run specific endpoint
go test -v -tags=integration -run TestGetPlayerCareerStats ./endpoints/player
```

**Note**: Integration tests can also be run manually via GitHub Actions workflow. They are not run automatically on every push to avoid unnecessary API calls and rate limiting.

### CI/CD Pipeline 🚀
This project uses GitHub Actions for continuous integration:

- **Linter**: Automatically runs `golangci-lint` on every push and pull request
- **Unit Tests**: Automatically runs all unit tests with coverage on every push and pull request
- **Integration Tests**: Can be manually triggered via GitHub Actions workflow (Actions → CI → Run workflow)

To run integration tests manually:
1. Go to the [Actions](https://github.com/utkonoser/nba-api-go/actions) tab
2. Select the "CI" workflow
3. Click "Run workflow"
4. Choose the branch and click "Run workflow"

**Test Statistics**:
- ✅ **91/91 unit tests** passing (100%)
- ✅ **83/83 integration tests** passing (100%)
- ⏭️ **0 integration tests** skipped
- ❌ **0 integration tests** failing
- 🎯 **Total: 174/174 tests** passing (100% success rate)

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Disclaimer

This library is not affiliated with or endorsed by the NBA. All data is accessed from publicly available NBA.com APIs. Please refer to [NBA.com Terms of Use](https://www.nba.com/termsofuse) for information about data usage.

## Credits

This Go client is inspired by the Python [nba_api](https://github.com/swar/nba_api) library.

## Support

For questions, issues, or feature requests, please open an issue on GitHub.

---

Made with ❤️ for NBA fans and Go developers

