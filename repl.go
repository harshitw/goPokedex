package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := inputMessage(reader.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, exists := checkCommand()[commandName]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Print(err)
			}
		} else {
			fmt.Print("unknown command input")
			continue
		}

	}
}

func inputMessage(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name     string
	message  string
	callback func() error
}

func checkCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:     "help",
			message:  "Displays the help message",
			callback: commandHelp,
		},
		"exit": {
			name:     "exit",
			message:  "Exits the pokedex cli",
			callback: commandExit,
		},
	}
}

func commandHelp() error {

	fmt.Println()
	fmt.Println("Welcome to Pokedex a land of pokemons ...")

	fmt.Println()
	fmt.Println("Usage details : ")
	fmt.Println()

	for _, cmd := range checkCommand() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.message)
	}
	fmt.Println()

	return nil
}

func commandExit() error {

	fmt.Println("Exiting the pokedex cli ...")
	os.Exit(0)

	return nil
}
