package data

import (
	"encoding/json"
	"time"
	"fmt"

	"github.com/seplak/nhl-scoreboard/utils"
	log "github.com/sirupsen/logrus"
)

type FullSchedule struct {
	Copyright    string `json:"copyright"`
	TotalItems   int    `json:"totalItems"`
	TotalEvents  int    `json:"totalEvents"`
	TotalGames   int    `json:"totalGames"`
	TotalMatches int    `json:"totalMatches"`
	MetaData     struct {
		TimeStamp string `json:"timeStamp"`
	} `json:"metaData"`
	Wait  int `json:"wait"`
	Dates []struct {
		Date         string `json:"date"`
		TotalItems   int    `json:"totalItems"`
		TotalEvents  int    `json:"totalEvents"`
		TotalGames   int    `json:"totalGames"`
		TotalMatches int    `json:"totalMatches"`
		Games        []struct {
			GamePk   int       `json:"gamePk"`
			Link     string    `json:"link"`
			GameType string    `json:"gameType"`
			Season   string    `json:"season"`
			GameDate time.Time `json:"gameDate"`
			Status   struct {
				AbstractGameState string `json:"abstractGameState"`
				CodedGameState    string `json:"codedGameState"`
				DetailedState     string `json:"detailedState"`
				StatusCode        string `json:"statusCode"`
				StartTimeTBD      bool   `json:"startTimeTBD"`
			} `json:"status"`
			Teams struct {
				Away struct {
					LeagueRecord struct {
						Wins   int    `json:"wins"`
						Losses int    `json:"losses"`
						Ot     int    `json:"ot"`
						Type   string `json:"type"`
					} `json:"leagueRecord"`
					Score int `json:"score"`
					Team  struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
						Link string `json:"link"`
					} `json:"team"`
				} `json:"away"`
				Home struct {
					LeagueRecord struct {
						Wins   int    `json:"wins"`
						Losses int    `json:"losses"`
						Ot     int    `json:"ot"`
						Type   string `json:"type"`
					} `json:"leagueRecord"`
					Score int `json:"score"`
					Team  struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
						Link string `json:"link"`
					} `json:"team"`
				} `json:"home"`
			} `json:"teams"`
			Venue struct {
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"venue,omitempty"`
			Content struct {
				Link string `json:"link"`
			} `json:"content"`
		} `json:"games"`
		Events  []interface{} `json:"events"`
		Matches []interface{} `json:"matches"`
	} `json:"dates"`
}

type Game struct {
	GamePk   int       `json:"gamePk"`
	Link     string    `json:"link"`
	GameType string    `json:"gameType"`
	Season   string    `json:"season"`
	GameDate time.Time `json:"gameDate"`
	Status   struct {
		AbstractGameState string `json:"abstractGameState"`
		CodedGameState    string `json:"codedGameState"`
		DetailedState     string `json:"detailedState"`
		StatusCode        string `json:"statusCode"`
		StartTimeTBD      bool   `json:"startTimeTBD"`
	} `json:"status"`
	Teams struct {
		Away struct {
			LeagueRecord struct {
				Wins   int    `json:"wins"`
				Losses int    `json:"losses"`
				Ot     int    `json:"ot"`
				Type   string `json:"type"`
			} `json:"leagueRecord"`
			Score int `json:"score"`
			Team  struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"team"`
		} `json:"away"`
		Home struct {
			LeagueRecord struct {
				Wins   int    `json:"wins"`
				Losses int    `json:"losses"`
				Ot     int    `json:"ot"`
				Type   string `json:"type"`
			} `json:"leagueRecord"`
			Score int `json:"score"`
			Team  struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"team"`
		} `json:"home"`
	} `json:"teams"`
	Venue struct {
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"venue,omitempty"`
	Content struct {
		Link string `json:"link"`
	} `json:"content"`
}

func fetchSchedule(date string) FullSchedule {
	data := utils.MakeRequest("https://statsapi.web.nhl.com/api/v1/schedule?date=" + date)

	var fullSchedule FullSchedule
	if err := json.Unmarshal(data, &fullSchedule); err != nil {
		log.Error("Could not unmarshal schedule JSON")
	}

	return fullSchedule
}

func FetchGames(date string) []Game {
	var games []Game

	// For now, we'll only ever get one day at a time
	daysSchedule := fetchSchedule(date)

	// No games today
	if len(daysSchedule.Dates) == 0 {
		return games
	}

	for _, game := range daysSchedule.Dates[0].Games {
		games = append(games, game)
	}

	return games
}

// Get the full name of the home team of a game
func (g *Game) HomeTeamName() string {
	return g.Teams.Home.Team.Name
}

// Get the full name of the away team of a game
func (g *Game) AwayTeamName() string {
	return g.Teams.Away.Team.Name
}

// Get the abbreviated name of the home team of a game
func (g *Game) HomeTeamAbbr() string {
	team := fetchTeam(g.Teams.Home.Team.ID)
	return team.TeamAbbreviation()
}

// Get the abbreviated name of the away team of a game
func (g *Game) AwayTeamAbbr() string {
	team := fetchTeam(g.Teams.Away.Team.ID)
	return team.TeamAbbreviation()
}

// Get the score of the away team of a game
func (g *Game) AwayTeamScore() int {
	return g.Teams.Away.Score
}

// Get the score of the home team of a game
func (g *Game) HomeTeamScore() int {
	return g.Teams.Home.Score
}

// Get the local time of a game
func (g *Game) GetGameDate() string {
	return g.GameDate.Local().Format("15:04:05")
}

// Whether or not a game is complete
func (g *Game) IsComplete() bool {
	return g.Status.AbstractGameState == "Final"
}

// Whether or not a game has started. A game has started
// if it is in progress or complete.
func (g *Game) HasStarted() bool {
	return g.IsComplete() || g.Status.AbstractGameState == "Live"
}

func (g *Game) GetLinecore() Linescore {
	url := fmt.Sprintf("https://statsapi.web.nhl.com/api/v1/game/%d/linescore", g.GamePk)
	data := utils.MakeRequest(url)

	var linescore Linescore
	if err := json.Unmarshal(data, &linescore); err != nil {
		log.Error("Could not unmarshal schedule JSON")
	}
	return linescore
}
