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
	"strings"
	"time"
)

// VersionString represents the current version of the apg-go CLI
const VersionString string = "0.4.1"

// Help text
const usage = `apg-go // A "Automated Password Generator"-clone
Copyright (c) 2021 Winni Neessen

apg [-a <algo>] [-m <length>] [-x <length>] [-L] [-U] [-N] [-S] [-H] [-C]
    [-l] [-M mode] [-E char_string] [-n num_of_pass] [-v] [-h] [-t]

Options:
    -a ALGORITH          Choose the password generation algorithm (Default: 1)
                            - 0: pronounceable password generation (koremutake syllables)
                            - 1: random password generation according to password modes/flags
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
                            - Note: this feature requires internet connectivity
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

	pwList := make([]string, 0)
	sylList := map[string][]string{}

	// Choose the type of password generation based on the selected algo
	for i := 0; i < cfgObj.NumOfPass; i++ {
		pwLength := config.GetPwLengthFromParams(&cfgObj)

		switch cfgObj.PwAlgo {
		case 0:
			pwString := ""
			pwSyls := make([]string, 0)

			charSylSet := chars.KoremutakeSyllables
			charSylSet = append(charSylSet,
				strings.Split(chars.PwNumbersHuman, "")...)
			charSylSet = append(charSylSet,
				strings.Split(chars.PwSpecialCharsHuman, "")...)
			charSylSetLen := len(charSylSet)
			for len(pwString) < pwLength {
				randNum, err := random.GetNum(charSylSetLen)
				if err != nil {
					log.Fatalf("error generating Koremutake syllable: %s", err)
				}
				nextSyl := charSylSet[randNum]
				if random.CoinFlip() {
					sylLen := len(nextSyl)
					charPos, err := random.GetNum(sylLen)
					if err != nil {
						log.Fatalf("error generating random number: %s", err)
					}
					ucChar := string(nextSyl[charPos])
					nextSyl = strings.ReplaceAll(nextSyl, ucChar, strings.ToUpper(ucChar))
				}

				pwString += nextSyl
				pwSyls = append(pwSyls, nextSyl)
			}
			pwList = append(pwList, pwString)
			sylList[pwString] = pwSyls
		default:
			charRange := chars.GetRange(&cfgObj)
			pwString, err := random.GetChar(charRange, pwLength)
			if err != nil {
				log.Fatalf("error generating random character range: %s\n", err)
			}
			pwList = append(pwList, pwString)
		}
	}

	for _, p := range pwList {
		switch cfgObj.OutputMode {
		case 1:
			spelledPw, err := spelling.String(p)
			if err != nil {
				log.Fatalf("error spelling out password: %s\n", err)
			}
			fmt.Printf("%v (%v)\n", p, spelledPw)
		case 2:
			fmt.Printf("%s", p)
			if cfgObj.SpellPron {
				spelledPw, err := spelling.Koremutake(sylList[p])
				if err != nil {
					log.Fatalf("error spelling out password: %s", err)
				}
				fmt.Printf(" (%s)", spelledPw)
			}
			fmt.Println()
		default:
			fmt.Println(p)
		}

		if cfgObj.CheckHibp {
			hc := hibp.New(hibp.WithHttpTimeout(time.Second*2), hibp.WithPwnedPadding())
			pwnObj, _, err := hc.PwnedPassApi.CheckPassword(p)
			if err != nil {
				log.Printf("unable to check HIBP database: %v", err)
			}
			if pwnObj != nil && pwnObj.Count != 0 {
				fmt.Print("^-- !!WARNING: The previously generated password was found in HIBP database. Do not use it!!\n")
			}
		}
	}
}
