package main

import (
	"fmt"
)

func (m model) View() string {
	var s string

	s += "What is your starter Pokémon?\n\n"

	for i, p := range m.pokemons {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		selected := " "
		if i == m.selected {
			selected = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, selected, p.name)
	}

	s += fmt.Sprintf("\n\n%s\n", m.pokemons[m.selected].Picture())

	return s
}
