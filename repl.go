package main

import (
	"fmt"
	"os"
	"pokedex/internal/pokecache"
	"strings"
	"math/rand"
)

type config struct {
	url *string
	next *string
	previous *string
	cache pokecache.Cache
}

func cleanInput(text string) []string {
	words := []string{}
	text = strings.TrimSpace(text)
	words = strings.Fields(text)
	cleanWords := []string{}
	for i := 0; i < len(words); i++ {
		word := words[i]
		word = strings.TrimSpace(word)
		word = strings.ToLower(word)
		cleanWords = append(cleanWords, word)
	}
	return cleanWords
}

func commandExit(cfg *config, s ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, s ...string) error {
			fmt.Print("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\n")
			return nil
}

func commandMapForward(cfg *config, s ...string) error {
	baseurl := "https://pokeapi.co/api/v2/location-area/"
	if cfg.url == nil{
		cfg.url = &baseurl
	} else {
		cfg.url = cfg.next
	}
	printout := getLocationAreaData(*cfg.url, cfg)
	fmt.Print(printout)
	return nil
}

func commandMapBack(cfg *config, s ...string) error {
	if cfg.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	cfg.url = cfg.previous
	printout := getLocationAreaData(*cfg.url, cfg)
	fmt.Print(printout)
	return nil

}

func commandExploreLocation(cfg *config, location ...string) error {
	if len(location) == 0 {
		fmt.Print("Please provide a location to explore")
		return nil
	}
	url := "https://pokeapi.co/api/v2/location-area/" + location[0]
	fmt.Printf("Exploring %s...\n", location[0])
	fmt.Println(getPokemonEncounters(url, cfg))
	return nil
}

func commandCatch(cfg *config, pokemon ...string) error {
	if len(pokemon) == 0 {
		fmt.Print("Please provide the pokemon you want to catch")
		return nil
	}
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon[0])
	rate := getPokemonCatchRate(url, cfg)
	if rand.Intn(rate) > (rate / 2){
		fmt.Printf("%s was caught!\n", pokemon[0])
		fmt.Printf("You may now inspect %s by entering 'inspect %s'\n", pokemon[0], pokemon[0])
		dexEntry := dex[pokemon[0]]
		dexEntry.caught++
		dexEntry.seen = true
		dex[pokemon[0]] = dexEntry
	} else {
		fmt.Printf("%s escaped!\n", pokemon[0])
		dexEntry := dex[pokemon[0]]
		dexEntry.seen = true
		dex[pokemon[0]] = dexEntry
		}
	return nil
}

func commandInspect(_ *config, pokemon ...string) error {
	if len(pokemon) == 0 {
		fmt.Print("Please provide the pokemon you want to inspect")
		return nil
	}
	dexEntry := dex[pokemon[0]]
	if !dexEntry.seen {
		fmt.Printf("You haven't seen %s yet!\n", pokemon[0])
		return nil
	} else {
		var types string
		data := dexEntry.data
		if len(data.Types) == 1 {
			types = data.Types[0].Type.Name
		} else {
			types = data.Types[0].Type.Name + "\n" + " - " + data.Types[1].Type.Name
		}
		fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n -hp: %d\n -attack: %d\n -defense: %d\n -special-attack: %d\n -special-defense: %d\n -speed: %d\nTypes:\n - %s\n", pokemon[0], data.Height, data.Weight, data.Stats[0].BaseStat, data.Stats[1].BaseStat, data.Stats[2].BaseStat, data.Stats[3].BaseStat, data.Stats[4].BaseStat, data.Stats[5].BaseStat, types)
	}
	return nil
}

func commandPokedex(_ *config, s ...string) error {
	var output string
	keys := make([]string, 0, len(dex))
	for key := range dex {
		keys = append(keys, key)
	}
	for i := 0; i < len(dex); i++ {
		pokemon := dex[keys[i]]
		output += fmt.Sprintf("- %s | Seen: %t | Caught: %d\n", pokemon.name, pokemon.seen, pokemon.caught)
	}
	fmt.Print(output)
	return nil
}