package config

import (
	"flag"
	"github.com/wneessen/apg-go/random"
	"log"
)

// Config is a struct that holds the different config parameters for the apg-go
// application
type Config struct {
	MinPassLen    int    // Minimum password length
	MaxPassLen    int    // Maximum password length
	NumOfPass     int    // Number of passwords to be generated
	UseComplex    bool   // Force complex password generation (implies all other Use* Options to be true)
	UseLowerCase  bool   // Allow lower-case chars in passwords
	UseUpperCase  bool   // Allow upper-case chars in password
	UseNumber     bool   // Allow numbers in passwords
	UseSpecial    bool   // Allow special chars in passwords
	HumanReadable bool   // Generated passwords use the "human readable" character set
	CheckHibp     bool   // Check generated are validated against the HIBP API for possible leaks
	ExcludeChars  string // List of characters to be excluded from the PW generation charset
	NewStyleModes string // Use the "new style" parameters instead of the single params
	SpellPassword bool   // Spell out passwords in the output
	ShowHelp      bool   // Display the help message in the CLI
	ShowVersion   bool   // Display the version string in the CLI
	OutputMode    int    // Interal parameter to control the output mode of the CLI
	PwAlgo        int    // PW generation algorithm to use (0: random PW based on flags, 1: pronouncable)
	SpellPron     bool   // Spell out the pronouncable password
}

// DefaultMinLength reflects the default minimum length of a generated password
const DefaultMinLength int = 12

// DefaultMaxLength reflects the default maximum length of a generated password
const DefaultMaxLength int = 20

// DefaultPwAlgo reflects the default password generation algorithm
const DefaultPwAlgo int = 1

// New parses the CLI flags and returns a new config object
func New() Config {
	var switchConf Config
	defaultSwitches := Config{
		UseLowerCase:  true,
		UseUpperCase:  true,
		UseNumber:     true,
		UseSpecial:    false,
		UseComplex:    false,
		HumanReadable: false,
	}
	config := Config{
		UseLowerCase:  defaultSwitches.UseLowerCase,
		UseUpperCase:  defaultSwitches.UseUpperCase,
		UseNumber:     defaultSwitches.UseNumber,
		UseSpecial:    defaultSwitches.UseSpecial,
		UseComplex:    defaultSwitches.UseComplex,
		HumanReadable: defaultSwitches.HumanReadable,
	}

	// Read and set all flags
	flag.BoolVar(&switchConf.UseLowerCase, "L", false, "Use lower case characters in passwords")
	flag.BoolVar(&switchConf.UseUpperCase, "U", false, "Use upper case characters in passwords")
	flag.BoolVar(&switchConf.UseNumber, "N", false, "Use numerich characters in passwords")
	flag.BoolVar(&switchConf.UseSpecial, "S", false, "Use special characters in passwords")
	flag.BoolVar(&switchConf.UseComplex, "C", false, "Generate complex passwords (implies -L -U -N -S, disables -H)")
	flag.BoolVar(&switchConf.HumanReadable, "H", false, "Generate human-readable passwords")
	flag.BoolVar(&config.SpellPassword, "l", false, "Spell generated password")
	flag.BoolVar(&config.CheckHibp, "p", false, "Check the HIBP database if the generated password was leaked before")
	flag.BoolVar(&config.ShowVersion, "v", false, "Show version")
	flag.IntVar(&config.MinPassLen, "m", DefaultMinLength, "Minimum password length")
	flag.IntVar(&config.MaxPassLen, "x", DefaultMaxLength, "Maxiumum password length")
	flag.IntVar(&config.NumOfPass, "n", 6, "Number of passwords to generate")
	flag.StringVar(&config.ExcludeChars, "E", "", "Exclude list of characters from generated password")
	flag.StringVar(&config.NewStyleModes, "M", "",
		"New style password parameters (higher priority than single parameters)")
	flag.IntVar(&config.PwAlgo, "a", DefaultPwAlgo, "Password generation algorithm")
	flag.BoolVar(&config.SpellPron, "t", false, "In pronouncable password mode, spell out the password")
	flag.Parse()

	// Invert-switch the defaults
	if switchConf.UseLowerCase {
		config.UseLowerCase = !defaultSwitches.UseLowerCase
	}
	if switchConf.UseUpperCase {
		config.UseUpperCase = !defaultSwitches.UseUpperCase
	}
	if switchConf.UseNumber {
		config.UseNumber = !defaultSwitches.UseNumber
	}
	if switchConf.UseSpecial {
		config.UseSpecial = !defaultSwitches.UseSpecial
	}
	if switchConf.UseComplex {
		config.UseComplex = !defaultSwitches.UseComplex
	}
	if switchConf.HumanReadable {
		config.HumanReadable = !defaultSwitches.HumanReadable
	}

	// Parse additional parameters and new-style switches
	parseParams(&config)

	return config
}

// Parse the parameters and set the according config flags
func parseParams(config *Config) {
	parseNewStyleParams(config)

	// Complex overrides everything
	if config.UseComplex {
		config.UseUpperCase = true
		config.UseLowerCase = true
		config.UseSpecial = true
		config.UseNumber = true
		config.HumanReadable = false
	}

	if !config.UseUpperCase &&
		!config.UseLowerCase &&
		!config.UseNumber &&
		!config.UseSpecial {
		log.Fatalf("No password mode set. Cannot generate password from empty character set.")
	}

	// Set output mode
	switch config.PwAlgo {
	case 0:
		config.OutputMode = 2
	default:
		config.OutputMode = 0
		if config.SpellPassword {
			config.OutputMode = 1
		}
	}
}

// Parse the new style parameters
func parseNewStyleParams(config *Config) {
	if config.NewStyleModes == "" {
		return
	}

	for _, curParam := range config.NewStyleModes {
		switch curParam {
		case 'S':
			config.UseSpecial = true
		case 's':
			config.UseSpecial = false
		case 'N':
			config.UseNumber = true
		case 'n':
			config.UseNumber = false
		case 'L':
			config.UseLowerCase = true
		case 'l':
			config.UseLowerCase = false
		case 'U':
			config.UseUpperCase = true
		case 'u':
			config.UseUpperCase = false
		case 'H':
			config.HumanReadable = true
		case 'h':
			config.HumanReadable = false
		case 'C':
			config.UseComplex = true
		case 'c':
			config.UseComplex = false
		default:
			log.Fatalf("Unknown password style parameter: %q\n", string(curParam))
		}
	}
}

// GetPwLengthFromParams extracts the password length from the given cli flags and stores
// in the provided config object
func GetPwLengthFromParams(config *Config) int {
	if config.MinPassLen > config.MaxPassLen {
		config.MaxPassLen = config.MinPassLen
	}
	lenDiff := config.MaxPassLen - config.MinPassLen + 1
	randAdd, err := random.GetNum(lenDiff)
	if err != nil {
		log.Fatalf("Failed to generated password length: %v", err)
	}
	retVal := config.MinPassLen + randAdd
	if retVal <= 0 {
		return 1
	}

	return retVal
}
