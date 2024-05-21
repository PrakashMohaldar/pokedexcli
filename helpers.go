package Helpers

import (
	"fmt"
	"strings"
)

func cleanInput(input string)[]string{
	lower := strings.ToLower(input)
	commands := strings.Fields(lower)
	return commands
}
func callbackHelp(){
	fmt.Println("Pokedex description")
	available := commandDetails.getCommands()
	for key, value := range available {
		fmt.Println(key, "  --  ", value.description)
	}
}
func callbackExit(){
	fmt.Println("Exiting Pokedex....")
	os.Exit(0)
}