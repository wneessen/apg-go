package main

import (
	"fmt"
	"os"
)

// Parse the parameters and set the according config flags
func parseParams() {
	parseNewStyleParams()

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
		fmt.Printf("No password mode set. Cannot generate password from empty character set.")
		os.Exit(1)
	}

	// Set output mode
	if config.spellPassword {
		config.outputMode = 1
	}
}

// Get the password length from the given cli flags
func getPwLengthFromParams() int {
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
func parseNewStyleParams() {
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
			fmt.Printf("Unknown password style parameter: %q\n", string(curParam))
			os.Exit(1)
		}
	}
}
