// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

func hasLetters(remark string) bool {
	for _, r := range remark {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func isShouting(remark string) bool {
	if remark == strings.ToUpper(remark) && hasLetters(remark) {
		return true
	}
	return false
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// if string.all == caps && string.last == ?
	// if string.all == caps || string.last == !
	// if string.last == ?
	// if string == ""
	// else

	// refactor to use case

	var reply string
	remark = strings.TrimSpace(remark)
	if isShouting(remark) && strings.HasSuffix(remark, "?") {
		reply = "Calm down, I know what I'm doing!"
	} else if isShouting(remark) {
		reply = "Whoa, chill out!"
	} else if strings.HasSuffix(remark, "?") {
		reply = "Sure."
	} else if remark == "" {
		reply = "Fine. Be that way!"
	} else {
		reply = "Whatever."
	}
	return reply
}
