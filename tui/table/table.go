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
		table.WithColumns(generateScheduleColumns()),
		table.WithRows(generateRowsFromData(model.games)),
		table.WithFocused(true),
		table.WithHeight(len(model.games) + 1),
	)
	return model
}

func generateScheduleColumns() []table.Column {
	return []table.Column{
		{Title: "Matchup", Width: 50},
		{Title: "Time", Width: 10},
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
		return []table.Row{
			{"No games today.", "N/A"},
		}
	}

	for _, game := range games {
		row := table.Row{
			fmt.Sprintf("%s @ %s", game.AwayTeamName(), game.HomeTeamName()),
			game.GetGameDate(),
		}

		rows = append(rows, row)
	}

	return rows
}
