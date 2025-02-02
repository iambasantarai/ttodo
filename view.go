package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	todos   []Todo
	cursor  int
	zenMode bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.todos)-1 {
				m.cursor++
			}
		case "z":
			var cmd tea.Cmd
			if m.zenMode {
				cmd = tea.ExitAltScreen
			} else {
				cmd = tea.EnterAltScreen
			}
			m.zenMode = !m.zenMode
			return m, cmd
		case "a":
			// add todo
		case "u":
			// update todo
		case "t":
			// toggle todo
		case "r":
			// remove todo
		case "X":
			// clean completed todos
		}
	}

	return m, nil
}

func (m model) View() string {
	s := mood() + "\n"
	s += "TODOS\n\n"

	for i, todo := range m.todos {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %d [%t] %s\n", cursor, todo.Id, todo.Completed, todo.Title)
	}

	s += "\nPress \n z: to toggle zen mode. \n q: to quit.\n"

	return s
}
