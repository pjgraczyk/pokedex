package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

const (
	AppName   = "Pokedex"
	Reset     = "\033[0m"
	Black     = "\033[30m"
	Red       = "\033[31m"
	Green     = "\033[32m"
	Yellow    = "\033[33m"
	Blue      = "\033[34m"
	Magenta   = "\033[35m"
	Cyan      = "\033[36m"
	Gray      = "\033[37m"
	White     = "\033[97m"
	Bold      = "\033[1m"
	Italic    = "\033[3m"
	Underline = "\033[4m"
	Invert    = "\033[7m"
	Newline   = "\n"
)

func CleanInput(text string) []string {
	cleanString := strings.Trim(text, " ")
	cleanString = strings.ToLower(cleanString)
	for strings.Contains(cleanString, "  ") {
		cleanString = strings.ReplaceAll(cleanString, "  ", " ")
	}
	splitString := strings.Split(cleanString, " ")
	return splitString
}

func createRegistry() map[string]cliCommand {
	registry := map[string]cliCommand{}
	registry["exit"] = cliCommand{
		name:        "exit",
		description: fmt.Sprintf("Exit the %v", AppName),
		callback: func() error {
			fmt.Printf(Bold+Red+"Closing the %v... Goodbye!"+Newline+Reset, AppName)
			os.Exit(0)
			return nil
		},
	}
	registry["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback: func() error {
			fmt.Println("Usage:")
			for k, v := range registry {
				fmt.Printf("%s: %s\n", k, v.description)
			}
			return nil
		},
	}
	return registry
}

func callRegistryCallback(registry map[string]cliCommand, command string) {
	if cmd, ok := registry[command]; ok {
		fmt.Println()
		cmd.callback()
	} else {
		fmt.Printf(Red+Bold+Underline+strings.ToUpper("The command doesn't exist: %s!")+Newline+Reset,
			command)
	}
}

func REPL() {
	scanner := bufio.NewScanner(os.Stdin)
	registry := createRegistry()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for k, v := range registry {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	for {
		fmt.Printf(Magenta+"%v >"+Reset, AppName)
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		output := CleanInput(input)
		callRegistryCallback(registry, output[0])
		if scanner.Err() != nil {
			os.Exit(1)
		}
	}
}
