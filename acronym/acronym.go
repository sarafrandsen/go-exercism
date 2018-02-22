package acronym

import (
	"strings"
	"unicode"
)

// isSpecialChar takes c with type byte and returns true if it is a special character and false if it is a letter
func isSpecialChar(char byte) bool {
	return !unicode.IsLetter(rune(char))
}

// Abbreviate will return a string a containing only the first letter of each word in the given string s.
func Abbreviate(s string) string {
	acronym := ""
	const Space byte = 32
	for i, char := range s {
		if i == 0 || (i != 0 && isSpecialChar(s[i-1]) && s[i] != Space) {
			acronym += string(char)
		}
	}
	return strings.ToUpper(acronym)
}
