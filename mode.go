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
func MaskSetMode(ma ModeMask, mo Mode) ModeMask { return ModeMask(uint8(ma) | uint8(mo)) }

// MaskClearMode clears a specific Mode from a given Mode bitmask
func MaskClearMode(ma ModeMask, mo Mode) ModeMask { return ModeMask(uint8(ma) &^ uint8(mo)) }

// MaskToggleMode toggles a specific Mode in a given Mode bitmask
func MaskToggleMode(ma ModeMask, mo Mode) ModeMask { return ModeMask(uint8(ma) ^ uint8(mo)) }

// MaskHasMode returns true if a given Mode bitmask holds a specific Mode
func MaskHasMode(ma ModeMask, mo Mode) bool { return uint8(ma)&uint8(mo) != 0 }

func ModesFromFlags(ms string) ModeMask {
	cl := strings.Split(ms, "")
	var mm ModeMask
	for _, m := range cl {
		switch m {
		case "C":
			mm = MaskSetMode(mm, ModeLowerCase|ModeNumeric|ModeSpecial|ModeUpperCase)
		case "h":
			mm = MaskClearMode(mm, ModeHumanReadable)
		case "H":
			mm = MaskSetMode(mm, ModeHumanReadable)
		case "l":
			mm = MaskClearMode(mm, ModeLowerCase)
		case "L":
			mm = MaskSetMode(mm, ModeLowerCase)
		case "n":
			mm = MaskClearMode(mm, ModeNumeric)
		case "N":
			mm = MaskSetMode(mm, ModeNumeric)
		case "s":
			mm = MaskClearMode(mm, ModeSpecial)
		case "S":
			mm = MaskSetMode(mm, ModeSpecial)
		case "u":
			mm = MaskClearMode(mm, ModeUpperCase)
		case "U":
			mm = MaskSetMode(mm, ModeUpperCase)
		}
	}

	return mm

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
