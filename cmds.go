package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/cloudlena/pokedex/api"
)

func (m model) fetchDetailsCmd() tea.Msg {
	details, err := api.GetPokemon(m.pokemons[m.selected].id)
	if err != nil {
		return setErrorMsg{err}
	}
	return setDetailsMsg{details}
}
