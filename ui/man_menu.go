package main

import (
	"fmt"
	"io/ioutil"
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

func (m model) Save() {
	// Write choices to file
	choicesData := []byte(strings.Join(m.choices, "\n"))
	err := ioutil.WriteFile("../data/data_choices.txt", choicesData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Write selected to file
	var selectedData []string
	for k := range m.selsected {
		selectedData = append(selectedData, fmt.Sprintf("%d", k))
	}
	selectedDataStr := strings.Join(selectedData, "\n")
	selectedDataBytes := []byte(selectedDataStr)
	err = ioutil.WriteFile("../data/data_selected.txt", selectedDataBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (m *model) Load() {
	// Read choices from file
	choicesData, err := ioutil.ReadFile("../data/data_choices.txt")
	if err != nil {
		log.Fatal(err)
	}
	m.choices = strings.Split(string(choicesData), "\n")

	// Read selected from file
	selectedData, err := ioutil.ReadFile("../data/data_selected.txt")
	if err != nil {
		log.Fatal(err)
	}
	selectedDataStr := string(selectedData)
	if selectedDataStr != "" {
		selectedDataArr := strings.Split(selectedDataStr, "\n")
		for _, v := range selectedDataArr {
			if v != "" {
				k := 0
				fmt.Sscanf(v, "%d", &k)
				m.selsected[k] = true
			}
		}
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
			m.Save()
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
	m := InitModel()
	m.Load()
	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		log.Fatal("error running program:", err)
	}
}
