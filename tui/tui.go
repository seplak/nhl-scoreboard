package tui

import (
	"time"
	"strings"

	"github.com/seplak/nhl-scoreboard/data"
	"github.com/seplak/nhl-scoreboard/utils"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/seplak/nhl-scoreboard/tui/table"
	"github.com/seplak/nhl-scoreboard/tui/help"

)

type MainModel struct {
	games  []data.Game // The games for the selected day
	table  table.Model
	date   time.Time
	help   help.Model
}

func InitialModel() MainModel {
	// past := time.Date(2022, time.March, 15, 5, 0, 0, 0, time.UTC) // Past date with games and scores
	// future := time.Date(2022, time.October, 8, 5, 0, 0, 0, time.UTC) // Future date with games and scores
	model := MainModel{
		date: time.Now(),
		// Date:  past, // -- Past date with games and scores
		// Date:  future, // -- Future date with games and no scores
		help: help.NewModel(),
	}
	model.games = data.FetchGames(utils.FormatDate(model.date))
	model.table = model.table.NewModel(model.games)

	return model
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var helpCmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Exit the program
		case "ctrl+c", "q":
			return m, tea.Quit
		// Next day
		case "n":
			m.date = m.date.AddDate(0, 0, 1)
			m.refreshTable()
			return m, nil
		// Previous day
		case "p":
			m.date = m.date.AddDate(0, 0, -1)
			m.refreshTable()
			return m, nil
		// Refresh the data
		case "r":
			m.refreshTable()
			return m, nil
		}
	}


	m.help, helpCmd = m.help.Update(msg)

	return m, tea.Batch(helpCmd)
}

func (m MainModel) View() string {
	ui := strings.Builder{}

	ui.WriteString("Schedule for: " + utils.PrintDate(m.date) + "\n\n")

	ui.WriteString(m.table.View())
	ui.WriteString("\n")
	ui.WriteString(m.help.View())

	return ui.String()
}

func (m *MainModel) refreshTable() {
	m.games = data.FetchGames(utils.FormatDate(m.date))
	m.table = m.table.NewModel(m.games)
}
