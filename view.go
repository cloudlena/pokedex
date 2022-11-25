package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var titleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#7D56F4")).
	Padding(2).
	Width(22).
	Align(lipgloss.Center)

var boldStyle = lipgloss.NewStyle().Bold(true)

var paddedStyle = lipgloss.NewStyle().Padding(4)

func (m model) View() string {
	var s string

	switch m.view {
	case viewDetails:
		s = m.viewDetails()
	case viewList:
		s = m.viewList()
	}

	return paddedStyle.Render(s)
}

func (m model) viewDetails() string {
	if m.loading {
		return fmt.Sprintf("\n\n %s Loading %s...", m.spinner.View(), m.pokemons[m.selected].name)
	}

	s := titleStyle.Render(pokemons[m.selected].name)
	s += fmt.Sprintf("\n\n%s: %v cm", boldStyle.Render("Height"), m.details.Height*30)
	s += fmt.Sprintf("\n%s: %v kg\n", boldStyle.Render("Weight"), m.details.Weight/2)

	types := make([]string, 0, len(m.details.Types))
	for _, t := range m.details.Types {
		types = append(types, t.Type.Name)
	}
	s += fmt.Sprintf("\n%s: %s\n", boldStyle.Render("Types"), strings.Join(types, ", "))

	abilities := make([]string, 0, len(m.details.Abilities))
	for _, a := range m.details.Abilities {
		abilities = append(abilities, a.Ability.Name)
	}
	s += fmt.Sprintf("\n%s: %s\n", boldStyle.Render("Abilities"), strings.Join(abilities, ", "))

	games := make([]string, 0, len(m.details.GameIndices))
	for _, i := range m.details.GameIndices {
		games = append(games, i.Version.Name)
	}
	s += fmt.Sprintf("\n%s: %s\n\n", boldStyle.Render("Games"), strings.Join(games, ", "))

	imgStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(typeColor(m.details.Types[0].Type.Name)))
	s += fmt.Sprintf("\n\n%s\n\n\n", imgStyle.Render(pokemons[m.selected].Picture()))

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

func typeColor(typeStr string) string {
	switch typeStr {
	case "electric":
		return "#FFE333"
	case "fire":
		return "#FA1212"
	case "grass":
		return "#05D60F"
	case "water":
		return "#0367E7"
	}
	return "#FFFFFF"
}
