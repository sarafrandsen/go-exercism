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
	var reply string
	remark = strings.TrimSpace(remark)

	switch {
	case isShouting(remark) && strings.HasSuffix(remark, "?"):
		reply = "Calm down, I know what I'm doing!"
	case strings.HasSuffix(remark, "?"):
		reply = "Sure."
	case isShouting(remark):
		reply = "Whoa, chill out!"
	case remark == "":
		reply = "Fine. Be that way!"
	default:
		reply = "Whatever."
	}
	return reply
}
