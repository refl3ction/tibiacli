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

	highscores := []struct {
		Key   string
		Value string
	}{
		{"Achievements", "achievements"},
		{"Axe Fighting", "axefighting"},
		{"Charm Points", "charmpoints"},
		{"Club Fighting", "clubfighting"},
		{"Distance Fighting", "distancefighting"},
		{"Experience", "experience"},
		{"Fishing", "fishing"},
		{"Fist Fighting", "fistfighting"},
		{"Goshnars Taint", "goshnarstaint"},
		{"Loyalty Points", "loyaltypoints"},
		{"Magic Level", "magiclevel"},
		{"Shielding", "shielding"},
		{"Sword Fighting", "swordfighting"},
		{"Drome Score", "dromescore"},
		{"Boss Points", "bosspoints"},
	}

	var options []huh.Option[string]
	for _, hs := range highscores {
		options = append(options, huh.NewOption(hs.Key, hs.Value))
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("Highscores (Top 50)").Options(
				options...,
			).Value(&highscore),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	getData(highscore)
}

func getData(category string) {
	var res Root
	vocation := "all"
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
