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

	if gamesHaveStarted(games) {
		return []table.Column{
			{Title: "Matchup", Width: 50},
			{Title: "Score", Width: 20},
			{Title: "Status", Width: 20},
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

	if gamesHaveStarted(games) {
		for _, game := range games {
			var row table.Row
			if game.IsComplete() {
				row = table.Row{
					renderMatchup(game.AwayTeamName(), game.HomeTeamName()),
					renderScore(game.AwayTeamAbbr(), game.HomeTeamAbbr(), game.AwayTeamScore(), game.HomeTeamScore()),
					fmt.Sprintln("F"),
				}
			} else if game.HasStarted() {
				stats := game.GetLinecore()
				row = table.Row{
					renderMatchup(game.AwayTeamName(), game.HomeTeamName()),
					renderScore(game.AwayTeamAbbr(), game.HomeTeamAbbr(), game.AwayTeamScore(), game.HomeTeamScore()),
					fmt.Sprintf("%s %s", stats.GetPeriodTimeLeft(), stats.GetPeriod()),
				}
			} else {
				row = table.Row{
					renderMatchup(game.AwayTeamName(), game.HomeTeamName()),
					fmt.Sprintln("-"),
					game.GetGameDate(),
				}
			}
			rows = append(rows, row)

		}
	} else {
		for _, game := range games {
			row := table.Row{
				renderMatchup(game.AwayTeamName(), game.HomeTeamName()),
				game.GetGameDate(),
			}
			rows = append(rows, row)
		}
	}

	return rows
}

func renderMatchup(away, home string) string {
	return fmt.Sprintf("%s @ %s", away, home)
}

func renderScore(away, home string, awayScore, homeScore int) string {
	return fmt.Sprintf("%s %d - %s %d", away, awayScore, home, homeScore)
}

// If any of the games have started
func gamesHaveStarted(games []data.Game) bool {
	started := false
	for _, game := range games {
		started = started || game.HasStarted()
	}
	return started
}
