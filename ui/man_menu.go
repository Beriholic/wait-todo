package main

import (
	"fmt"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	index     int
	choices   []string
	selsected map[int]bool
}

func InitModel() model {
	return model{
		choices:   []string{"item1", "item2", "item3", "item4", "item5", "item6", "item7", "item8", "item9", "item10"},
		selsected: make(map[int]bool),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "up", "k":
			if m.index > 0 {
				m.index--
			}
		case "down", "j":
			if m.index < len(m.choices)-1 {
				m.index++
			}
		case "enter", " ":
			if m.selsected[m.index] {
				delete(m.selsected, m.index)
			} else {
				m.selsected[m.index] = true
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	title := "todo list"
	view := strings.Builder{}

	view.WriteString(fmt.Sprintf("%s\n\n", title))

	for i, choice := range m.choices {
		cursor := " "
		if m.index == i {
			cursor = ">"
		}
		checked := " "
		if _, ok := m.selsected[i]; ok {
			checked = "x"
		}

		view.WriteString(fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice))
	}

	return view.String()
}

func main() {
	p := tea.NewProgram(InitModel())
	if err := p.Start(); err != nil {
		log.Fatal("error running program:", err)
	}
}
