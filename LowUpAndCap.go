package GoReloaded

import (
	"regexp"
	"strconv"
	"strings"
)

func LowUpAndCap(s string) string {
	re := regexp.MustCompile(`\((up|low|cap)(?:,\s*(\d+))?\)`)

	for {
		loc := re.FindStringSubmatchIndex(s)
		if loc == nil {
			break
		}

		sub := re.FindStringSubmatch(s)

		mode := sub[1]
		n := 1

		if sub[2] != "" {
			val, _ := strconv.Atoi(sub[2])
			n = val
		}

		before := s[:loc[0]]
		after := s[loc[1]:]

		words := strings.Fields(before)

		for i := len(words) - n; i < len(words); i++ {
			if i < 0 {
				continue
			}

			switch mode {
			case "up":
				words[i] = strings.ToUpper(words[i])

			case "low":
				words[i] = strings.ToLower(words[i])

			case "cap":
				w := strings.ToLower(words[i])
				if len(w) > 0 {
					words[i] = strings.ToUpper(string(w[0])) + w[1:]
				}
			}
		}

		before = strings.Join(words, " ")
		s = before + after
	}

	return s
}
