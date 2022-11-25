package main

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

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
			if m.cursor < len(m.pokemons)-1 {
				m.cursor++
			}

		case " ":
			m.selected = m.cursor

		case "enter":
			m.loading = true
			m.view = viewDetails
			cmd = m.fetchDetailsCmd

		case "esc":
			m.view = viewList
		}

	case setDetailsMsg:
		m.loading = false
		m.details = msg.pokemon

	case setErrorMsg:
		m.loading = false
		m.err = &msg.err

	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
	}

	return m, cmd
}
