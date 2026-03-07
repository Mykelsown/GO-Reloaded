package function

import (
	"regexp"
	"strconv"
)

func HexAndBin(s string) string {
	re := regexp.MustCompile(`([0-9A-Fa-f]+)\s*\((hex|bin)\)`) // compiles the pattern to check for

	result := re.ReplaceAllStringFunc(s, func(match string) string { // Replace for every pattern match

		sub := re.FindStringSubmatch(match) // isolates the value and the transormation denoter in a slice of strings

		value := sub[1] // value to e converted e.g "1E"
		mode := sub[2] // Denotor e.g "(hex)"

		var base int

		switch mode {
		case "hex":
			base = 16
		case "bin":
			base = 2
		}

		// Conversion of the value to an integer
		num, err := strconv.ParseInt(value, base, 64) 
		// Returns the string as it is if there is an error
		if err != nil {
			return match
		}

		// Converts to base 10 of num
		return strconv.FormatInt(num, 10)
	})

	return result
}
