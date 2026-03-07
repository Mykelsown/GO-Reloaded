package function

import (
	"regexp"
	"strings"
)

func PunctuationsHandler(s string) string {

	// Match one or more whitespace characters immediately before any punctuation
	// mark(s), and remove the space, keeping only the punctuation — e.g. "Hello !" → "Hello!"
	reSpaceBefore := regexp.MustCompile(`\s+([.,!?;:]+)`)
	s = reSpaceBefore.ReplaceAllString(s, "$1")

	// Match a punctuation mark, whitespace, then another punctuation mark.
	// Loop until no more matches exist, since ReplaceAllString only finds
	// non-overlapping matches — e.g. ". . ." needs two passes to become "..."
	reSpaceInside := regexp.MustCompile(`([.,!?;:])\s+([.,!?;:])`)
	for reSpaceInside.MatchString(s) {
		s = reSpaceInside.ReplaceAllString(s, "$1$2")
	}

	// Match punctuation immediately followed by a letter (no space between them),
	// and inject a space — e.g. "Hello,world" → "Hello, world"
	reMissingAfter := regexp.MustCompile(`([.,!?;:]+)([A-Za-z])`)
	s = reMissingAfter.ReplaceAllString(s, "$1 $2")

	// Match a single quote followed by whitespace (opening quote with trailing space),
	// and remove the space — e.g. "' hello'" → "'hello'"
	reQuoteStart := regexp.MustCompile(`'\s+`)
	s = reQuoteStart.ReplaceAllString(s, "'")

	// Match whitespace immediately before a single quote (closing quote with leading space),
	// and remove the space — e.g. "'hello '" → "'hello'"
	reQuoteEnd := regexp.MustCompile(`\s+'`)
	s = reQuoteEnd.ReplaceAllString(s, "'")

	// Trim any remaining leading or trailing whitespace from the final result
	finRes := strings.TrimSpace(s)

	return finRes
}
