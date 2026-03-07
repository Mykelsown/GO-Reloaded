package function

func MyLogic(text string) string {
	// This calls al the function on the unrefined text passed in from the file handler, nd then returns a refined text
	final := HexAndBin(text)
	final = PunctuationsHandler(final)
	final = LowUpAndCap(final)
	final = ArticleHandler(final)
	return final
}