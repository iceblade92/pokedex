package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	str := text
	str = strings.ToLower(str)
	str = strings.TrimSpace(str)
	result := strings.Fields(str)
	return result
}

func Repl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		str := cleanInput(text)
		if len(str) == 0 {
			fmt.Println("No text prompt provided")
			continue
		}
		fmt.Printf("Your command was: %s\n", str[0])

	}
}
