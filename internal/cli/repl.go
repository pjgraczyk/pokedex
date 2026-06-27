package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pjgraczyk/pokedexcli/internal/constants"
)

func REPL() {
	scanner := bufio.NewScanner(os.Stdin)
	registry := createRegistry()
	fmt.Println("Welcome to the Pokedex!")
	for {
		fmt.Printf(constants.Magenta+"%v > "+constants.Reset, constants.AppName)
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
