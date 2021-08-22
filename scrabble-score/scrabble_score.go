package scrabble

import "strings"

type lettersPoint struct {
	letters string
	points  int
}

var points = []lettersPoint{
	{"AEIOULNRST", 1},
	{"DG", 2},
	{"BCMP", 3},
	{"FHVWY", 4},
	{"K", 5},
	{"JX", 8},
	{"QZ", 10},
}

var pointMap = map[rune]int{
	'A': 1,
	'E': 1,
	'I': 1,
	'O': 1,
	'U': 1,
	'L': 1,
	'N': 1,
	'R': 1,
	'S': 1,
	'T': 1,
	'G': 2,
	'D': 2,
	'B': 3,
	'C': 3,
	'M': 3,
	'P': 3,
	'F': 4,
	'H': 4,
	'V': 4,
	'W': 4,
	'Y': 4,
	'K': 5,
	'J': 8,
	'X': 8,
	'Q': 10,
	'Z': 10,
}

//Score Map
func Score(input string) int {
	var res int = 0
	toUpper := strings.ToUpper(input)
	for _, letter := range toUpper {
		res = res + pointMap[letter]
	}
	return res
}

// Score test input
func ScoreStruct(input string) int {
	var res int = 0
	toUpper := strings.ToUpper(input)

	for _, letter := range toUpper {
		for _, candidate := range points {
			if strings.Contains(candidate.letters, string(letter)) {
				res = res + candidate.points
			}
		}
	}
	return res
}
