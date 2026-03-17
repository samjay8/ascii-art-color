package colorize

import (
	"bufio"
	"os"
	"testing"
)

func TestAsciiColor(t *testing.T) {
	type testCase struct {
		name      string
		input     string
		substring string
		color     string
		expected  string
	}

	file, err := os.Open("../standard.txt")
	if err != nil {
		t.Fatal("Error opening file", err)
	}
	defer file.Close()

	var bannerlines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bannerlines = append(bannerlines, scanner.Text())
	}

	tests := []testCase{
		{"no color", "Hello", "", "", " _    _          _   _          \n| |  | |        | | | |         \n| |__| |   ___  | | | |   ___   \n|  __  |  / _ \\ | | | |  / _ \\  \n| |  | | |  __/ | | | | | (_) | \n|_|  |_|  \\___| |_| |_|  \\___/  \n                                \n                                \n"},
		{"full color", "Hi", "", "\033[31m", "\x1b[31m _    _  \x1b[0m\x1b[31m _  \x1b[0m\n\x1b[31m| |  | | \x1b[0m\x1b[31m(_) \x1b[0m\n\x1b[31m| |__| | \x1b[0m\x1b[31m _  \x1b[0m\n\x1b[31m|  __  | \x1b[0m\x1b[31m| | \x1b[0m\n\x1b[31m| |  | | \x1b[0m\x1b[31m| | \x1b[0m\n\x1b[31m|_|  |_| \x1b[0m\x1b[31m|_| \x1b[0m\n\x1b[31m         \x1b[0m\x1b[31m    \x1b[0m\n\x1b[31m         \x1b[0m\x1b[31m    \x1b[0m\n"},
		{"empty input", "", "", "", ""},
		{"newline input", `\n`, "", "", "\n"},
		{"substring color", "Hello", "He", "\033[32m", "\x1b[32m _    _  \x1b[0m\x1b[32m       \x1b[0m _   _          \n\x1b[32m| |  | | \x1b[0m\x1b[32m       \x1b[0m| | | |         \n\x1b[32m| |__| | \x1b[0m\x1b[32m  ___  \x1b[0m| | | |   ___   \n\x1b[32m|  __  | \x1b[0m\x1b[32m / _ \\ \x1b[0m| | | |  / _ \\  \n\x1b[32m| |  | | \x1b[0m\x1b[32m|  __/ \x1b[0m| | | | | (_) | \n\x1b[32m|_|  |_| \x1b[0m\x1b[32m \\___| \x1b[0m|_| |_|  \\___/  \n\x1b[32m         \x1b[0m\x1b[32m       \x1b[0m                \n\x1b[32m         \x1b[0m\x1b[32m       \x1b[0m                \n"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			positions := ColoredPositions(tc.input, tc.substring)
			got := AsciiArt(tc.input, bannerlines, tc.color, positions)
			if got != tc.expected {
				t.Errorf("AsciiArt(%q) = %q, expected = %q", tc.input, got, tc.expected)
			}
		})
	}
}
