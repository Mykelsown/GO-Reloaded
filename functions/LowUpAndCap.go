package function

import (
	"regexp"
	"strconv"
	"strings"
)

func LowUpAndCap(s string) string {
	// Match tags like (up), (low), (cap), (up, 3), (low, 2), (cap, 5)
	// Group 1 captures the mode (up/low/cap), group 2 optionally captures the number N
	re := regexp.MustCompile(`\((up|low|cap)(?:,\s*(\d+))?\)`)

	for {
		// Find the position of the next tag in the string
		loc := re.FindStringSubmatchIndex(s)
		if loc == nil {
			break // No more tags found, we're done
		}

		// Extract the tag content: mode and optional N
		sub := re.FindStringSubmatch(s)

		mode := sub[1] // "up", "low", or "cap"
		n := 1         // default: affect only 1 word

		// If a number was provided in the tag, use it as N
		if sub[2] != "" {
			val, _ := strconv.Atoi(sub[2])
			n = val
		}

		// Split the string around the tag
		before := s[:loc[0]] // everything before the tag
		after := s[loc[1]:]  // everything after the tag

		// Split the preceding text into individual words
		words := strings.Fields(before)

		// Apply the transformation to the last N words before the tag
		for i := len(words) - n; i < len(words); i++ {
			if i < 0 {
				continue // skip if N exceeds the number of available words
			}

			switch mode {
			case "up":
				words[i] = strings.ToUpper(words[i])

			case "low":
				words[i] = strings.ToLower(words[i])

			case "cap":
				// Lowercase the whole word first, then capitalize the first letter
				w := strings.ToLower(words[i])
				if len(w) > 0 {
					words[i] = strings.ToUpper(string(w[0])) + w[1:]
				}
			}
		}

		// Reassemble the string with the transformed words, dropping the tag
		before = strings.Join(words, " ")
		s = before + after
	}

	return s
}