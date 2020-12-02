package str

import (
	"strings"
)

// CountCharactersInString counts the number of a certain character in a string
func CountCharactersInString(value string, character string) int {
	count := 0
	for _, c := range value {
		if string(c) == character {
			count++
		}
	}
	return count
}

// RemoveEmptyLines removes all empty lines in a string
func RemoveEmptyLines(value string) string {
	var result []string

	for _, line := range strings.Split(value, "\n") {
		if line != "" {
			result = append(result, line)
		}
	}

	return strings.Join(result, "\n")
}
