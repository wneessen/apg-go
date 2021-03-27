package main

import (
	"flag"
	"log"
)

// Parse the CLI flags
func parseFlags() Config {
	var config Config

	// Read and set all flags
	flag.BoolVar(&config.useLowerCase, "L", true, "Use lower case characters in passwords")
	flag.BoolVar(&config.useUpperCase, "U", true, "Use upper case characters in passwords")
	flag.BoolVar(&config.useNumber, "N", true, "Use numbers in passwords")
	flag.BoolVar(&config.useSpecial, "S", false, "Use special characters in passwords")
	flag.BoolVar(&config.useComplex, "C", false, "Generate complex passwords (implies -L -U -N -S, disables -H)")
	flag.BoolVar(&config.spellPassword, "l", false, "Spell generated password")
	flag.BoolVar(&config.humanReadable, "H", false, "Generate human-readable passwords")
	flag.BoolVar(&config.showVersion, "v", false, "Show version")
	flag.IntVar(&config.minPassLen, "m", DefaultPwLenght, "Minimum password length")
	flag.IntVar(&config.maxPassLen, "x", DefaultPwLenght, "Maxiumum password length")
	flag.IntVar(&config.numOfPass, "n", 1, "Number of passwords to generate")
	flag.StringVar(&config.excludeChars, "E", "", "Exclude list of characters from generated password")
	flag.StringVar(&config.newStyleModes, "M", "",
		"New style password parameters (higher priority than single parameters)")
	flag.Parse()

	// Parse additional parameters
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
	pwLength := config.minPassLen
	if pwLength < config.minPassLen {
		pwLength = config.minPassLen
	}
	if pwLength > config.maxPassLen {
		pwLength = config.maxPassLen
	}

	return pwLength
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
