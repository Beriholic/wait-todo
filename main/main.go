package main

import (
	"log"
	"wait-to-do/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := ui.InitModel()
	m.Load()
	p := tea.NewProgram(m)

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
