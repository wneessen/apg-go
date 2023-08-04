package apg

// Mode represents a mode of characters
type Mode uint8

const (
	// ModeNumber sets the bitmask to include numbers in the generated passwords
	ModeNumber = 1 << iota
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
	// CharRangeNumber represents all numerical characters
	CharRangeNumber = "1234567890"
	// CharRangeNumberHuman represents all human-readable numerical characters
	CharRangeNumberHuman = "23456789"
	// CharRangeSpecial represents all special characters
	CharRangeSpecial = `!\"#$%&'()*+,-./:;<=>?@[\\]^_{|}~`
	// CharRangeSpecialHuman represents all human-readable special characters
	CharRangeSpecialHuman = `#%*+-:;=`
)

// SetMode sets a specific Mode to a given Mode bitmask
func SetMode(b, m Mode) Mode { return b | m }

// ClearMode clears a specific Mode from a given Mode bitmask
func ClearMode(b, m Mode) Mode { return b &^ m }

// ToggleMode toggles a specific Mode in a given Mode bitmask
func ToggleMode(b, m Mode) Mode { return b ^ m }

// HasMode returns true if a given Mode bitmask holds a specific Mode
func HasMode(b, m Mode) bool { return b&m != 0 }
