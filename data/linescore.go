package data

import (
	"time"
)

type Linescore struct {
	Copyright                  string `json:"copyright"`
	CurrentPeriod              int    `json:"currentPeriod"`
	CurrentPeriodOrdinal       string `json:"currentPeriodOrdinal"`
	CurrentPeriodTimeRemaining string `json:"currentPeriodTimeRemaining"`
	Periods                    []struct {
		PeriodType string    `json:"periodType"`
		StartTime  time.Time `json:"startTime"`
		Num        int       `json:"num"`
		OrdinalNum string    `json:"ordinalNum"`
		Home       struct {
			Goals       int    `json:"goals"`
			ShotsOnGoal int    `json:"shotsOnGoal"`
			RinkSide    string `json:"rinkSide"`
		} `json:"home"`
		Away struct {
			Goals       int    `json:"goals"`
			ShotsOnGoal int    `json:"shotsOnGoal"`
			RinkSide    string `json:"rinkSide"`
		} `json:"away"`
	} `json:"periods"`
	ShootoutInfo struct {
		Away struct {
			Scores   int `json:"scores"`
			Attempts int `json:"attempts"`
		} `json:"away"`
		Home struct {
			Scores   int `json:"scores"`
			Attempts int `json:"attempts"`
		} `json:"home"`
	} `json:"shootoutInfo"`
	Teams struct {
		Home struct {
			Team struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"team"`
			Goals        int  `json:"goals"`
			ShotsOnGoal  int  `json:"shotsOnGoal"`
			GoaliePulled bool `json:"goaliePulled"`
			NumSkaters   int  `json:"numSkaters"`
			PowerPlay    bool `json:"powerPlay"`
		} `json:"home"`
		Away struct {
			Team struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"team"`
			Goals        int  `json:"goals"`
			ShotsOnGoal  int  `json:"shotsOnGoal"`
			GoaliePulled bool `json:"goaliePulled"`
			NumSkaters   int  `json:"numSkaters"`
			PowerPlay    bool `json:"powerPlay"`
		} `json:"away"`
	} `json:"teams"`
	PowerPlayStrength string `json:"powerPlayStrength"`
	HasShootout       bool   `json:"hasShootout"`
	IntermissionInfo  struct {
		IntermissionTimeRemaining int  `json:"intermissionTimeRemaining"`
		IntermissionTimeElapsed   int  `json:"intermissionTimeElapsed"`
		InIntermission            bool `json:"inIntermission"`
	} `json:"intermissionInfo"`
	PowerPlayInfo struct {
		SituationTimeRemaining int  `json:"situationTimeRemaining"`
		SituationTimeElapsed   int  `json:"situationTimeElapsed"`
		InSituation            bool `json:"inSituation"`
	} `json:"powerPlayInfo"`
}

func (l Linescore) GetPeriod() string {
	return l.CurrentPeriodOrdinal
}

func (l Linescore) GetPeriodTimeLeft() string {
	return l.CurrentPeriodTimeRemaining
}
