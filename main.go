package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	pokecache "github.com/jelijs51/pokedexcli/internal"
	"github.com/jelijs51/pokedexcli/pokeapi"
)

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(commands map[string]cliCommand) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	
	for _, cmd := range commands {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}
	return nil
}

func printLocationArea(locationArea pokeapi.LocationArea) {
	for _, location := range locationArea.Results {
		fmt.Println(location.Name)
	}
}

func commandMap(config *pokeapi.LocationAreaConfig, next bool, cache *pokecache.Cache) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if config.Next != nil && next {
		url = *config.Next
	}
	if config.Prev != nil && !next {
		url = *config.Prev
	}
	if config.Prev == nil && !next {
		fmt.Println("you're on the first page")
		return nil
	}
	if cachedData, ok := cache.Get(url); ok {
		var locationAreas pokeapi.LocationArea
		if err := json.Unmarshal(cachedData, &locationAreas); err != nil{
			return err
		}
		printLocationArea(locationAreas)
		config.Next = locationAreas.Next
		config.Prev = locationAreas.Previous
		return nil
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	cache.Add(url, body)
	var locationArea pokeapi.LocationArea
	if err := json.Unmarshal(body, &locationArea); err != nil {
		return err
	}
	config.Next = locationArea.Next
	config.Prev = locationArea.Previous
	printLocationArea(locationArea)
	return nil
}

func commandExplore (locationArea string) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v", locationArea)
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var locationDetail pokeapi.PokemonList
	if err := json.Unmarshal(body, &locationDetail); err != nil {
		return err
	}
	for _, pokemon := range locationDetail.PokemonEncounter {
		fmt.Println(pokemon.PokemonDetail.Name)
	}
	return nil
}

func commandCatch (pokemon string, pokedex *map[string]pokeapi.Pokemon) error {
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon)
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v/", pokemon)
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return err
	}
	if res.StatusCode == http.StatusNotFound {
		fmt.Println("Pokemon not found")
		return nil
	}
	if rand.Intn(100) < 50 {
		fmt.Printf("%v escaped!\n", pokemon)
		return nil
	}
	fmt.Printf("%v was caught!\n", pokemon)
	var catchedPokemon pokeapi.Pokemon
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&catchedPokemon); err != nil {
		return err
	}
	(*pokedex)[pokemon] = catchedPokemon
	return nil
}

func commandInspect (pokemon string, pokedex map[string]pokeapi.Pokemon) error {
	poke, ok := pokedex[pokemon]
	if !ok {
		fmt.Printf("%v not caught yet\n",pokemon)
		return nil
	}
	fmt.Printf("Name: %v\n", poke.Name)
	fmt.Printf("Height: %v\n", poke.Height)
	fmt.Printf("Weight: %v\n", poke.Weight)
	fmt.Println("Stats:")
	for _, pokemonStat := range poke.PokemonStats{
		fmt.Printf("  -%v: %v\n", pokemonStat.Stat.Name, pokemonStat.Value)
	}
	fmt.Println("Types:")
	for _, pokemonType := range poke.PokemonType{
		fmt.Printf("  -%v\n",pokemonType.Type.Name)
	}
	return nil
}

func commandPokedex (pokedex map[string]pokeapi.Pokemon) error {
	fmt.Println("Your Pokedex:")
	for key, _ := range pokedex {
		fmt.Printf(" - %v\n",key)
	}
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) []string{
	text = strings.ToLower(text)
	return strings.Fields(text)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cache := pokecache.NewCache(5 * time.Minute)
	var LocationAreaConfig pokeapi.LocationAreaConfig
	var area string
	var catchPokemon string
	var inspectPokemon string
	pokedex := make(map[string]pokeapi.Pokemon)
	getCommands := make(map[string]cliCommand)
	getCommands = map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "List of commands to user the repl",
			callback: func() error {
				return commandHelp(getCommands)
			},
		},
		"map": {
			name: "map",
			description: "Show 20 next location areas in pokemon world",
			callback: func() error {
				return commandMap(&LocationAreaConfig, true, cache)
			},
		},
		"mapb": {
			name: "mapb",
			description: "Show 20 previous location areas in pokemon world",
			callback: func() error {
				return commandMap(&LocationAreaConfig, false, cache)
			},
		},
		"explore": {
			name: "explore",
			description: "explore the location area and show list of all pokemon in the area",
			callback: func() error {
				return commandExplore(area)
			},
		},
		"catch": {
			name: "catch",
			description: "catch the pokemon in area",
			callback: func() error {
				return commandCatch(catchPokemon, &pokedex)
			},
		},
		"inspect": {
			name: "inspect",
			description: "inspect pokemon in your pokedex",
			callback: func() error {
				return commandInspect(inspectPokemon, pokedex)
			},
		},
		"pokedex": {
			name: "pokedex",
			description: "show all pokemon in pokedex",
			callback: func() error {
				return commandPokedex(pokedex)
			},
		},
	}
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanedInput := cleanInput(text)
		if len(cleanedInput) == 0{
			continue
		}
		commandName := cleanedInput[0]
		if commandName == "explore" {
			area = cleanedInput[1]
		}
		if commandName == "catch" {
			catchPokemon = cleanedInput[1]
		}
		if commandName == "inspect" {
			inspectPokemon = cleanedInput[1]
		}
		if command, ok := getCommands[commandName]; ok {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}			
		}else {
			fmt.Println("Unknown Command")
		}
	}
}