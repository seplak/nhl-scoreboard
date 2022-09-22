package table

import (
	"strings"
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	"github.com/seplak/nhl-scoreboard/data"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	table table.Model
	games []data.Game

}

func (m Model) NewModel(games []data.Game) Model {
	model := Model{
		games: games,
	}

	model.table = table.New(
		table.WithColumns(generateColumnsFromData(model.games)),
		table.WithRows(generateRowsFromData(model.games)),
		table.WithFocused(true),
		table.WithHeight(len(model.games) + 1),
	)
	return model
}


func generateColumnsFromData(games []data.Game) []table.Column {
	if len(games) == 0 {
		return []table.Column{}
	}

	if allComplete(games) {
		return []table.Column{
			{Title: "Matchup", Width: 50},
			{Title: "Score", Width: 20},
		}
	} else {
		return []table.Column{
			{Title: "Matchup", Width: 50},
			{Title: "Time", Width: 10},
		}
	}
}


func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
	body := strings.Builder{}
	body.WriteString(m.table.View())
	return body.String()
}

func generateRowsFromData(games []data.Game) []table.Row {
	rows := []table.Row{}

	if len(games) == 0 {
		return []table.Row{}
	}

	if allComplete(games) {
		for _, game := range games {
			row := table.Row{
				fmt.Sprintf("%s @ %s", game.AwayTeamName(), game.HomeTeamName()),
				fmt.Sprintf("%s %d - %s %d", game.AwayTeamAbbr(), game.AwayTeamScore(), game.HomeTeamAbbr(), game.HomeTeamScore()),
			}
			rows = append(rows, row)
		}
	} else {
		for _, game := range games {
			row := table.Row{
				fmt.Sprintf("%s @ %s", game.AwayTeamName(), game.HomeTeamName()),
				game.GetGameDate(),
			}

			rows = append(rows, row)
		}
	}

	return rows
}

// Whether or not all of the games are complete
func allComplete(games []data.Game) bool {
	allComplete := true
	for _, game := range games {
		allComplete = allComplete && game.IsComplete()
	}
	return allComplete
}
