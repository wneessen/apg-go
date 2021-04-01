package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// Constants
const DefaultPwLenght int = 20
const VersionString string = "0.3.0"

type Config struct {
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
	ShowHelp      bool
	showVersion   bool
	outputMode    int
}

// Help text
const usage = `apg-go // A "Automated Password Generator"-clone
Copyright (c) 2021 Winni Neessen

apg [-m <length>] [-x <length>] [-L] [-U] [-N] [-S] [-H] [-C]
    [-l] [-M mode] [-E char_string] [-n num_of_pass] [-v] [-h]

Options:
    -m LENGTH            Minimum length of the password to be generated (Default: 20)
    -x LENGTH            Maximum length of the password to be generated (Default: 20)
    -n NUMBER            Amount of password to be generated (Default: 1)
    -E CHARS             List of characters to be excluded in the generated password
    -M [LUNSHClunshc]    New style password parameters (upper case: on, lower case: off)
    -L                   Use lower case characters in passwords (Default: on)
    -U                   Use upper case characters in passwords (Default: on)
    -N                   Use numeric characters in passwords (Default: on)
    -S                   Use special characters in passwords (Default: on)
    -H                   Avoid ambiguous characters in passwords (i. e.: 1, l, I, O, 0) (Default: off)
    -C                   Enable complex password mode (implies -L -U -N -S and disables -H) (Default: off)
    -l                   Spell generated passwords in phonetic alphabet (Default: off)
    -h                   Show this help text
    -v                   Show version string`

// Main function that generated the passwords and returns them
func main() {
	// Log config
	log.SetFlags(log.Ltime | log.Ldate | log.Lshortfile)

	// Read and parse flags
	flag.Usage = func() { _, _ = fmt.Fprintf(os.Stderr, "%s\n", usage) }
	var config = parseFlags()

	// Show version and exit
	if config.showVersion {
		_, _ = os.Stderr.WriteString(`apg-go // A "Automated Password Generator"-clone v` + VersionString + "\n")
		_, _ = os.Stderr.WriteString("(C) 2021 by Winni Neessen\n")
		os.Exit(0)
	}

	// Set PW length and available characterset
	pwLength := getPwLengthFromParams(&config)
	charRange := getCharRange(&config)

	// Generate passwords
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
