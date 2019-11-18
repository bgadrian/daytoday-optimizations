package stringsconcat

import (
	"strings"
)

func concatenation(list ...string) string {
	var result string
	for i := range list {
		result = result + list[i]
	}
	return result
}

func builder(list ...string) string {
	var result strings.Builder
	for i := range list {
		result.WriteString(list[i])
	}
	return result.String()
}

func builderWithPrealloc(list ...string) string {
	var result strings.Builder
	magicNumberGuessAvgItemSize := len(list[0])
	result.Grow(magicNumberGuessAvgItemSize * len(list))
	for i := range list {
		result.WriteString(list[i])
	}
	return result.String()
}
