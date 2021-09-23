package chars

import (
	"github.com/wneessen/apg-go/config"
	"regexp"
)

// PwLowerCharsHuman is the range of lower-case characters in human-readable mode
const PwLowerCharsHuman string = "abcdefghjkmnpqrstuvwxyz"

// PwUpperCharsHuman is the range of upper-case characters in human-readable mode
const PwUpperCharsHuman string = "ABCDEFGHJKMNPQRSTUVWXYZ"

// PwLowerChars is the range of lower-case characters
const PwLowerChars string = "abcdefghijklmnopqrstuvwxyz"

// PwUpperChars is the range of upper-case characters
const PwUpperChars string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// PwSpecialCharsHuman is the range of special characters in human-readable mode
const PwSpecialCharsHuman string = "\"#%*+-/:;=\\_|~"

// PwSpecialChars is the range of special characters
const PwSpecialChars string = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

// PwNumbersHuman is the range of numbers in human-readable mode
const PwNumbersHuman string = "23456789"

// PwNumbers is the range of numbers
const PwNumbers string = "1234567890"

// GetRange provides the range of available characters based on configured parameters
func GetRange(config *config.Config) string {
	pwUpperChars := PwUpperChars
	pwLowerChars := PwLowerChars
	pwNumbers := PwNumbers
	pwSpecialChars := PwSpecialChars
	if config.HumanReadable {
		pwUpperChars = PwUpperCharsHuman
		pwLowerChars = PwLowerCharsHuman
		pwNumbers = PwNumbersHuman
		pwSpecialChars = PwSpecialCharsHuman
	}

	var charRange string
	if config.UseLowerCase {
		charRange = charRange + pwLowerChars
	}
	if config.UseUpperCase {
		charRange = charRange + pwUpperChars
	}
	if config.UseNumber {
		charRange = charRange + pwNumbers
	}
	if config.UseSpecial {
		charRange = charRange + pwSpecialChars
	}
	if config.ExcludeChars != "" {
		regExp := regexp.MustCompile("[" + regexp.QuoteMeta(config.ExcludeChars) + "]")
		charRange = regExp.ReplaceAllLiteralString(charRange, "")
	}

	return charRange
}
