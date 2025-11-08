package main

import (
	"bufio"
	"fmt"
	"os"
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
		fmt.Println("Your command was:", cleanInput[0])
	}
}
