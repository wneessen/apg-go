package chars

import (
	"github.com/wneessen/apg-go/config"
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
