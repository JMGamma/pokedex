package main


import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokecache"
	"time"
)

type cliCommand struct {
	name string
	description string
	callback func(*config, ...string) error
}

type Pokemon struct {
	name string
	seen bool
	caught int
	data PokemonData
}

var dex = make(map[string]Pokemon)

func main() {
	var commands = map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
				},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays 20 locations, advancing on each usage",
			callback: commandMapForward,
		},
		"mapb": {
			name: "map back",
			description: "Displays the previous 20 locations",
			callback: commandMapBack,
		},
		"explore": {
			name: "explore",
			description: "Displays all pokemon found in the provided location",
			callback: commandExploreLocation,
		},
		"catch": {
			name: "catch",
			description: "Attempts to catch the named pokemon",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Displays information about the named pokemon",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "Displays all pokemon you've seen or caught",
			callback: commandPokedex,
		},
		}
	var cfg config
	cfg.cache = pokecache.NewCache(5 * time.Second)
	
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		clean := cleanInput(input)
		cmd, ok := commands[clean[0]]
		args := clean[1:]
		if ok {
			cmd.callback(&cfg, args...)
		} else {
			fmt.Print("Unknown command")
		}
	}
}
