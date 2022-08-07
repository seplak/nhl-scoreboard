package main

import (
	log "github.com/sirupsen/logrus"
	tea "github.com/charmbracelet/bubbletea"
	tui "github.com/seplak/nhl-scoreboard/tui"
)


func main() {
	p := tea.NewProgram(tui.InitialModel(), tea.WithAltScreen())
    	if err := p.Start(); err != nil {
        	log.Error("Alas, there's been an error: %v", err)
    	}
}
