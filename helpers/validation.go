package helpers

import (
	"unicode"
)

func IsPascalCase(s string) bool {
	if s == "" {
		return false
	}

	runes := []rune(s)

	if !unicode.IsUpper(runes[0]) {
		return false
	}

	for _, r := range runes {
		if !(unicode.IsLetter(r) || unicode.IsDigit(r)) {
			return false
		}
	}

	hasLower := false
	for _, r := range runes {
		if unicode.IsLower(r) {
			hasLower = true
			break
		}
	}

	return hasLower
}
