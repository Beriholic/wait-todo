package ui

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	index         int
	choices       []string
	selsected     map[int]bool
	isRunItemProg bool
}

func InitModel() model {
	return model{
		choices:   []string{},
		selsected: make(map[int]bool),
	}
}

func (m model) Save() {
	// Write choices to file
	choicesData := []byte(strings.Join(m.choices, "\n"))
	err := ioutil.WriteFile("../.data/data_choices.txt", choicesData, 0644)
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
	err = ioutil.WriteFile("../.data/data_selected.txt", selectedDataBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (m *model) Load() {
	// Read choices from file
	choicesData, err := ioutil.ReadFile("../.data/data_choices.txt")
	if err != nil {
		log.Fatal(err)
	}
	if string(choicesData) == "" {
		return
	}
	m.choices = strings.Split(string(choicesData), "\n")

	// Read selected from file
	selectedData, err := ioutil.ReadFile("../.data/data_selected.txt")
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

func (m *model) DeleteItem(index int) int {
	m.choices = append(m.choices[:index], m.choices[index+1:]...)
	delete(m.selsected, index)

	m.Save()

	for k := range m.selsected {
		if k > index {
			m.selsected[k-1] = true
			delete(m.selsected, k)
		}
	}
	switch index {
	case len(m.choices):
		return index - 1
	case 0:
		return 0
	}
	return index - 1
}

func (m *model) Newitem() {
	p := tea.NewProgram(InitNewitemModel(m))

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}

func (m *model) ResetItem() {
	m.choices = []string{}
	m.selsected = make(map[int]bool)
	m.Save()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
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
		case "d":
			if m.index >= 0 && m.index < len(m.choices) {
				m.index = m.DeleteItem(m.index)
			}
		case "n":
			if !m.isRunItemProg {
				m.Newitem()
			}
		case "r":
			m.ResetItem()
		}
	}
	return m, nil
}

func (m model) View() string {
	view := strings.Builder{}

	if len(m.choices) == 0 {
		view.WriteString("No items found.\n")
		view.WriteString("[N]ew")
		return view.String()
	}

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

	view.WriteString("\n")
	view.WriteString("[N]ew  ")
	view.WriteString("[D]el  ")
	view.WriteString("[R]eset")

	return view.String()
}
