package colorize

import "strings"

func ColoredPositions(input, substring string) []bool {

	boolslice := make([]bool, len(input)) // This gives a slice of bool having the length of the input

	if substring == "" {
		for i := range boolslice {
			boolslice[i] = true // If the substring is empty, make the position marker return true for all input
		}
		return boolslice

	} else {
		for pos := 0; pos <= len(input); pos++ {
			find := strings.Index(input[pos:], substring) // This locates the index of the first instance of substring in the input
			if find == -1 {
				break // If it is at its end, it breaks!
			}
			for i := pos + find; i < pos+find+len(substring); i++ { // This locates the indexes of the remaining substring in the input and returns true for them
				boolslice[i] = true
			}
			pos = pos + find + len(substring) // Updates the position
		}

	}
	return boolslice
}
