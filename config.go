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
	MinLength    int64
	MinLowerCase int64
	MinNumeric   int64
	MinSpecial   int64
	MinUpperCase int64
	// Mode holds the different character modes for the Random algorithm
	Mode ModeMask
	// NumberPass sets the number of passwords that are generated
	// and returned by the generator
	NumberPass int64
}

// Option is a function that can override default Config settings
type Option func(*Config)

// NewConfig creates a new Config instance and pre-fills it with sane
// default settings. The Config is returned as pointer value
func NewConfig(o ...Option) *Config {
	c := &Config{
		MaxLength:  DefaultMaxLength,
		MinLength:  DefaultMinLength,
		Mode:       DefaultMode,
		NumberPass: DefaultNumberPass,
	}

	// Override defaults with optionally provided config.Option functions
	for _, co := range o {
		if co == nil {
			continue
		}
		co(c)
	}
	return c
}

// WithMinLength overrides the minimum password length
func WithMinLength(l int64) Option {
	return func(c *Config) {
		c.MinLength = l
	}
}

// WithMaxLength overrides the maximum password length
func WithMaxLength(l int64) Option {
	return func(c *Config) {
		c.MaxLength = l
	}
}

// WithNumberPass overrides the number of generated passwords setting
func WithNumberPass(n int64) Option {
	return func(c *Config) {
		c.NumberPass = n
	}
}
