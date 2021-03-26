package config

import (
	"flag"
	"log"
	"os"
)

// Constants
const DefaultPwLenght int = 20
const VersionString string = "0.2.7"

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
	ExcludeChars  string
	NewStyleModes string
	SpellPassword bool
	ShowHelp      bool
	showVersion   bool
	OutputMode    int
}

var config Config

func init() {
	// Bool flags
	flag.BoolVar(&config.UseLowerCase, "L", true, "Use lower case characters in passwords")
	flag.BoolVar(&config.UseUpperCase, "U", true, "Use upper case characters in passwords")
	flag.BoolVar(&config.UseNumber, "N", true, "Use numbers in passwords")
	flag.BoolVar(&config.UseSpecial, "S", false, "Use special characters in passwords")
	flag.BoolVar(&config.UseComplex, "C", false, "Generate complex passwords (implies -L -U -N -S, disables -H)")
	flag.BoolVar(&config.SpellPassword, "l", false, "Spell generated password")
	flag.BoolVar(&config.HumanReadable, "H", false, "Generate human-readable passwords")
	flag.BoolVar(&config.showVersion, "v", false, "Show version")

	// Int flags
	flag.IntVar(&config.MinPassLen, "m", DefaultPwLenght, "Minimum password length")
	flag.IntVar(&config.MaxPassLen, "x", DefaultPwLenght, "Maxiumum password length")
	flag.IntVar(&config.NumOfPass, "n", 1, "Number of passwords to generate")

	// String flags
	flag.StringVar(&config.ExcludeChars, "E", "", "Exclude list of characters from generated password")
	flag.StringVar(&config.NewStyleModes, "M", "",
		"New style password parameters (higher priority than single parameters)")
}

func NewConfig() *Config {
	flag.Parse()

	if config.showVersion {
		_, _ = os.Stderr.WriteString("Advanced Password Generator Clone (apg.go) v" + VersionString + "\n")
		_, _ = os.Stderr.WriteString("(C) 2021 by Winni Neessen\n")
		os.Exit(0)
	}

	return &config
}

// Parse the parameters and set the according config flags
func (config *Config) ParseParams() {
	config.parseNewStyleParams()

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

// Get the password length from the given cli flags
func (config *Config) GetPwLengthFromParams() int {
	pwLength := config.MinPassLen
	if pwLength < config.MinPassLen {
		pwLength = config.MinPassLen
	}
	if pwLength > config.MaxPassLen {
		pwLength = config.MaxPassLen
	}

	return pwLength
}

// Parse the new style parameters
func (config *Config) parseNewStyleParams() {
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
