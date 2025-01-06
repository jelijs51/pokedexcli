package pokeapi

type Pokemon struct {
	Id int `json:"id"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Name string `json:"name"`
	PokemonStats []Stats `json:"stats"`
}

type Stats struct {
	Stat Stat `json:"stat"`
	Value int `json:"base_stat"`
}

type Stat struct {
	Name string `json:"string"`
}