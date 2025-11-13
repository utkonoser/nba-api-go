package main

import (
	"context"
	"fmt"
	"log"

	"github.com/utkonoser/nba-api-go/endpoints/player"
)

func main() {
	// Nikola Jokic ID: 203999
	playerID := "203999"
	playerName := "Nikola Jokic"

	fmt.Printf("Getting stats for: %s (ID: %s)\n\n", playerName, playerID)

	// Create player stats client
	statsClient := player.NewClient(nil)

	// Get career stats
	params := player.PlayerCareerStatsParams{
		PlayerId:  playerID,
		PerMode36: "PerGame",
	}

	response, err := statsClient.GetPlayerCareerStats(context.Background(), params)
	if err != nil {
		log.Fatalf("Failed to fetch player stats: %v", err)
	}

	// Get season totals
	seasonStats, err := response.GetDataSet("SeasonTotalsRegularSeason")
	if err != nil {
		log.Fatalf("Failed to get season stats: %v", err)
	}

	fmt.Println("Career Statistics (Per Game):")
	fmt.Println("================================================================================")
	fmt.Printf("%-10s %-5s %-6s %-6s %-6s %-6s %-6s %-6s\n",
		"Season", "GP", "MIN", "PTS", "REB", "AST", "STL", "BLK")
	fmt.Println("--------------------------------------------------------------------------------")

	statsData := seasonStats.ToMap()
	for _, season := range statsData {
		fmt.Printf("%-10v %-5v %-6v %-6v %-6v %-6v %-6v %-6v\n",
			season["SEASON_ID"],
			season["GP"],
			season["MIN"],
			season["PTS"],
			season["REB"],
			season["AST"],
			season["STL"],
			season["BLK"])
	}

	// Get career totals
	careerTotals, err := response.GetDataSet("CareerTotalsRegularSeason")
	if err == nil && careerTotals.RowCount() > 0 {
		career, _ := careerTotals.GetRow(0)
		fmt.Println("\nCareer Totals:")
		fmt.Printf("  Games Played: %v\n", career["GP"])
		fmt.Printf("  Points: %v\n", career["PTS"])
		fmt.Printf("  Rebounds: %v\n", career["REB"])
		fmt.Printf("  Assists: %v\n", career["AST"])
	}
}
