package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(commands map[string]cliCommand) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	
	for _, cmd := range commands {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
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
	}
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		text := scanner.Text()
		cleanedInput := cleanInput(text)
		if len(cleanedInput) == 0{
			continue
		}
		commandName := cleanedInput[0]
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