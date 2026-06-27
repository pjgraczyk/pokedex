package cli

import (
	"fmt"
	"os"

	"github.com/pjgraczyk/pokedexcli/internal/api"
	"github.com/pjgraczyk/pokedexcli/internal/constants"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func createRegistry() map[string]cliCommand {
	registry := map[string]cliCommand{}

	registry["exit"] = cliCommand{
		name:        "exit",
		description: fmt.Sprintf("Exit the %v", constants.AppName),
		callback: func() error {
			fmt.Printf(constants.Bold+constants.Red+"Closing the %v... Goodbye!"+constants.Newline+constants.Reset, constants.AppName)
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

	registry["map"] = cliCommand{
		name:        "map",
		description: "Lists all places",
		callback: func() error {
			data, err := api.FetchData[api.LocationAreaResponse](api.ApiBaseUrl, "location-area")
			if err != nil {
				return err
			}
			for _, el := range data.Results {
				fmt.Println(constants.Green + constants.Bold + el.Name + constants.Reset)
			}
			return nil
		},
	registry["explore"] = cliCommand{
		name:        "explore",
		description: "Explore a place given by map",
		callback: func() error {
			data, err := api.FetchData[api.LocationAreaResponse](api.ApiBaseUrl, "location-area")
			if err != nil {
				return err
			}
			for _, el := range data.Results {
				fmt.Println(constants.Green + constants.Bold + el.Name + constants.Reset)
			}
			return nil
		},
	}

	return registry
}

func callRegistryCallback(registry map[string]cliCommand, command string) {
	if cmd, ok := registry[command]; ok {
		cmd.callback()
	} else {
		fmt.Printf(constants.Red+constants.Bold+constants.Underline+"The command doesn't exist: "+"%s"+constants.Newline+constants.Reset, command)
	}
}
