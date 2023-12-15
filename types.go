package main

type Root struct {
	Highscores    Highscores    `json:"highscores"`
	CharacterData CharacterData `json:"character"`
}

// Highscore represents an entry in the highscore list.
type Highscore struct {
	Level    int    `json:"level"`
	Name     string `json:"name"`
	Rank     int    `json:"rank"`
	Title    string `json:"title"`
	Value    int    `json:"value"`
	Vocation string `json:"vocation"`
	World    string `json:"world"`
}

// HighscorePage contains information about pagination in highscore list.
type HighscorePage struct {
	CurrentPage  int `json:"current_page"`
	TotalPages   int `json:"total_pages"`
	TotalRecords int `json:"total_records"`
}

// Highscores represents the highscores section of the JSON.
type Highscores struct {
	Category      string        `json:"category"`
	HighscoreAge  int           `json:"highscore_age"`
	HighscoreList []Highscore   `json:"highscore_list"`
	HighscorePage HighscorePage `json:"highscore_page"`
	Vocation      string        `json:"vocation"`
	World         string        `json:"world"`
}

type Badge struct {
	Description string `json:"description"`
	IconURL     string `json:"icon_url"`
	Name        string `json:"name"`
}

type AccountInformation struct {
	Created      string `json:"created"`
	LoyaltyTitle string `json:"loyalty_title"`
	Position     string `json:"position"`
}

type Achievement struct {
	Grade  int    `json:"grade"`
	Name   string `json:"name"`
	Secret bool   `json:"secret"`
}

type Guild struct {
	Name string `json:"name"`
	Rank string `json:"rank"`
}

type House struct {
	HouseID int    `json:"houseid"`
	Name    string `json:"name"`
	Paid    string `json:"paid"`
	Town    string `json:"town"`
}

type CharacterData struct {
	Character          `json:"character"`
	AccountInformation `json:"account_information"`
}

type Character struct {
	AccountStatus     string   `json:"account_status"`
	AchievementPoints int      `json:"achievement_points"`
	Comment           string   `json:"comment"`
	DeletionDate      string   `json:"deletion_date"`
	FormerNames       []string `json:"former_names"`
	FormerWorlds      []string `json:"former_worlds"`
	Guild             Guild    `json:"guild"`
	Houses            []House  `json:"houses"`
	LastLogin         string   `json:"last_login"`
	Level             int      `json:"level"`
	MarriedTo         string   `json:"married_to"`
	Name              string   `json:"name"`
	Position          string   `json:"position"`
	Residence         string   `json:"residence"`
	Sex               string   `json:"sex"`
	Title             string   `json:"title"`
	Traded            bool     `json:"traded"`
	UnlockedTitles    int      `json:"unlocked_titles"`
	Vocation          string   `json:"vocation"`
	World             string   `json:"world"`
}
