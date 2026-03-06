package function

import (
	"regexp"
	"strconv"
)

func HexAndBin(s string) string {
	re := regexp.MustCompile(`([0-9A-Fa-f]+)\s*\((hex|bin)\)`)

	result := re.ReplaceAllStringFunc(s, func(match string) string {

		sub := re.FindStringSubmatch(match)

		value := sub[1]
		mode := sub[2]

		var base int

		switch mode {
		case "hex":
			base = 16
		case "bin":
			base = 2
		}

		num, err := strconv.ParseInt(value, base, 64)
		if err != nil {
			return match
		}

		return strconv.FormatInt(num, 10)
	})

	return result
}
