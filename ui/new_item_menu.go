package ui

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type NewitemModel struct {
	ta       textarea.Model
	mainMenu *model
}

func InitNewitemModel(mainMenu *model) NewitemModel {
	ta := textarea.New()
	ta.Focus()

	ta.Prompt = ">"
	ta.CharLimit = 280

	ta.SetWidth(30)
	ta.SetHeight(1)

	// ta.FocusedStyle.CursorLine = lipgloss.NewStyle()
	ta.ShowLineNumbers = false
	return NewitemModel{
		ta:       ta,
		mainMenu: mainMenu,
	}
}

func (m NewitemModel) Init() tea.Cmd {
	return textarea.Blink
}

func (m NewitemModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q":
			return m, tea.Quit
		case "enter":
			m.mainMenu.choices = append(m.mainMenu.choices, m.ta.Value())
			m.mainMenu.Save()
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.ta, cmd = m.ta.Update(msg)

	return m, tea.Batch(cmd)
}

func (m NewitemModel) View() string {
	var s string
	s += m.ta.View()
	return s
}
