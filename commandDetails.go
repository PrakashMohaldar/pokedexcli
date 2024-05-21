package CommandDetails

import (
	"fmt"
)
type cliCommand struct {
	name string
	description string
	callback func()
}

func getCommands()map[string]cliCommand{
	return map[string]cliCommand{
		"help":{
			name: "help",
			description: "shows help information",
			callback: callbackHelp,
		}
		"exit":{
			name: "exit",
			description: "exits the application",
			callback: callbackExit,
		}
	}
}