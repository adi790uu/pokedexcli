package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var apiURL string = "https://pokeapi.co/api/v2/location-area/"

type cliCommand struct {
	Name string
	Description string
	Callback func() error
}

type locations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type config struct {
	Previous string
	Next string
}

func commandHelp() (error){
	fmt.Println("Welcome to the Pokedex!\nUsage:\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return nil
}

func Map(con *config) (error){
	locs, err := fetchLocations(con.Next, con)

	if err != nil {
		fmt.Println("Some error occurred", err)
		return err
	}

	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func MapBack(con *config) (error) {

	if con.Previous == "" {
		fmt.Println("This is the starting point. No previous locations to map.")
		return nil
	}

	locs, err := fetchLocations(con.Previous, con)

	if err != nil {
		fmt.Println("Some error occurred", err)
		return err
	}

	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandExit() (error){
	os.Exit(0)
	return nil
}

func fetchLocations(url string, con *config) (locations, error) {
	res, err := http.Get(url)
	if err != nil {
		return locations{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return locations{}, err
	}

	var locs locations
	err = json.Unmarshal(body, &locs)

	if err != nil {
		return locations{}, err
	}

	con.Next = locs.Next
	con.Previous = locs.Previous

	return locs, nil
}

func mapper() map[string]cliCommand {
	var con =  config{
		Next: apiURL,
		Previous: "",
	}
	return map[string]cliCommand{
		"help": {
			Name: "help",
			Description: "Displays a help message",
			Callback: commandHelp,
		},
		"exit": {
			Name: "exit",
			Description: "Exit the pokedex",
			Callback: commandExit,
		},
		"map": {
			Name: "map",
			Description: "Get Next locations",
			Callback: func() error { return Map(&con) },
		},
		"mapb": {
			Name: "mapb",
			Description: "Get Previous locations",
			Callback: func() error { return MapBack(&con) },
		},
	}
}

func main() {

	m := mapper()

	for  {
		var i string
		fmt.Printf("pokedex > ")
		fmt.Scanln(&i)
		if command, found := m[i]; found {
			err := command.Callback()
			if err != nil {
				fmt.Println("Some error occurred!")
			}
		} else {
			fmt.Println("Unknown command:", i)
		}
	}
}