package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/the-1aw/pokemon-gogo/commands"
)

const Prompt = "Pokedex > "

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(Prompt)
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error reading stdin", err)
		}
		cleanInput := CleanInput(scanner.Text())
		if len(cleanInput) == 0 {
			continue
		}
		if cmd, ok := commands.Registry[cleanInput[0]]; ok == true {
			if err := cmd.Run(); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
