package colorize

import (
	"fmt"
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
			fmt.Println()
			continue
		}

		for row := 0; row < 8; row++ {
			var rowString string
			for col := 0; col < len(char); col++ {
				post := int(char[col]-32) * 9
				if color != "" && positions[globalpos] == true {
					rowString += color + bannerlines[row+post] + "\033[0m"
				} else {
					rowString += bannerlines[row+post]
				}
				globalpos++
			}
			result += rowString + "\n"

		}

	}
	return result
}
