package spelling

import (
	"fmt"
	"strings"
)

var (
	symbNumNames = map[byte]string{
		'1': "ONE",
		'2': "TWO",
		'3': "THREE",
		'4': "FOUR",
		'5': "FIVE",
		'6': "SIX",
		'7': "SEVEN",
		'8': "EIGHT",
		'9': "NINE",
		'0': "ZERO",
		33:  "EXCLAMATION_POINT",
		34:  "QUOTATION_MARK",
		35:  "CROSSHATCH",
		36:  "DOLLAR_SIGN",
		37:  "PERCENT_SIGN",
		38:  "AMPERSAND",
		39:  "APOSTROPHE",
		40:  "LEFT_PARENTHESIS",
		41:  "RIGHT_PARENTHESIS",
		42:  "ASTERISK",
		43:  "PLUS_SIGN",
		44:  "COMMA",
		45:  "HYPHEN",
		46:  "PERIOD",
		47:  "SLASH",
		58:  "COLON",
		59:  "SEMICOLON",
		60:  "LESS_THAN",
		61:  "EQUAL_SIGN",
		62:  "GREATER_THAN",
		63:  "QUESTION_MARK",
		64:  "AT_SIGN",
		91:  "LEFT_BRACKET",
		92:  "BACKSLASH",
		93:  "RIGHT_BRACKET",
		94:  "CIRCUMFLEX",
		95:  "UNDERSCORE",
		96:  "GRAVE",
		123: "LEFT_BRACE",
		124: "VERTICAL_BAR",
		125: "RIGHT_BRACE",
		126: "TILDE",
	}
	alphabetNames = map[byte]string{
		'A': "Alfa",
		'B': "Bravo",
		'C': "Charlie",
		'D': "Delta",
		'E': "Echo",
		'F': "Foxtrot",
		'G': "Golf",
		'H': "Hotel",
		'I': "India",
		'J': "Juliett",
		'K': "Kilo",
		'L': "Lima",
		'M': "Mike",
		'N': "November",
		'O': "Oscar",
		'P': "Papa",
		'Q': "Quebec",
		'R': "Romeo",
		'S': "Sierra",
		'T': "Tango",
		'U': "Uniform",
		'V': "Victor",
		'W': "Whiskey",
		'X': "X_ray",
		'Y': "Yankee",
		'Z': "Zulu",
	}
)

// String returns an english spelled version of the given string
func String(pwString string) (string, error) {
	var returnString []string
	for _, curChar := range pwString {
		curSpellString, err := ConvertCharToName(byte(curChar))
		if err != nil {
			return "", err
		}
		returnString = append(returnString, curSpellString)
	}
	return strings.Join(returnString, "/"), nil
}

// ConvertCharToName converts a given ascii byte into the corresponding english spelled
// name
func ConvertCharToName(charByte byte) (string, error) {
	var returnString string
	if charByte > 64 && charByte < 91 {
		returnString = alphabetNames[charByte]
	} else if charByte > 96 && charByte < 123 {
		returnString = strings.ToLower(alphabetNames[charByte-32])
	} else {
		returnString = symbNumNames[charByte]
	}

	if returnString == "" {
		err := fmt.Errorf("cannot convert to character to name: %q is an unknown character", charByte)
		return "", err
	}
	return returnString, nil
}
