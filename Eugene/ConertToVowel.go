package function

import "regexp"

func ArticleHandler(s string) string {
	// Matches "a" or "A" followed by a word starting with vowel or h
	re := regexp.MustCompile(`\b([aA])\s+([aeiouhAEIOUH]\w*)`)

	return re.ReplaceAllStringFunc(s, func(match string) string {
		sub := re.FindStringSubmatch(match)

		article := sub[1]
		word := sub[2]

		if article == "A" {
			return "An " + word
		}

		return "an " + word
	})
}