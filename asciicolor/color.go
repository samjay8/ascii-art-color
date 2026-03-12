package colorize

import (
	"fmt"
	"strings"
)

func ParseColor(flag string) string {
	ColorMap := map[string]string{
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"white":   "\033[37m",
	}

	parts := strings.Split(flag, "=")
	if parts[0] != "--color" {
		fmt.Println("Usage: go run . [OPTION] [STRING] EX: go run . --color=<color> <substring to be colored> \"something\"")
		return ""
	}

	ansicode := ColorMap[parts[1]]

	if ansicode == "" {
		fmt.Println("Usage: go run . [OPTION] [STRING] EX: go run . --color=<color> <substring to be colored> \"something\"")
		return ""
	}

	return ansicode

}
