package GoReloaded

func MyLogic(text string) string {
	final := HexAndBin(text)
	final = PunctuationsHandler(final)
	final = LowUpAndCap(final)
	final = ArticleHandler(final)
	return final
}