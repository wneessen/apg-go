package apg

// List of default values for Config instances
const (
	// DefaultMinLength reflects the default minimum length of a generated password
	DefaultMinLength int64 = 12
	// DefaultMaxLength reflects the default maximum length of a generated password
	DefaultMaxLength int64 = 20
	// DefaultMode sets the default character set mode bitmask to a combination of
	// lower- and upper-case characters as well as numbers
	DefaultMode ModeMask = ModeLowerCase | ModeNumeric | ModeUpperCase
	// DefaultNumberPass reflects the default amount of passwords returned by the generator
	DefaultNumberPass int64 = 6
)

// Config represents the apg.Generator config parameters
type Config struct {
	// Algo
	Algorithm Algorithm
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
	// Mode holds the different character modes for the Random algorithm
	Mode ModeMask
	// NumberPass sets the number of passwords that are generated
	// and returned by the generator
	NumberPass int64
	// SpellPassword if set will spell the generated passwords in the phonetic alphabet
	SpellPassword bool
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

// WithMinLength overrides the minimum password length
func WithMinLength(length int64) Option {
	return func(config *Config) {
		config.MinLength = length
	}
}

// WithMaxLength overrides the maximum password length
func WithMaxLength(length int64) Option {
	return func(config *Config) {
		config.MaxLength = length
	}
}

// WithNumberPass overrides the amount of generated passwords setting
func WithNumberPass(amount int64) Option {
	return func(config *Config) {
		config.NumberPass = amount
	}
}
