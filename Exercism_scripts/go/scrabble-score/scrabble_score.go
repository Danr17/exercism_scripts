package scrabble

import (
	"unicode"
)

var letters = map[rune]int{'A': 1, 'B': 3, 'C': 3, 'D': 2, 'E': 1, 'F': 4, 'G': 2, 'H': 4, 'I': 1, 'J': 8,
	'K': 5, 'L': 1, 'M': 3, 'N': 1, 'O': 1, 'P': 3, 'R': 1, 'Q': 10, 'S': 1, 'T': 1, 'U': 1, 'V': 4, 'W': 4, 'X': 8, 'Y': 4, 'Z': 10}

//Score is a Scrabble score
func Score(input string) int {

	number := 0
	for _, r := range input {
		number += letters[unicode.ToUpper(r)]
	}

	return number
}
