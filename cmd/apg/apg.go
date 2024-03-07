// Package main is the APG command line client that makes use of the apg-go library

package main

import (
	"flag"
	"fmt"
	"os"

	"src.neessen.cloud/wneessen/apg-go"
)

// MinimumAmountTooHigh is an error message displayed when a minimum amount of
// parameter has been set to a too high value
const MinimumAmountTooHigh = "WARNING: You have selected a minimum amount of characters that is bigger\n" +
	"than 50% of the minimum password length to be generated. This can lead\n" +
	"to extraordinary calculation times resulting in apg-go never finishing\n" +
	"the job. Please consider lowering the value.\n\n"

func main() {
	c := apg.NewConfig()

	// Configure and parse the CLI flags
	// See usage() for flag details
	var al int
	var ms string
	var co, hr, lc, nu, sp, uc bool
	flag.IntVar(&al, "a", 1, "")
	flag.BoolVar(&lc, "L", false, "")
	flag.Int64Var(&c.MinLowerCase, "mL", c.MinLowerCase, "")
	flag.BoolVar(&uc, "U", false, "")
	flag.Int64Var(&c.MinUpperCase, "mU", c.MinUpperCase, "")
	flag.BoolVar(&nu, "N", false, "")
	flag.Int64Var(&c.MinNumeric, "mN", c.MinNumeric, "")
	flag.BoolVar(&sp, "S", false, "")
	flag.Int64Var(&c.MinSpecial, "mS", c.MinSpecial, "")
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
		c.Mode = apg.MaskToggleMode(c.Mode, apg.ModeNumeric)
	}
	if sp {
		c.Mode = apg.MaskToggleMode(c.Mode, apg.ModeSpecial)
	}
	if co {
		c.Mode = apg.MaskSetMode(c.Mode, apg.ModeLowerCase|apg.ModeNumeric|
			apg.ModeSpecial|apg.ModeUpperCase)
		c.Mode = apg.MaskClearMode(c.Mode, apg.ModeHumanReadable)
	}

	// New style character modes (has higher priority than the old style modes)
	if ms != "" {
		c.Mode = apg.ModesFromFlags(ms)
	}

	// For the "minimum amount of" modes we need to imply at the type
	// of character mode is set
	if c.MinLowerCase > 0 {
		if float64(c.MinLength)/2 < float64(c.MinNumeric) {
			_, _ = os.Stderr.WriteString(MinimumAmountTooHigh)
		}
		c.Mode = apg.MaskSetMode(c.Mode, apg.ModeLowerCase)
	}
	if c.MinNumeric > 0 {
		if float64(c.MinLength)/2 < float64(c.MinLowerCase) {
			_, _ = os.Stderr.WriteString(MinimumAmountTooHigh)
		}
		c.Mode = apg.MaskSetMode(c.Mode, apg.ModeNumeric)
	}
	if c.MinSpecial > 0 {
		if float64(c.MinLength)/2 < float64(c.MinSpecial) {
			_, _ = os.Stderr.WriteString(MinimumAmountTooHigh)
		}
		c.Mode = apg.MaskSetMode(c.Mode, apg.ModeSpecial)
	}
	if c.MinUpperCase > 0 {
		if float64(c.MinLength)/2 < float64(c.MinUpperCase) {
			_, _ = os.Stderr.WriteString(MinimumAmountTooHigh)
		}
		c.Mode = apg.MaskSetMode(c.Mode, apg.ModeUpperCase)
	}

	// Check if algorithm is supported
	c.Algorithm = apg.IntToAlgo(al)
	if c.Algorithm == apg.AlgoUnsupported {
		_, _ = fmt.Fprintf(os.Stderr, "unsupported algorithm value: %d\n", al)
		os.Exit(1)
	}

	// Generate the password based on the given flags
	g := apg.New(c)
	for i := int64(0); i < c.NumberPass; i++ {
		p, err := g.Generate()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to generate password: %s\n", err)
			os.Exit(1)
		}
		fmt.Println(p)
	}
}

// usage is used by the flag package to display the CLI usage message
func usage() {
	// Usage text
	const ut = `apg-go v` +
		apg.VERSION + "\n" +
		`A OSS "Automated Password Generator"-clone -- https://src.neessen.cloud/wneessen/apg-go/
Created 2021-2024 by Winni Neessen (MIT licensed)

apg [-a <algo>] [-m <length>] [-x <length>] [-L] [-U] [-N] [-S] [-H] [-C]
    [-l] [-M mode] [-E char_string] [-n num_of_pass] [-v] [-h] [-t]

Flags:
    -a ALGORITH          Choose the password generation algorithm (Default: 1)
                          - 0: pronounceable password generation (koremutake syllables)
                          - 1: random password generation according to password modes/flags
                          - 2: coinflip (returns heads or tails)
    -m LENGTH            Minimum length of the password to be generated (Default: 12)
    -x LENGTH            Maximum length of the password to be generated (Default: 20)
    -f LENGTH            Fixed length of the password to be generated (Ignores -m and -x)
    -n NUMBER            Amount of password to be generated (Default: 6)
    -E CHARS             List of characters to be excluded in the generated password
    -M [LUNSHClunshc]    New style password flags
                          - Note: new-style flags have higher priority than any of the old-style flags
    -mL NUMBER           Minimum amount of lower-case characters (implies -L)
    -mN NUMBER           Minimum amount of numeric characters (imlies -N)
    -mS NUMBER           Minimum amount of special characters (imlies -S)
    -mU NUMBER           Minimum amount of upper-case characters (imlies -U)
                          - Note: any of the "Minimum amount of" modes may result in
                            extraordinarily long calculation times
    -C                   Enable complex password mode (implies -L -U -N -S and disables -H)
    -H                   Avoid ambiguous characters in passwords (i. e.: 1, l, I, O, 0) (Default: off)
    -L                   Toggle lower-case characters in passwords (Default: on)
    -N                   Toggle numeric characters in passwords (Default: on)
    -S                   Toggle special characters in passwords (Default: off)
    -U                   Toggle upper-case characters in passwords (Default: on)
                          - Note: this flag has higher priority than the other old-style flags
    -l                   Spell generated passwords in phonetic alphabet (Default: off)
    -p                   Check the HIBP database if the generated passwords was found in a leak before (Default: off)
                          - Note: this feature requires internet connectivity
    -h                   Show this help text
    -v                   Show version string`

	_, _ = os.Stderr.WriteString(ut + "\n\n")
}
