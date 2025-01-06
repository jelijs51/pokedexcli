# Pokédex CLI

A command-line interface Pokédex application that interacts with the PokeAPI to provide Pokémon information.

## Description

This CLI tool allows users to search and explore Pokémon data through a simple command-line interface. It features local caching to improve performance and reduce API calls.

## Features

- Interactive REPL (Read-Eval-Print Loop) interface
- Pokémon information lookup
- Local caching system
- Error handling for invalid inputs

## Requirements

- Go
- Internet connection (for initial API calls)

## Usage

Run the program:
```go run .```

### Available Commands:

- `help`: Display available commands
- `exit`: Exit the program
- `map`: Explore location to find pokemon
- `mapb`: Explore previous location
- `explore`: Explore area and get a list of pokemon
- `catch`: Catch a pokemon (not always success)
- `inspect`: Inspect pokemon from your pokedex
- `pokedex`: Show all pokemon in your pokedex 

## Cache System

The application implements a caching system that stores previously fetched data locally using map:
- Reduce API calls
- Improve response times
- Work offline with previously accessed data