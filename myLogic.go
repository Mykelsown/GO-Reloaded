package GoReloaded

func MyLogic(text string) string {
	final := HexAndBin(text)
	final = LowUpAndCap(text)
	return final
}