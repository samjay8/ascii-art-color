package main

import (
	"bufio"
	colorize "colorize/asciicolor"
	"fmt"
	"os"
)

func main() {
	var input, substring string
	var color string

	if len(os.Args) == 2 {
		input = os.Args[1]

	} else if len(os.Args) == 3 {
		color = colorize.ParseColor(os.Args[1])
		input = os.Args[2]
	} else if len(os.Args) == 4 {
		color = colorize.ParseColor(os.Args[1])
		substring = os.Args[2]
		input = os.Args[3]
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING] EX: go run . --color=<color> <substring to be colored> \"something\"")
	}

	file, err := os.Open("standard.txt")

	if err != nil {
		fmt.Println("Error Opening file", err)
	}
	defer file.Close()

	var bannerlines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		bannerlines = append(bannerlines, scanner.Text())
	}
	fmt.Println(input, substring, color)
}
