package main

import (
	"flag"
	"log"
)

// Parse the CLI flags
func parseFlags() Config {
	var switchConf Config
	defaultSwitches := Config{
		useLowerCase:  true,
		useUpperCase:  true,
		useNumber:     true,
		useSpecial:    false,
		useComplex:    false,
		humanReadable: false,
	}
	config := Config{
		useLowerCase:  defaultSwitches.useLowerCase,
		useUpperCase:  defaultSwitches.useUpperCase,
		useNumber:     defaultSwitches.useNumber,
		useSpecial:    defaultSwitches.useSpecial,
		useComplex:    defaultSwitches.useComplex,
		humanReadable: defaultSwitches.humanReadable,
	}

	// Read and set all flags
	flag.BoolVar(&switchConf.useLowerCase, "L", false, "Use lower case characters in passwords")
	flag.BoolVar(&switchConf.useUpperCase, "U", false, "Use upper case characters in passwords")
	flag.BoolVar(&switchConf.useNumber, "N", false, "Use numerich characters in passwords")
	flag.BoolVar(&switchConf.useSpecial, "S", false, "Use special characters in passwords")
	flag.BoolVar(&switchConf.useComplex, "C", false, "Generate complex passwords (implies -L -U -N -S, disables -H)")
	flag.BoolVar(&switchConf.humanReadable, "H", false, "Generate human-readable passwords")
	flag.BoolVar(&config.spellPassword, "l", false, "Spell generated password")
	flag.BoolVar(&config.showVersion, "v", false, "Show version")
	flag.IntVar(&config.minPassLen, "m", DefaultMinLenght, "Minimum password length")
	flag.IntVar(&config.maxPassLen, "x", DefaultMaxLenght, "Maxiumum password length")
	flag.IntVar(&config.numOfPass, "n", 6, "Number of passwords to generate")
	flag.StringVar(&config.excludeChars, "E", "", "Exclude list of characters from generated password")
	flag.StringVar(&config.newStyleModes, "M", "",
		"New style password parameters (higher priority than single parameters)")
	flag.Parse()

	// Invert-switch the defaults
	if switchConf.useLowerCase {
		config.useLowerCase = !defaultSwitches.useLowerCase
	}
	if switchConf.useUpperCase {
		config.useUpperCase = !defaultSwitches.useUpperCase
	}
	if switchConf.useNumber {
		config.useNumber = !defaultSwitches.useNumber
	}
	if switchConf.useSpecial {
		config.useSpecial = !defaultSwitches.useSpecial
	}
	if switchConf.useComplex {
		config.useComplex = !defaultSwitches.useComplex
	}
	if switchConf.humanReadable {
		config.humanReadable = !defaultSwitches.humanReadable
	}

	// Parse additional parameters and new-style switches
	parseParams(&config)

	return config
}

// Parse the parameters and set the according config flags
func parseParams(config *Config) {
	parseNewStyleParams(config)

	// Complex overrides everything
	if config.useComplex {
		config.useUpperCase = true
		config.useLowerCase = true
		config.useSpecial = true
		config.useNumber = true
		config.humanReadable = false
	}

	if config.useUpperCase == false &&
		config.useLowerCase == false &&
		config.useNumber == false &&
		config.useSpecial == false {
		log.Fatalf("No password mode set. Cannot generate password from empty character set.")
	}

	// Set output mode
	if config.spellPassword {
		config.outputMode = 1
	}
}

// Get the password length from the given cli flags
func getPwLengthFromParams(config *Config) int {
	if config.minPassLen > config.maxPassLen {
		config.maxPassLen = config.minPassLen
	}
	lenDiff := config.maxPassLen - config.minPassLen + 1
	randAdd, err := getRandNum(lenDiff)
	if err != nil {
		log.Fatalf("Failed to generated password length: %v", err)
	}
	retVal := config.minPassLen + randAdd
	if retVal <= 0 {
		return 1
	}

	return retVal
}

// Parse the new style parameters
func parseNewStyleParams(config *Config) {
	if config.newStyleModes == "" {
		return
	}

	for _, curParam := range config.newStyleModes {
		switch curParam {
		case 'S':
			config.useSpecial = true
			break
		case 's':
			config.useSpecial = false
			break
		case 'N':
			config.useNumber = true
			break
		case 'n':
			config.useNumber = false
			break
		case 'L':
			config.useLowerCase = true
			break
		case 'l':
			config.useLowerCase = false
			break
		case 'U':
			config.useUpperCase = true
			break
		case 'u':
			config.useUpperCase = false
			break
		case 'H':
			config.humanReadable = true
			break
		case 'h':
			config.humanReadable = false
			break
		case 'C':
			config.useComplex = true
			break
		case 'c':
			config.useComplex = false
			break
		default:
			log.Fatalf("Unknown password style parameter: %q\n", string(curParam))
		}
	}
}
