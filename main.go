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
		"highscores",
		"characters",
	}
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("Category.").Options(
				huh.NewOptions(categories...)...,
			).Value(&category),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	if category == "highscores" {
		highscores()
	} else if category == "characters" {
		fmt.Println("characters")
	}
}

func highscores() {
	var highscore string
	highscores := []string{
		"achievements",
		"axefighting",
		"charmpoints",
		"clubfighting",
		"distancefighting",
		"experience",
		"fishing",
		"fistfighting",
		"goshnarstaint",
		"loyaltypoints",
		"magiclevel",
		"shielding",
		"swordfighting",
		"dromescore",
		"bosspoints",
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("Highscores.").Options(
				huh.NewOptions(highscores...)...,
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
