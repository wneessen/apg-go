package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// Constants
const DefaultPwLenght int = 20
const VersionString string = "0.2.6"
const PwLowerCharsHuman string = "abcdefghjkmnpqrstuvwxyz"
const PwUpperCharsHuman string = "ABCDEFGHJKMNPQRSTUVWXYZ"
const PwLowerChars string = "abcdefghijklmnopqrstuvwxyz"
const PwUpperChars string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const PwSpecialCharsHuman string = "\"#%*+-/:;=\\_|~"
const PwSpecialChars string = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
const PwNumbersHuman string = "23456789"
const PwNumbers string = "1234567890"

type cliOpts struct {
	minPassLen    int
	maxPassLen    int
	numOfPass     int
	useComplex    bool
	useLowerCase  bool
	useUpperCase  bool
	useNumber     bool
	useSpecial    bool
	humanReadable bool
	excludeChars  string
	newStyleModes string
	spellPassword bool
	showHelp      bool
	showVersion   bool
	outputMode    int
}

var config cliOpts

// Read flags
func init() {
	// Bool flags
	flag.BoolVar(&config.useLowerCase, "L", true, "Use lower case characters in passwords")
	flag.BoolVar(&config.useUpperCase, "U", true, "Use upper case characters in passwords")
	flag.BoolVar(&config.useNumber, "N", true, "Use numbers in passwords")
	flag.BoolVar(&config.useSpecial, "S", false, "Use special characters in passwords")
	flag.BoolVar(&config.useComplex, "C", false, "Generate complex passwords (implies -L -U -N -S, disables -H)")
	flag.BoolVar(&config.spellPassword, "l", false, "Spell generated password")
	flag.BoolVar(&config.humanReadable, "H", false, "Generate human-readable passwords")
	flag.BoolVar(&config.showVersion, "v", false, "Show version")

	// Int flags
	flag.IntVar(&config.minPassLen, "m", DefaultPwLenght, "Minimum password length")
	flag.IntVar(&config.maxPassLen, "x", DefaultPwLenght, "Maxiumum password length")
	flag.IntVar(&config.numOfPass, "n", 1, "Number of passwords to generate")

	// String flags
	flag.StringVar(&config.excludeChars, "E", "", "Exclude list of characters from generated password")
	flag.StringVar(&config.newStyleModes, "M", "",
		"New style password parameters (higher priority than single parameters)")

	flag.Parse()
	if config.showVersion {
		_, _ = os.Stderr.WriteString("Winni's Advanced Password Generator Clone (apg.go) v" + VersionString + "\n")
		_, _ = os.Stderr.WriteString("(C) 2021 by Winni Neessen\n")
		os.Exit(0)
	}

	log.SetFlags(log.Ltime | log.Ldate | log.Lshortfile)
}

// Main function that generated the passwords and returns them
func main() {
	parseParams()
	pwLength := getPwLengthFromParams()
	charRange := getCharRange()

	for i := 1; i <= config.numOfPass; i++ {
		pwString, err := getRandChar(&charRange, pwLength)
		if err != nil {
			log.Fatalf("getRandChar returned an error: %q\n", err)
		}

		switch config.outputMode {
		case 1:
			{
				spelledPw, err := spellPasswordString(pwString)
				if err != nil {
					log.Fatalf("spellPasswordString returned an error: %q\n", err.Error())
				}
				fmt.Printf("%v (%v)\n", pwString, spelledPw)
				break
			}
		default:
			{
				fmt.Println(pwString)
				break
			}
		}
	}
}
