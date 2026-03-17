package colorize

import (
	"fmt"
	"strings"
)

func ParseColor(flag string) string {
	ColorMap := map[string]string{ // This is the color map that maps the ansicodes
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"white":   "\033[37m",
	}

	find := strings.HasPrefix(flag, "--color=")

	if find != true {
		fmt.Println("Usage: go run . [OPTION] [STRING] EX: go run . --color=<color> <substring to be colored> \"something\"")
		return ""
	}

	parts := strings.Split(flag, "=") //Splits the flag based on the delimiter "="

	if parts[0] != "--color" { // Checks if the first part is not equal to the flag and returns Usage.
		fmt.Println("Usage: go run . [OPTION] [STRING] EX: go run . --color=<color> <substring to be colored> \"something\"")
		return ""
	}

	ansicode := ColorMap[parts[1]] // This stores the color requested i.e 'red', 'blue', e.t.c

	if ansicode == "" {
		fmt.Println("Usage: go run . [OPTION] [STRING] EX: go run . --color=<color> <substring to be colored> \"something\"")
		return ""
	}

	return ansicode

}
