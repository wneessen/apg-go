// Package main is the APG command line client that makes use of the apg-go library

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/wneessen/apg-go"
)

func main() {
	c := apg.NewConfig()

	// Configure and parse the CLI flags
	// See usage() for flag details
	var ms string
	var co, hr, lc, nu, sp, uc bool
	flag.BoolVar(&lc, "L", false, "")
	flag.BoolVar(&uc, "U", false, "")
	flag.BoolVar(&nu, "N", false, "")
	flag.BoolVar(&sp, "S", false, "")
	flag.BoolVar(&co, "C", false, "")
	flag.BoolVar(&hr, "H", false, "")
	flag.Int64Var(&c.FixedLength, "f", 0, "")
	flag.Int64Var(&c.MinLength, "m", c.MinLength, "")
	flag.Int64Var(&c.MaxLength, "x", c.MaxLength, "")
	flag.StringVar(&ms, "M", "", "")
	flag.Int64Var(&c.NumberPass, "n", c.NumberPass, "")
	flag.Usage = usage
	flag.Parse()

	// Old style character modes
	if hr {
		c.Mode = apg.MaskToggleMode(c.Mode, apg.ModeHumanReadable)
	}
	if lc {
		c.Mode = apg.MaskToggleMode(c.Mode, apg.ModeLowerCase)
	}
	if uc {
		c.Mode = apg.MaskToggleMode(c.Mode, apg.ModeUpperCase)
	}
	if nu {
		c.Mode = apg.MaskToggleMode(c.Mode, apg.ModeNumber)
	}
	if sp {
		c.Mode = apg.MaskToggleMode(c.Mode, apg.ModeSpecial)
	}
	if co {
		c.Mode = apg.MaskSetMode(c.Mode, apg.ModeLowerCase|apg.ModeNumber|
			apg.ModeSpecial|apg.ModeUpperCase)
		c.Mode = apg.MaskClearMode(c.Mode, apg.ModeHumanReadable)
	}

	// New style character modes (has higher priority than the old style modes)
	if ms != "" {
		c.Mode = apg.ModesFromFlags(ms)
	}

	// Generate the password based on the given flags
	g := apg.New(c)
	for i := int64(0); i < c.NumberPass; i++ {
		pl, err := g.GetPasswordLength()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error during password generation: %s\n", err)
		}
		fmt.Printf("PW length: %d\n", pl)
	}
}

// usage is used by the flag package to display the CLI usage message
func usage() {
	// Usage text
	const ut = `apg-go // A "Automated Password Generator"-clone
Copyleft (c) 2021-2023 Winni Neessen

apg [-a <algo>] [-m <length>] [-x <length>] [-L] [-U] [-N] [-S] [-H] [-C]
    [-l] [-M mode] [-E char_string] [-n num_of_pass] [-v] [-h] [-t]

Options:
    -a ALGORITH          Choose the password generation algorithm (Default: 1)
                          - 0: pronounceable password generation (koremutake syllables)
                          - 1: random password generation according to password modes/flags
    -m LENGTH            Minimum length of the password to be generated (Default: 12)
    -x LENGTH            Maximum length of the password to be generated (Default: 20)
    -f LENGTH            Fixed length of the password to be generated (Ignores -m and -x)
    -n NUMBER            Amount of password to be generated (Default: 6)
    -E CHARS             List of characters to be excluded in the generated password
    -M [LUNSHClunshc]    New style password parameters
                          - Note: new-style flags have higher priority than any of the old-style flags
    -L                   Toggle lower case characters in passwords (Default: on)
    -U                   Toggle upper case characters in passwords (Default: on)
    -N                   Toggle numeric characters in passwords (Default: on)
    -S                   Toggle special characters in passwords (Default: off)
    -H                   Avoid ambiguous characters in passwords (i. e.: 1, l, I, O, 0) (Default: off)
    -C                   Enable complex password mode (implies -L -U -N -S and disables -H)
                          - Note: this flag has higher priority than the other old-style flags
    -l                   Spell generated passwords in phonetic alphabet (Default: off)
    -p                   Check the HIBP database if the generated passwords was found in a leak before (Default: off)
                          - Note: this feature requires internet connectivity
    -h                   Show this help text
    -v                   Show version string

`
	_, _ = os.Stderr.WriteString(ut)
}
