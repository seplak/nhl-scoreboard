package tui

import (
	"time"

	"github.com/seplak/nhl-scoreboard/data"
	"github.com/seplak/nhl-scoreboard/utils"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/seplak/nhl-scoreboard/tui/components/table"

)

type MainModel struct {
	Games  []data.Game // The games for the selected day
	Table  table.Model
	Date   time.Time
}

func InitialModel() MainModel {
	// past := time.Date(2022, time.March, 15, 5, 0, 0, 0, time.UTC) // Past date with games and scores
	// future := time.Date(2022, time.October, 8, 5, 0, 0, 0, time.UTC) // Future date with games and scores
	model := MainModel{
		Date: time.Now(),
		// Date:  past, // -- Past date with games and scores
		// Date:  future, // -- Future date with games and no scores
	}
	model.Games = data.FetchGames(utils.FormatDate(model.Date))
	model.Table = model.Table.NewModel(model.Games)

	return model
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Exit the program
		case "ctrl+c", "q":
			return m, tea.Quit
		// Next day
		case "n":
			m.Date = m.Date.AddDate(0, 0, 1)
			m.refreshTable()
			return m, nil
		// Previous day
		case "p":
			m.Date = m.Date.AddDate(0, 0, -1)
			m.refreshTable()
			return m, nil
		// Refresh the data
		case "r":
			m.refreshTable()
			return m, nil
		}
	}
	return m, nil
}

func (m MainModel) View() string {
	var ui string

	ui = "Schedule for: " + utils.PrintDate(m.Date) + "\n\n"

	ui += m.Table.View()

	// TODO: Implement a help section
	ui += "\nPress n for next day, p for previous day,\nr to refresh, q to quit."

	return ui
}

func (m *MainModel) refreshTable() {
	m.Games = data.FetchGames(utils.FormatDate(m.Date))
	m.Table = m.Table.NewModel(m.Games)
}
