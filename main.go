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

	if len(os.Args) == 2 { // Checks for when the length of arguments in the terminal is equal 2
		input = os.Args[1] // Takes arguments [1] as input after go run . on the terminal

	} else if len(os.Args) == 3 { // Checks for when the length of arguments in the terminal is equal 3
		color = colorize.ParseColor(os.Args[1]) // Takes arguments[1] as the color
		input = os.Args[2]                      // Takes arguments[2] as the input
	} else if len(os.Args) == 4 { // Checks for when the length of arguments in the terminal is equal 4
		color = colorize.ParseColor(os.Args[1]) // Takes arguments[1] as the color
		substring = os.Args[2]                  // Takes arguments[2] as substring
		input = os.Args[3]                      // Takes arguments[3] as the input
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING] EX: go run . --color=<color> <substring to be colored> \"something\"") //Else, handle error
	}

	file, err := os.Open("standard.txt") // Opens the banner file

	if err != nil {
		fmt.Println("Error Opening file", err) // If err is not empty, handle the error
	}
	defer file.Close() // Close the file after banner file has been read

	var bannerlines []string

	scanner := bufio.NewScanner(file) // This scans the banner file opened

	for scanner.Scan() {
		bannerlines = append(bannerlines, scanner.Text()) // This iterates through, scans through the banner files opened, and stores a slice of strings of the result into bannerlines.
	}

	positions := colorize.ColoredPositions(input, substring) // This is the position marker that marks the particular input or string to be colored.

	result := colorize.AsciiArt(input, bannerlines, color, positions) // This return the graphical representation of the colored ascii-art

	fmt.Print(result) // This prints the result
}
