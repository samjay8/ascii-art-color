package colorize

import (
	"strings"
)

func AsciiArt(input string, bannerlines []string, color string, positions []bool) string {
	if input == "" {
		return ""
	}
	if input == `\n` {
		return "\n"
	}

	textsplit := strings.Split(input, `\n`)

	var result string
	globalpos := 0

	for _, char := range textsplit {
		if char == "" {
			result += "\n"
			globalpos++
			continue
		}
		rowString := make([]string, 8)

		for col := 0; col < len(char); col++ {
			post := int(char[col]-32) * 9

			for row := 0; row < 8; row++ {

				if color != "" && positions[globalpos] == true {
					rowString[row] += color + bannerlines[row+post] + "\033[0m"
				} else {
					rowString[row] += bannerlines[row+post]
				}
			}
			globalpos++
		}
		for _, row := range rowString {
			result += row + "\n"
		}

	}
	return result
}
