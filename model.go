package main

type model struct {
	pokemons []Pokemon
	cursor   int
	selected int
}

func initialModel() model {
	return model{
		pokemons: pokemons,
		cursor:   0,
		selected: 0,
	}
}
