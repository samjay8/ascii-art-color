package colorize

import "strings"

func ColoredPositions(input, substring string) []bool {

	// Create positions based on visible characters only (excluding newlines)
	// This matches how AsciiArt processes input after splitting by newline
	inputNoNewlines := strings.ReplaceAll(input, "\n", "")
	boolslice := make([]bool, len(inputNoNewlines)) // This gives a slice of bool having the length of the input without newlines

	if substring == "" {
		for i := range boolslice {
			boolslice[i] = true // If the substring is empty, make the position marker return true for all input
		}
		return boolslice

	} else {
		pos := 0
		for pos < len(inputNoNewlines) {
			find := strings.Index(inputNoNewlines[pos:], substring) // This locates the index of the first instance of substring in the input
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
