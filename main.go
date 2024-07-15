package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	Name string
	Description string
	Callback func() error
}

func commandHelp() (error){
	fmt.Println("Welcome to the Pokedex!\nUsage:\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return nil
}

func commandExit() (error){
	os.Exit(0)
	return nil
}

func mapper() (map[string]cliCommand) {
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