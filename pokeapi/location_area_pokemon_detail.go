package pokeapi

type PokemonList struct {
	PokemonEncounter []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	PokemonDetail PokemonName `json:"pokemon"`
}

type PokemonName struct {
	Name string `json:"name"`
}