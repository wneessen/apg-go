// SPDX-FileCopyrightText: 2021-2024 Winni Neessen <wn@neessen.dev>
//
// SPDX-License-Identifier: MIT

package apg

// List of default values for Config instances
const (
	// DefaultMinLength reflects the default minimum length of a generated password
	DefaultMinLength int64 = 12
	// DefaultMaxLength reflects the default maximum length of a generated password
	DefaultMaxLength int64 = 20
	// DefaultBinarySize is the default byte size for generating binary random bytes
	DefaultBinarySize int64 = 32
	// DefaultMode sets the default character set mode bitmask to a combination of
	// lower- and upper-case characters as well as numbers
	DefaultMode ModeMask = ModeLowerCase | ModeNumeric | ModeUpperCase
	// DefaultNumberPass reflects the default amount of passwords returned by the generator
	DefaultNumberPass int64 = 6
)

// Config represents the apg.Generator config parameters
type Config struct {
	// Algorithm sets the Algorithm used for the password generation
	Algorithm Algorithm
	// BinaryHexMode if set will output the hex representation of the generated
	// binary random string
	BinaryHexMode bool
	// BinaryNewline if set will print out a new line in AlgoBinary mode
	BinaryNewline bool
	// CheckHIBP sets a flag if the generated password has to be checked
	// against the HIBP pwned password database
	CheckHIBP bool
	// ExcludeChars is a list of characters that should be excluded from
	// generated passwords
	ExcludeChars string
	// FixedLength sets a fixed length for generated passwords and ignores
	// the MinLength and MaxLength values
	FixedLength int64
	// MaxLength sets the maximum length for a generated password
	MaxLength int64
	// MinLength sets the minimum length for a generated password
	MinLength int64
	// MinLowerCase represents the minimum amount of lower-case characters that have
	// to be part of the generated password
	MinLowerCase int64
	// MinNumeric represents the minimum amount of numeric characters that have
	// to be part of the generated password
	MinNumeric int64
	// MinSpecial represents the minimum amount of special characters that have
	// to be part of the generated password
	MinSpecial int64
	// MinUpperCase represents the minimum amount of upper-case characters that have
	// to be part of the generated password
	MinUpperCase int64
	// MobileGrouping indicates if the generated password should be grouped in a
	// mobile-friendly manner
	MobileGrouping bool
	// Mode holds the different character modes for the Random algorithm
	Mode ModeMask
	// NumberPass sets the number of passwords that are generated
	// and returned by the generator
	NumberPass int64
	// SpellPassword if set will spell the generated passwords in the phonetic alphabet
	SpellPassword bool
	// SpellPronounceable if set will spell the generated pronounceable passwords in
	// as its corresponding syllables
	SpellPronounceable bool
}

// Option is a function that can override default Config settings
type Option func(*Config)

// NewConfig creates a new Config instance and pre-fills it with sane
// default settings. The Config is returned as pointer value
func NewConfig(opts ...Option) *Config {
	config := &Config{
		MaxLength:  DefaultMaxLength,
		MinLength:  DefaultMinLength,
		Mode:       DefaultMode,
		NumberPass: DefaultNumberPass,
	}

	// Override defaults with optionally provided config.Option functions
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		opt(config)
	}
	return config
}

// WithAlgorithm overrides the algorithm mode for the password generation
func WithAlgorithm(algo Algorithm) Option {
	return func(config *Config) {
		config.Algorithm = algo
	}
}

// WithBinaryHexMode sets the hex mode for the AlgoBinary
func WithBinaryHexMode() Option {
	return func(config *Config) {
		config.BinaryHexMode = true
	}
}

// WithExcludeChars sets a list of characters to be excluded in the generated
// passwords
func WithExcludeChars(chars string) Option {
	return func(config *Config) {
		config.ExcludeChars = chars
	}
}

// WithFixedLength sets a fixed password length
func WithFixedLength(length int64) Option {
	return func(config *Config) {
		config.FixedLength = length
	}
}

// WithMinLength overrides the minimum password length
func WithMinLength(length int64) Option {
	return func(config *Config) {
		config.MinLength = length
	}
}

// WithMinLowercase sets the minimum amount of lowercase characters that
// the generated password should contain
//
// CAVEAT: using to high values with this option, can lead to extraordinary
// calculation times, resulting in apg-go to never finish
func WithMinLowercase(amount int64) Option {
	return func(config *Config) {
		config.MinLowerCase = amount
	}
}

// WithMinNumeric sets the minimum amount of numeric characters that
// the generated password should contain
//
// CAVEAT: using to high values with this option, can lead to extraordinary
// calculation times, resulting in apg-go to never finish
func WithMinNumeric(amount int64) Option {
	return func(config *Config) {
		config.MinNumeric = amount
	}
}

// WithMinSpecial sets the minimum amount of special characters that
// the generated password should contain
//
// CAVEAT: using to high values with this option, can lead to extraordinary
// calculation times, resulting in apg-go to never finish
func WithMinSpecial(amount int64) Option {
	return func(config *Config) {
		config.MinSpecial = amount
	}
}

// WithMinUppercase sets the minimum amount of uppercase characters that
// the generated password should contain
//
// CAVEAT: using to high values with this option, can lead to extraordinary
// calculation times, resulting in apg-go to never finish
func WithMinUppercase(amount int64) Option {
	return func(config *Config) {
		config.MinUpperCase = amount
	}
}

// WithMaxLength overrides the maximum password length
func WithMaxLength(length int64) Option {
	return func(config *Config) {
		config.MaxLength = length
	}
}

// WithMobileGrouping enables the mobile-friendly character grouping for AlgoRandom
func WithMobileGrouping() Option {
	return func(config *Config) {
		config.MobileGrouping = true
	}
}

// WithModeMask overrides the default mode mask for the random algorithm
func WithModeMask(mask ModeMask) Option {
	return func(config *Config) {
		config.Mode = mask
	}
}
