package apg

import (
	"strings"
)

// Mode represents a mode of characters
type Mode uint8

// ModeMask represents a bitmask of character modes
type ModeMask uint8

const (
	// ModeNumeric sets the bitmask to include numeric in the generated passwords
	ModeNumeric = 1 << iota
	// ModeLowerCase sets the bitmask to include lower case characters in the
	// generated passwords
	ModeLowerCase
	// ModeUpperCase sets the bitmask to include upper case characters in the
	// generated passwords
	ModeUpperCase
	// ModeSpecial sets the bitmask to include special characters in the
	// generated passwords
	ModeSpecial
	// ModeHumanReadable sets the bitmask to generate human readable passwords
	ModeHumanReadable
)

const (
	// CharRangeAlphaLower represents all lower-case alphabetical characters
	CharRangeAlphaLower = "abcdefghijklmnopqrstuvwxyz"
	// CharRangeAlphaLowerHuman represents the human-readable lower-case alphabetical characters
	CharRangeAlphaLowerHuman = "abcdefghjkmnpqrstuvwxyz"
	// CharRangeAlphaUpper represents all upper-case alphabetical characters
	CharRangeAlphaUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// CharRangeAlphaUpperHuman represents the human-readable upper-case alphabetical characters
	CharRangeAlphaUpperHuman = "ABCDEFGHJKMNPQRSTUVWXYZ"
	// CharRangeNumeric represents all numerical characters
	CharRangeNumeric = "1234567890"
	// CharRangeNumericHuman represents all human-readable numerical characters
	CharRangeNumericHuman = "23456789"
	// CharRangeSpecial represents all special characters
	CharRangeSpecial = `!\"#$%&'()*+,-./:;<=>?@[\\]^_{|}~`
	// CharRangeSpecialHuman represents all human-readable special characters
	CharRangeSpecialHuman = `#%*+-:;=`
)

// MaskSetMode sets a specific Mode to a given Mode bitmask
func MaskSetMode(mask ModeMask, mode Mode) ModeMask { return ModeMask(uint8(mask) | uint8(mode)) }

// MaskClearMode clears a specific Mode from a given Mode bitmask
func MaskClearMode(mask ModeMask, mode Mode) ModeMask { return ModeMask(uint8(mask) &^ uint8(mode)) }

// MaskToggleMode toggles a specific Mode in a given Mode bitmask
func MaskToggleMode(mask ModeMask, mode Mode) ModeMask { return ModeMask(uint8(mask) ^ uint8(mode)) }

// MaskHasMode returns true if a given Mode bitmask holds a specific Mode
func MaskHasMode(mask ModeMask, mode Mode) bool { return uint8(mask)&uint8(mode) != 0 }

func ModesFromFlags(maskString string) ModeMask {
	cl := strings.Split(maskString, "")
	var modeMask ModeMask
	for _, m := range cl {
		switch m {
		case "C":
			modeMask = MaskSetMode(modeMask, ModeLowerCase|ModeNumeric|ModeSpecial|ModeUpperCase)
		case "h":
			modeMask = MaskClearMode(modeMask, ModeHumanReadable)
		case "H":
			modeMask = MaskSetMode(modeMask, ModeHumanReadable)
		case "l":
			modeMask = MaskClearMode(modeMask, ModeLowerCase)
		case "L":
			modeMask = MaskSetMode(modeMask, ModeLowerCase)
		case "n":
			modeMask = MaskClearMode(modeMask, ModeNumeric)
		case "N":
			modeMask = MaskSetMode(modeMask, ModeNumeric)
		case "s":
			modeMask = MaskClearMode(modeMask, ModeSpecial)
		case "S":
			modeMask = MaskSetMode(modeMask, ModeSpecial)
		case "u":
			modeMask = MaskClearMode(modeMask, ModeUpperCase)
		case "U":
			modeMask = MaskSetMode(modeMask, ModeUpperCase)
		}
	}

	return modeMask
}

// String satisfies the fmt.Stringer interface for the Mode type
func (m Mode) String() string {
	switch m {
	case ModeHumanReadable:
		return "Human-readable"
	case ModeLowerCase:
		return "Lower-case"
	case ModeNumeric:
		return "Numeric"
	case ModeSpecial:
		return "Special"
	case ModeUpperCase:
		return "Upper-case"
	default:
		return "Unknown"
	}
}
