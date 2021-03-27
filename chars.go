package main

import (
	"regexp"
)

const PwLowerCharsHuman string = "abcdefghjkmnpqrstuvwxyz"
const PwUpperCharsHuman string = "ABCDEFGHJKMNPQRSTUVWXYZ"
const PwLowerChars string = "abcdefghijklmnopqrstuvwxyz"
const PwUpperChars string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const PwSpecialCharsHuman string = "\"#%*+-/:;=\\_|~"
const PwSpecialChars string = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
const PwNumbersHuman string = "23456789"
const PwNumbers string = "1234567890"

// Provide the range of available characters based on provided parameters
func getCharRange(config *Config) string {
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
		regExp := regexp.MustCompile("[" + regexp.QuoteMeta(config.excludeChars) + "]")
		charRange = regExp.ReplaceAllLiteralString(charRange, "")
	}

	return charRange
}
