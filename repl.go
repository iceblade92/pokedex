package main

import (
	"strings"
)

func cleanInput(text string) []string {
	str := text
	str = strings.ToLower(str)
	str = strings.TrimSpace(str)
	result := strings.Fields(str)
	return result
}
