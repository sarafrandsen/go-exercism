package bob

import (
	"strings"
	"unicode"
)

// hasLetters checks if there are any alpha characters
func hasLetters(remark string) bool {
	for _, r := range remark {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

// isShouting checks if all alpha characters are uppercase
func isShouting(remark string) bool {
	if remark == strings.ToUpper(remark) && hasLetters(remark) {
		return true
	}
	return false
}

// Hey returns the response Bob would give based on a given remark
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	switch {
	case isShouting(remark) && strings.HasSuffix(remark, "?"):
		return "Calm down, I know what I'm doing!"
	case strings.HasSuffix(remark, "?"):
		return "Sure."
	case isShouting(remark):
		return "Whoa, chill out!"
	case remark == "":
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
