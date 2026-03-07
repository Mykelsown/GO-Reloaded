package function

import "regexp"

func ArticleHandler(s string) string {
	// Matches "a" or "A" followed by a word starting with vowel or h
	re := regexp.MustCompile(`\b([aA])\s+([aeiouhAEIOUH]\w*)`)

	return re.ReplaceAllStringFunc(s, func(match string) string { // this replaces each and every match found in the string with the compiled regex 
		// This find the mathes and store them in a slice of string; it find the article, the word after the article also. Where the index zero stores the article and the next word after it.
		sub := re.FindStringSubmatch(match)

		article := sub[1] // stroes the article
		word := sub[2] // stores the word after the article

		// This doea the conversion of "A" to "An" i.e the capitalization
		if article == "A" {
			return "An " + word
		}

		// return a lowercase "an" instead
		return "an " + word
	})
}