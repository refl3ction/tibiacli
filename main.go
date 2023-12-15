package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/carlmjohnson/requests"
	"github.com/charmbracelet/huh"
	"github.com/olekukonko/tablewriter"
)

var (
	TIBIA_API_HOST = "https://api.tibiadata.com"
)

func main() {
	var category string
	categories := []string{
		"Highscores",
		"Characters",
	}
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("Select the category").Options(
				huh.NewOptions(categories...)...,
			).Value(&category),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	if category == "Highscores" {
		highscores()
	} else if category == "Characters" {
		fmt.Println("Characters")
	}
}

func highscores() {
	var highscore string
	var vocation string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("Highscores (Top 50)").Options(
				huh.NewOption("Achievements", "achievements"),
				huh.NewOption("Axe Fighting", "axefighting"),
				huh.NewOption("Charm Points", "charmpoints"),
				huh.NewOption("Club Fighting", "clubfighting"),
				huh.NewOption("Distance Fighting", "distancefighting"),
				huh.NewOption("Experience", "experience"),
				huh.NewOption("Fishing", "fishing"),
				huh.NewOption("Fist Fighting", "fistfighting"),
				huh.NewOption("Goshnars Taint", "goshnarstaint"),
				huh.NewOption("Loyalty Points", "loyaltypoints"),
				huh.NewOption("Magic Level", "magiclevel"),
				huh.NewOption("Shielding", "shielding"),
				huh.NewOption("Sword Fighting", "swordfighting"),
				huh.NewOption("Drome Score", "dromescore"),
				huh.NewOption("Boss Points", "bosspoints"),
			).Value(&highscore),
			huh.NewSelect[string]().Title("Filter by vocation").Options(
				huh.NewOption("All", "all").Selected(true),
				huh.NewOption("Druid", "druid"),
				huh.NewOption("Knight", "knight"),
				huh.NewOption("Paladin", "paladin"),
				huh.NewOption("Sorcerer", "sorcerer"),
			).Value(&vocation),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	getData(highscore, vocation)
}

func getData(category, vocation string) {
	var res Root
	err := requests.
		URL(TIBIA_API_HOST).
		Pathf("/v4/highscores/%s/%s/%s/%d", "all", category, vocation, 1).
		ToJSON(&res).
		Fetch(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	showOutput(res.Highscores)
}

func showOutput(highscores Highscores) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Rank", "Name", "Level", "Value", "Vocation", "World"})

	for _, hs := range highscores.HighscoreList {
		row := []string{
			fmt.Sprint(hs.Rank),
			hs.Name,
			fmt.Sprint(hs.Level),
			fmt.Sprint(hs.Value),
			hs.Vocation,
			hs.World,
		}
		table.Append(row)
	}

	table.Render()
}
