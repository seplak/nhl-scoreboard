package data

import (
	"encoding/json"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/seplak/nhl-scoreboard/utils"
)

type Teams struct {
	Copyright string `json:"copyright"`
	Teams     []Team
}

type Team struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Link  string `json:"link"`
	Venue struct {
		Name     string `json:"name"`
		Link     string `json:"link"`
		City     string `json:"city"`
		TimeZone struct {
			ID     string `json:"id"`
			Offset int    `json:"offset"`
			Tz     string `json:"tz"`
		} `json:"timeZone"`
	} `json:"venue"`
	Abbreviation    string `json:"abbreviation"`
	TeamName        string `json:"teamName"`
	LocationName    string `json:"locationName"`
	FirstYearOfPlay string `json:"firstYearOfPlay"`
	Division        struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"division"`
	Conference struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"conference"`
	Franchise struct {
		FranchiseID int    `json:"franchiseId"`
		TeamName    string `json:"teamName"`
		Link        string `json:"link"`
	} `json:"franchise"`
	ShortName       string `json:"shortName"`
	OfficialSiteURL string `json:"officialSiteUrl"`
	FranchiseID     int    `json:"franchiseId"`
	Active          bool   `json:"active"`
}

func fetchTeam(id int) Team {
	var teams Teams
	data := utils.MakeRequest("https://statsapi.web.nhl.com/api/v1/teams/" + strconv.Itoa(id))

	if err := json.Unmarshal(data, &teams); err != nil {
		log.Error("Could not unmarshal team JSON")
	}

	// There's always one team
	return teams.Teams[0]
}

func (t Team) TeamAbbreviation() string {
	return t.Abbreviation
}
