package tests

import (
	function "GoReloaded/Eugene"
	"testing"
)


func TestConvertToVowel(t *testing.T) {
	samples := []struct {
		input    string
		expected string
	}{
		{
			"There it was. A amazing rock!",
			"There it was. An amazing rock!",
		},
	}

	for _, sentences := range samples {
		result := function.ArticleHandler(sentences.input)

		if result != sentences.expected {
			t.Errorf("Expected: %v \n Got: %v", sentences.expected, result)
		}
	}
}

func TestHexAndBin(t *testing.T) {
	samples := []struct {
		input    string
		expected string
	}{
		{
			"1E (hex) files were added",
			"30 files were added",
		},
		{
			"It has been 10 (bin) years", 
			"It has been 2 years",
		},
	}

	for _, sentences := range samples {
		result := function.HexAndBin(sentences.input)

		if result != sentences.expected {
			t.Errorf("Expected: %v \n Got: %v", sentences.expected, result)
		}
	}
}


func TestLowUpAndCap(t *testing.T) {
	samples := []struct {
		input    string
		expected string
	}{
		{
			"This is so exciting (up, 2)",
			"This is SO EXCITING",
		},
		{
			"Ready, set, go (up) !", 
			"Ready, set, GO !",
		},
	}

	for _, sentences := range samples {
		result := function.LowUpAndCap(sentences.input)

		if result != sentences.expected {
			t.Errorf("Expected: %v \n Got: %v", sentences.expected, result)
		}
	}
}

func TestPunctuations(t *testing.T) {
	samples := []struct {
		input    string
		expected string
	}{
		{
			"This is so EXCITING .  ...yes it is . ",
			"This is so EXCITING.... yes it is.",
		},
		{
			"Ready , set , go  ! ", 
			"Ready, set, go!",
		},
	}

	for _, sentences := range samples {
		result := function.PunctuationsHandler(sentences.input)

		if result != sentences.expected {
			t.Errorf("\nExpected: %v \n Got: %v", sentences.expected, result)
		}
	}
}
