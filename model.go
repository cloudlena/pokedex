package main

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/cloudlena/pokedex/api"
)

type viewID int

const (
	viewList viewID = iota
	viewDetails
)

type setErrorMsg struct {
	err error
}

type setDetailsMsg struct {
	pokemon api.Pokemon
}

type model struct {
	view     viewID
	pokemons []Pokemon
	cursor   int
	selected int
	details  api.Pokemon
	spinner  spinner.Model
	loading  bool
	err      *error
}

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	return model{
		view:     viewList,
		pokemons: pokemons,
		cursor:   0,
		selected: 0,
		spinner:  s,
	}
}
