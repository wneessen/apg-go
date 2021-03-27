package main

import (
	"fmt"
	"log"
	"os"
)

// Constants
const DefaultPwLenght int = 20
const VersionString string = "0.2.9"

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

// Main function that generated the passwords and returns them
func main() {
	// Log config
	log.SetFlags(log.Ltime | log.Ldate | log.Lshortfile)

	// Read and parse flags
	var config = parseFlags()

	// Show version and exit
	if config.showVersion {
		_, _ = os.Stderr.WriteString("Advanced Password Generator Clone (apg.go) v" + VersionString + "\n")
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
