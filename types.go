package main

type Root struct {
	Highscores Highscores `json:"highscores"`
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
