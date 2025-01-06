package pokeapi

type PokemonList struct {
	PokemonEncounter []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	PokemonDetail Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
}