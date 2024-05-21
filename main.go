package main

import (
	"github.com/PrakashMohaldar/pokedexcli/CommandDetails"
	"github.com/PrakashMohaldar/pokedexcli/Helpers"
	"os"
	"fmt"
	"bufio"
	"strings"
)


func main(){
	reader := bufio.NewReader(os.Stdin)
	for{
		fmt.Print("pokedex>> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error readiing input:", err)
			continue
		}

		inputs := helpers.cleanInput(input)
		if len(inputs) == 0 {
			continue
		}
		inputCommand := inputs[0]

		available := commandDetails.getCommands()
		
		availCommand, ok := available[inputCommand]
		if !ok {
			fmt.Println("Unknown command: ")
			continue
		}
		
		availableCommand.callback()

	}
}