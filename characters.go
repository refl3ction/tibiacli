package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/carlmjohnson/requests"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/olekukonko/tablewriter"
)

func characters() {
	var characterName string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Character name?").
				Prompt("> ").
				Value(&characterName),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	action := func() {
		getCharacter(characterName)
	}
	err = spinner.New().Title("Loading data...").Action(action).Run()
	if err != nil {
		log.Fatal(err)
	}
}

func getCharacter(characterName string) {
	var response Root
	err := requests.
		URL(TIBIA_API_HOST).
		Pathf("/v4/character/%s", characterName).
		ToJSON(&response).
		Fetch(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	printCharacter(response.CharacterData.Character)
}

func printCharacter(character Character) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetRowLine(true)
	table.SetHeader([]string{"Character Information"})

	// Adding rows for each field of the Character
	table.Append([]string{"Name", character.Name})
	table.Append([]string{"Title", character.Title})
	table.Append([]string{"Sex", character.Sex})
	table.Append([]string{"Vocation", character.Vocation})
	table.Append([]string{"Level", fmt.Sprint(character.Level)})
	table.Append([]string{"Achievement Points", fmt.Sprint(character.AchievementPoints)})
	table.Append([]string{"World", character.World})
	table.Append([]string{"Account Status", character.AccountStatus})
	// table.Append([]string{"Comment", character.Comment})
	table.Append([]string{"Deletion Date", character.DeletionDate})
	table.Append([]string{"Last Login", character.LastLogin})
	table.Append([]string{"Married To", character.MarriedTo})
	table.Append([]string{"Position", character.Position})
	table.Append([]string{"Residence", character.Residence})
	table.Append([]string{"Traded", fmt.Sprint(character.Traded)})
	table.Append([]string{"Unlocked Titles", fmt.Sprint(character.UnlockedTitles)})

	// Display guild info
	if character.Guild.Name != "" {
		table.Append([]string{"Guild Name", character.Guild.Name})
		table.Append([]string{"Guild Rank", character.Guild.Rank})
	}

	// Display houses
	for _, house := range character.Houses {
		table.Append([]string{"House", fmt.Sprintf("Name: %s, Town: %s", house.Name, house.Town)})
	}

	// Display former names and worlds, if any
	if len(character.FormerNames) > 0 {
		table.Append([]string{"Former Names", strings.Join(character.FormerNames, ", ")})
	}
	if len(character.FormerWorlds) > 0 {
		table.Append([]string{"Former Worlds", strings.Join(character.FormerWorlds, ", ")})
	}

	table.Render()
}
