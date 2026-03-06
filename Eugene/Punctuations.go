package function

import (
	"regexp"
	"strings"
)

func PunctuationsHandler(s string) string {

	reSpaceBefore := regexp.MustCompile(`\s+([.,!?;:]+)`)
	s = reSpaceBefore.ReplaceAllString(s, "$1")

	reSpaceInside := regexp.MustCompile(`([.,!?;:])\s+([.,!?;:])`)
	for reSpaceInside.MatchString(s) {
		s = reSpaceInside.ReplaceAllString(s, "$1$2")
	}

	reMissingAfter := regexp.MustCompile(`([.,!?;:]+)([A-Za-z])`)
	s = reMissingAfter.ReplaceAllString(s, "$1 $2")

	// remove space after opening quote
	reQuoteStart := regexp.MustCompile(`'\s+`)
	s = reQuoteStart.ReplaceAllString(s, "'")

	// remove space before closing quote
	reQuoteEnd := regexp.MustCompile(`\s+'`)
	s = reQuoteEnd.ReplaceAllString(s, "'")

	var finRes strings.Builder
	for i := 0; i < len(s)-1; i++ {
		if i == 0 && s[i] == ' ' || i == len(s)-i && s[len(s)-i] == ' '  {
			continue
		}
		finRes.WriteString(string())
	}

	return finRes.String()
}
