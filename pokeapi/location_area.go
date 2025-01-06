package pokeapi

type LocationArea struct {
    Count    int                  `json:"count"`
    Next     *string              `json:"next"`
    Previous *string             `json:"previous"`
    Results  []LocationAreaDetail `json:"results"`
}

type LocationAreaDetail struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}