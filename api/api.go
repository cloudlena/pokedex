package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Pokemon struct {
	Height      int32       `json:"height"`
	Weight      int32       `json:"weight"`
	Abilities   []Ability   `json:"abilities"`
	GameIndices []GameIndex `json:"game_indices"`
	Types       []Type      `json:"types"`
}

type Ability struct {
	Ability struct {
		Name string `json:"name"`
	} `json:"ability"`
}

type GameIndex struct {
	Version struct {
		Name string `json:"name"`
	} `json:"version"`
}

type Type struct {
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

func GetPokemon(id int) (Pokemon, error) {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + fmt.Sprint(id))
	if err != nil {
		return Pokemon{}, fmt.Errorf("error getting pokemon %v from API: %w", id, err)
	}
	defer resp.Body.Close()

	var pokemon Pokemon
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error decoding API response to JSON: %w", err)
	}

	return pokemon, nil
}
