package table

import (
	"strings"
	"fmt"

	"github.com/evertras/bubble-table/table"
	"github.com/seplak/nhl-scoreboard/data"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	columnKeyMatchup = "matchup"
	columnKeyTime    = "time"
)

type Model struct {
	table table.Model
	games []data.Game

}

func (m Model) NewModel(games []data.Game) Model {
	model := Model{
		games: games,
	}
	model.table = table.New(generateScheduleColumns()).WithRows(generateRowsFromData(model.games))

	return model
}

func generateScheduleColumns() []table.Column {
	return []table.Column{
		table.NewColumn(columnKeyMatchup, "Matchup", 50),
		table.NewColumn(columnKeyTime, "Time", 10),
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
			table.NewRow(table.RowData{
				columnKeyMatchup: "No games today.",
				columnKeyTime:    "N/A",
			}),
		}
	}

	for _, game := range games {
		row := table.NewRow(table.RowData{
			columnKeyMatchup: fmt.Sprintf("%s @ %s", game.AwayTeamName(), game.HomeTeamName()),
			columnKeyTime:    game.GetGameDate(),
		})

		rows = append(rows, row)
	}

	return rows
}
