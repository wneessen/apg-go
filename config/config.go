package config

import (
	"flag"
	"github.com/wneessen/apg-go/random"
	"log"
)

// Config is a struct that holds the different config parameters for the apg-go
// application
type Config struct {
	MinPassLen    int
	MaxPassLen    int
	NumOfPass     int
	UseComplex    bool
	UseLowerCase  bool
	UseUpperCase  bool
	UseNumber     bool
	UseSpecial    bool
	HumanReadable bool
	CheckHibp     bool
	ExcludeChars  string
	NewStyleModes string
	SpellPassword bool
	ShowHelp      bool
	ShowVersion   bool
	OutputMode    int
}

// DefaultMinLength reflects the default minimum length of a generated password
const DefaultMinLength int = 12

// DefaultMaxLength reflects the default maximum length of a generated password
const DefaultMaxLength int = 20

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

	if config.UseUpperCase == false &&
		config.UseLowerCase == false &&
		config.UseNumber == false &&
		config.UseSpecial == false {
		log.Fatalf("No password mode set. Cannot generate password from empty character set.")
	}

	// Set output mode
	if config.SpellPassword {
		config.OutputMode = 1
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
			break
		case 's':
			config.UseSpecial = false
			break
		case 'N':
			config.UseNumber = true
			break
		case 'n':
			config.UseNumber = false
			break
		case 'L':
			config.UseLowerCase = true
			break
		case 'l':
			config.UseLowerCase = false
			break
		case 'U':
			config.UseUpperCase = true
			break
		case 'u':
			config.UseUpperCase = false
			break
		case 'H':
			config.HumanReadable = true
			break
		case 'h':
			config.HumanReadable = false
			break
		case 'C':
			config.UseComplex = true
			break
		case 'c':
			config.UseComplex = false
			break
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
