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
}

func InitialModel() MainModel {
	model := MainModel{
		// On init, load today's games
		Games:  data.FetchGames(utils.FormatDate(time.Now())),
		// Games:  data.FetchGames("2022-10-08"), // -- Future date with games and no scores
		// Games:  data.FetchGames("2022-03-13"), // -- Old date with games and scores
	}
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
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m MainModel) View() string {
	var ui string

	ui += m.Table.View()

	// TODO: Implement a help section
	ui += "\nPress q to quit.\n"

	return ui
}
