package data

import (
	// "fmt"
	"net/http"
	"time"
	"io"
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

// const (
// 	baseEndpoint     = "https://statsapi.web.nhl.com/api/v1/"
// 	scheduleEndpoint = baseEndpoint + "schedule?date=2022-10-08"
// )

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
	resp, err := http.Get("https://statsapi.web.nhl.com/api/v1/schedule?date=" + date)
	if err != nil {
		log.Error("Could not get NHL game schedule.")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Could not read the schedule response body.")
	}

	var fullSchedule FullSchedule
	if err := json.Unmarshal(body, &fullSchedule); err != nil {
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

// Get the local time of a game
func (g *Game) GetGameDate() string {
	// TODO: Date formatting could/should come from a config
	return g.GameDate.Local().Format("15:04:05")
}


// func (g Game) HomeTeamAbbr() string {
// 	return g.Teams.Home.Team.
// }
