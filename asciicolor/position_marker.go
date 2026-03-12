package colorize

import "strings"

func ColoredPositions(input, substring string) []bool {

	boolslice := make([]bool, len(input))

	if substring == "" {
		for i := range boolslice {
			boolslice[i] = true
		}
		return boolslice

	} else {
		for pos := 0; pos <= len(input); pos++ {
			find := strings.Index(input[pos:], substring)
			if find == -1 {
				break
			}
			for i := pos + find; i < pos+find+len(substring); i++ {
				boolslice[i] = true
			}
			pos = pos + find + len(substring)
		}

	}
	return boolslice
}
