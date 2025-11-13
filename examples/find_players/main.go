package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/utkonoser/nba-api-go/endpoints/player"
)

func main() {
	fmt.Println("NBA Players Search Example")
	fmt.Println("=========================")
	fmt.Println()

	// Create player client
	client := player.NewClient(nil)

	// Get all NBA players
	params := player.CommonAllPlayersParams{
		IsOnlyCurrentSeason: "0",  // 0 = all players in history, 1 = current season only
		LeagueId:            "00", // NBA
		Season:              "2023-24",
	}

	ctx := context.Background()
	response, err := client.GetCommonAllPlayers(ctx, params)
	if err != nil {
		log.Fatalf("Failed to fetch players: %v", err)
	}

	// Get the players dataset
	dataset, err := response.GetDataSet("CommonAllPlayers")
	if err != nil {
		log.Fatalf("Failed to get dataset: %v", err)
	}

	players := dataset.ToMap()
	fmt.Printf("Total players in database: %d\n\n", len(players))

	// Example 1: Find player by name
	searchName := "Nikola Jokic"
	fmt.Printf("🔍 Searching for: %s\n", searchName)
	
	for _, playerData := range players {
		fullName := fmt.Sprintf("%v", playerData["DISPLAY_FIRST_LAST"])
		if strings.Contains(strings.ToLower(fullName), strings.ToLower(searchName)) {
			fmt.Printf("\n✅ Found!\n")
			fmt.Printf("  ID: %.0f\n", playerData["PERSON_ID"])
			fmt.Printf("  Name: %s\n", playerData["DISPLAY_FIRST_LAST"])
			fmt.Printf("  Team: %s\n", playerData["TEAM_NAME"])
			fmt.Printf("  Team Abbreviation: %s\n", playerData["TEAM_ABBREVIATION"])
			fmt.Printf("  From Year: %.0f\n", playerData["FROM_YEAR"])
			fmt.Printf("  To Year: %.0f\n", playerData["TO_YEAR"])
			fmt.Printf("  Roster Status: %.0f\n", playerData["ROSTERSTATUS"])
			break
		}
	}

	// Example 2: Find all players from a specific team
	fmt.Println("\n" + strings.Repeat("=", 50))
	teamName := "Lakers"
	fmt.Printf("\n🏀 All players from %s:\n\n", teamName)
	
	count := 0
	for _, playerData := range players {
		team := fmt.Sprintf("%v", playerData["TEAM_NAME"])
		if strings.Contains(strings.ToLower(team), strings.ToLower(teamName)) {
			rosterStatus := playerData["ROSTERSTATUS"]
			if fmt.Sprintf("%.0f", rosterStatus) == "1" { // Active roster
				fmt.Printf("  • %s\n", playerData["DISPLAY_FIRST_LAST"])
				count++
			}
		}
	}
	fmt.Printf("\nTotal active %s players: %d\n", teamName, count)

	// Example 3: Find players by last name
	fmt.Println("\n" + strings.Repeat("=", 50))
	lastName := "James"
	fmt.Printf("\n👥 All players with last name '%s':\n\n", lastName)
	
	for _, playerData := range players {
		fullName := fmt.Sprintf("%v", playerData["DISPLAY_FIRST_LAST"])
		if strings.Contains(fullName, lastName) {
			team := playerData["TEAM_NAME"]
			if team == "" || team == nil {
				team = "Free Agent"
			}
			fmt.Printf("  • %s - %v\n", fullName, team)
		}
	}
}

