package main

import (
	"fmt"
	"strings"
)

func (m model) View() string {
	var s string

	switch m.view {
	case viewDetails:
		s = m.viewDetails()
	case viewList:
		s = m.viewList()
	}

	return s
}

func (m model) viewDetails() string {
	if m.loading {
		return fmt.Sprintf("\n\n %s Loading %s...", m.spinner.View(), m.pokemons[m.selected].name)
	}

	s := pokemons[m.selected].name
	s += fmt.Sprintf("\n\n%s: %v cm", "Height", m.details.Height*30)
	s += fmt.Sprintf("\n%s: %v kg\n", "Weight", m.details.Weight/2)

	types := make([]string, 0, len(m.details.Types))
	for _, t := range m.details.Types {
		types = append(types, t.Type.Name)
	}
	s += fmt.Sprintf("\n%s: %s\n", "Types", strings.Join(types, ", "))

	abilities := make([]string, 0, len(m.details.Abilities))
	for _, a := range m.details.Abilities {
		abilities = append(abilities, a.Ability.Name)
	}
	s += fmt.Sprintf("\n%s: %s\n", "Abilities", strings.Join(abilities, ", "))

	games := make([]string, 0, len(m.details.GameIndices))
	for _, i := range m.details.GameIndices {
		games = append(games, i.Version.Name)
	}
	s += fmt.Sprintf("\n%s: %s\n\n", "Games", strings.Join(games, ", "))

	s += fmt.Sprintf("\n\n%s\n\n\n", pokemons[m.selected].Picture())

	return s
}

func (m model) viewList() string {
	s := "What is your starter Pokémon?\n\n"

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

	return s
}
