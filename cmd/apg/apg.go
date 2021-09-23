package main

import (
	"flag"
	"fmt"
	"github.com/wneessen/apg-go/chars"
	"github.com/wneessen/apg-go/config"
	"github.com/wneessen/apg-go/random"
	"github.com/wneessen/apg-go/spelling"
	"github.com/wneessen/go-hibp"
	"log"
	"os"
	"time"
)

const VersionString string = "0.4.0-dev"

// Help text
const usage = `apg-go // A "Automated Password Generator"-clone
Copyright (c) 2021 Winni Neessen

apg [-m <length>] [-x <length>] [-L] [-U] [-N] [-S] [-H] [-C]
    [-l] [-M mode] [-E char_string] [-n num_of_pass] [-v] [-h]

Options:
    -m LENGTH            Minimum length of the password to be generated (Default: 12)
    -x LENGTH            Maximum length of the password to be generated (Default: 20)
    -n NUMBER            Amount of password to be generated (Default: 6)
    -E CHARS             List of characters to be excluded in the generated password
    -M [LUNSHClunshc]    New style password parameters (upper case: on, lower case: off)
    -L                   Use lower case characters in passwords (Default: on)
    -U                   Use upper case characters in passwords (Default: on)
    -N                   Use numeric characters in passwords (Default: on)
    -S                   Use special characters in passwords (Default: off)
    -H                   Avoid ambiguous characters in passwords (i. e.: 1, l, I, O, 0) (Default: off)
    -C                   Enable complex password mode (implies -L -U -N -S and disables -H) (Default: off)
    -l                   Spell generated passwords in phonetic alphabet (Default: off)
    -p                   Check the HIBP database if the generated passwords was found in a leak before (Default: off)
                         '--> this feature requires internet connectivity 
    -h                   Show this help text
    -v                   Show version string`

// Main function that generated the passwords and returns them
func main() {
	// Log configuration
	log.SetFlags(log.Ltime | log.Ldate | log.Lshortfile)

	// Read and parse flags
	flag.Usage = func() { _, _ = fmt.Fprintf(os.Stderr, "%s\n", usage) }
	var cfgObj = config.New()

	// Show version and exit
	if cfgObj.ShowVersion {
		_, _ = os.Stderr.WriteString(`apg-go // A "Automated Password Generator"-clone v` + VersionString + "\n")
		_, _ = os.Stderr.WriteString("(C) 2021 by Winni Neessen\n")
		os.Exit(0)
	}

	// Set PW length and available characterset
	charRange := chars.GetRange(&cfgObj)

	// Generate passwords
	for i := 1; i <= cfgObj.NumOfPass; i++ {
		pwLength := config.GetPwLengthFromParams(&cfgObj)
		pwString, err := random.GetChar(&charRange, pwLength)
		if err != nil {
			log.Fatalf("error generating random character range: %s\n", err)
		}

		switch cfgObj.OutputMode {
		case 1:
			{
				spelledPw, err := spelling.String(pwString)
				if err != nil {
					log.Fatalf("error spelling out password: %s\n", err)
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

		if cfgObj.CheckHibp {
			hc := hibp.New(hibp.WithHttpTimeout(time.Second*2), hibp.WithPwnedPadding())
			pwnObj, _, err := hc.PwnedPassApi.CheckPassword(pwString)
			if err != nil {
				log.Printf("unable to check HIBP database: %v", err)
			}
			if pwnObj != nil && pwnObj.Count != 0 {
				fmt.Print("^-- !!WARNING: The previously generated password was found in HIPB database. Do not use it!!\n")
			}
		}
	}
}
