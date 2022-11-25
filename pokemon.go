package main

import (
	"embed"
	"fmt"
	"path"
	"strings"
)

//go:embed pokemons
var fs embed.FS

type Pokemon struct {
	id   int
	name string
}

func (p Pokemon) Picture() string {
	fileName := fmt.Sprintf("%s.txt", strings.ToLower(p.name))
	filePath := path.Join("pokemons", fileName)
	imageBytes, err := fs.ReadFile(filePath)
	if err != nil {
		return ""
	}
	return string(imageBytes)
}

var pokemons = []Pokemon{
	{
		id:   1,
		name: "Bulbasaur",
	},
	{
		id:   2,
		name: "Ivysaur",
	},
	{
		id:   3,
		name: "Venusaur",
	},
	{
		id:   4,
		name: "Charmander",
	},
	{
		id:   5,
		name: "Charmeleon",
	},
	{
		id:   6,
		name: "Charizard",
	},
	{
		id:   7,
		name: "Squirtle",
	},
	{
		id:   8,
		name: "Wartortle",
	},
	{
		id:   9,
		name: "Blastoise",
	},
	{
		id:   10,
		name: "Caterpie",
	},
	{
		id:   25,
		name: "Pikachu",
	},
}
