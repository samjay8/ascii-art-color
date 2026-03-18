package colorize

import (
	"strings"
)

func AsciiArt(input string, bannerlines []string, color string, positions []bool) string {
	if input == "" { // If input is empty, return empty
		return ""
	}
	if input == `\n` { // If input is equal to newline, return newline
		return "\n"
	}

	textsplit := strings.Split(input, `\n`) // Split the input based on the sep string `\n`

	var result string
	globalpos := 0 //This is the global position that increases per character

	for _, char := range textsplit {
		if char == "" {
			result += "\n" // This checks the character and if a character is empty, a new line should be appended into result
			globalpos++    // Move on from that forward to the next
			continue
		}
		rowString := make([]string, 8) // Create a variable that takes in a slice of string, with the length not more than 8 lines

		for col := 0; col < len(char); col++ { // For each character
			post := int(char[col]-32) * 9 // Calculates the position of exact character

			for row := 0; row < 8; row++ { // For each 8 segments the character has to pass through
				if color != "" && positions[globalpos] == true { // If the color is not empty and the character in position is true
					rowString[row] += color + bannerlines[post+1+row] + "\033[0m" // Append into rowString of the particular row the color, the bannnerlines, and the rest back to black.
				} else {
					rowString[row] += bannerlines[post+1+row] // Else return the bannerlines of that particular row into rowString.
				}
			}
			globalpos++ // Then loop over into the next character and iterate over again.
		}
		for _, row := range rowString {
			result += row + "\n" // In the range of rowString, append rowString plus a newline each into result.
		}

	}
	return result
}
