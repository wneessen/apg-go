// Package main is the APG command line client that makes use of the apg-go library

package main

import (
	"flag"
	"os"

	"github.com/wneessen/apg-go"
)

func main() {
	c := apg.NewConfig()

	// Configure and parse the CLI flags
	flag.Int64Var(&c.MinLength, "m", c.MinLength, "")
	flag.Int64Var(&c.MaxLength, "x", c.MaxLength, "")
	flag.Int64Var(&c.NumberPass, "n", c.NumberPass, "")
	flag.Usage = usage
	flag.Parse()

	/*
		g := apg.New(c)
		rb, err := g.RandomBytes(c.MinLength)
		if err != nil {
			fmt.Println("ERROR", err)
			os.Exit(1)
		}
		fmt.Printf("Random: %#v\n", rb)

	*/
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
    -v                   Show version string

`

	_, _ = os.Stderr.WriteString(ut)
}
