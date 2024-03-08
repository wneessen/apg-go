package apg

// VERSION represents the version string
const VERSION = "2.0.0"

// Generator is the password generator type of the APG package
type Generator struct {
	// config is a pointer to the apg config instance
	config *Config
}

// New returns a new password Generator type
func New(config *Config) *Generator {
	return &Generator{
		config: config,
	}
}
