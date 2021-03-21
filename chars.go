package main

import "regexp"

// Provide the range of available characters based on provided parameters
func getCharRange() string {
	pwUpperChars := PwUpperChars
	pwLowerChars := PwLowerChars
	pwNumbers := PwNumbers
	pwSpecialChars := PwSpecialChars
	if config.humanReadable {
		pwUpperChars = PwUpperCharsHuman
		pwLowerChars = PwLowerCharsHuman
		pwNumbers = PwNumbersHuman
		pwSpecialChars = PwSpecialCharsHuman
	}

	var charRange string
	if config.useLowerCase {
		charRange = charRange + pwLowerChars
	}
	if config.useUpperCase {
		charRange = charRange + pwUpperChars
	}
	if config.useNumber {
		charRange = charRange + pwNumbers
	}
	if config.useSpecial {
		charRange = charRange + pwSpecialChars
	}
	if config.excludeChars != "" {
		regExp := regexp.MustCompile("[" + config.excludeChars + "]")
		charRange = regExp.ReplaceAllLiteralString(charRange, "")
	}

	return charRange
}
