package internal

import "strings"

func ValidateAndPad(input string) string {
	if len(input) >= 16 {
		return input
	}

	padding := 16 - len(input)
	padString := strings.Repeat("1", padding)
	return input + padString
}
