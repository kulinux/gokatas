// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	trimmed := strings.Trim(remark, " \t\r\n")
	yell := strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark
	question := len(trimmed) > 0 && trimmed[len(trimmed)-1] == '?'

	if len(trimmed) == 0 {
		return "Fine. Be that way!"
	}
	if yell && question {
		return "Calm down, I know what I'm doing!"
	}
	if yell {
		return "Whoa, chill out!"
	}
	if question {
		return "Sure."
	}
	return "Whatever."
}
