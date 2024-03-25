// SPDX-FileCopyrightText: 2021-2024 Winni Neessen <wn@neessen.dev>
//
// SPDX-License-Identifier: MIT

// Package main is the APG command line client that makes use of the apg-go library

package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/wneessen/apg-go"
)

// MinimumAmountTooHigh is an error message displayed when a minimum amount of
// parameter has been set to a too high value
const MinimumAmountTooHigh = "WARNING: You have selected a minimum amount of characters that is bigger\n" +
	"than 50% of the minimum password length to be generated. This can lead\n" +
	"to extraordinary calculation times resulting in apg-go never finishing\n" +
	"the job. Please consider lowering the value.\n\n"

func main() {
	config := apg.NewConfig()

	// Configure and parse the CLI flags
	// See usage() for flag details
	var algorithm int
	var modeString string
	var complexPass, humanReadable, lowerCase, numeric, special, showVer, upperCase bool
	flag.IntVar(&algorithm, "a", 1, "")
	flag.BoolVar(&config.BinaryHexMode, "bh", false, "")
	flag.BoolVar(&config.BinaryNewline, "bn", false, "")
	flag.BoolVar(&complexPass, "C", false, "")
	flag.StringVar(&config.ExcludeChars, "E", "", "")
	flag.Int64Var(&config.FixedLength, "f", 0, "")
	flag.BoolVar(&config.MobileGrouping, "g", false, "")
	flag.BoolVar(&humanReadable, "H", false, "")
	flag.BoolVar(&config.SpellPassword, "l", false, "")
	flag.BoolVar(&lowerCase, "L", false, "")
	flag.Int64Var(&config.MinLength, "m", config.MinLength, "")
	flag.Int64Var(&config.MinLowerCase, "mL", config.MinLowerCase, "")
	flag.Int64Var(&config.MinNumeric, "mN", config.MinNumeric, "")
	flag.Int64Var(&config.MinSpecial, "mS", config.MinSpecial, "")
	flag.Int64Var(&config.MinUpperCase, "mU", config.MinUpperCase, "")
	flag.Int64Var(&config.NumberPass, "n", config.NumberPass, "")
	flag.StringVar(&modeString, "M", "", "")
	flag.BoolVar(&numeric, "N", false, "")
	flag.BoolVar(&config.CheckHIBP, "p", false, "")
	flag.BoolVar(&special, "S", false, "")
	flag.BoolVar(&config.SpellPronounceable, "t", false, "")
	flag.BoolVar(&upperCase, "U", false, "")
	flag.BoolVar(&showVer, "v", false, "")
	flag.Int64Var(&config.MaxLength, "x", config.MaxLength, "")
	flag.Usage = usage
	flag.Parse()

	// Show version and exit
	if showVer {
		_, _ = os.Stderr.WriteString(`apg-go // A "Automated Password Generator"-clone ` +
			`v` + apg.VERSION + "\n")
		_, _ = os.Stderr.WriteString("OS: " + runtime.GOOS + " // Arch: " +
			runtime.GOARCH + " \n")
		_, _ = os.Stderr.WriteString("(C) 2021-2024 by Winni Neessen\n")
		os.Exit(0)
	}

	// Old style character modes
	configOldStyle(config, humanReadable, lowerCase, upperCase, numeric,
		special, complexPass)

	// New style character modes (has higher priority than the old style modes)
	if modeString != "" {
		config.Mode = apg.ModesFromFlags(modeString)
	}

	// For the "minimum amount of" modes we need to imply at the type
	// of character mode is set
	configMinRequirement(config)

	// Check if algorithm is supported
	config.Algorithm = apg.IntToAlgo(algorithm)
	if config.Algorithm == apg.AlgoUnsupported {
		_, _ = fmt.Fprintf(os.Stderr, "unsupported algorithm value: %d\n", algorithm)
		os.Exit(1)
	}

	// Generate the password based on the given flags and print it to stdout
	generate(config)
}

// configMinRequirement configures the "minimum amount" feature
func configMinRequirement(config *apg.Config) {
	if config.MinLowerCase > 0 {
		if float64(config.MinLength)/2 < float64(config.MinNumeric) {
			_, _ = os.Stderr.WriteString(MinimumAmountTooHigh)
		}
		config.Mode = apg.MaskSetMode(config.Mode, apg.ModeLowerCase)
	}
	if config.MinNumeric > 0 {
		if float64(config.MinLength)/2 < float64(config.MinLowerCase) {
			_, _ = os.Stderr.WriteString(MinimumAmountTooHigh)
		}
		config.Mode = apg.MaskSetMode(config.Mode, apg.ModeNumeric)
	}
	if config.MinSpecial > 0 {
		if float64(config.MinLength)/2 < float64(config.MinSpecial) {
			_, _ = os.Stderr.WriteString(MinimumAmountTooHigh)
		}
		config.Mode = apg.MaskSetMode(config.Mode, apg.ModeSpecial)
	}
	if config.MinUpperCase > 0 {
		if float64(config.MinLength)/2 < float64(config.MinUpperCase) {
			_, _ = os.Stderr.WriteString(MinimumAmountTooHigh)
		}
		config.Mode = apg.MaskSetMode(config.Mode, apg.ModeUpperCase)
	}
}

// configOldStyle configures the old style character modes
func configOldStyle(config *apg.Config, humanReadable, lowerCase, upperCase,
	numeric, special, complexPass bool,
) {
	if humanReadable {
		config.Mode = apg.MaskToggleMode(config.Mode, apg.ModeHumanReadable)
	}
	if lowerCase {
		config.Mode = apg.MaskToggleMode(config.Mode, apg.ModeLowerCase)
	}
	if upperCase {
		config.Mode = apg.MaskToggleMode(config.Mode, apg.ModeUpperCase)
	}
	if numeric {
		config.Mode = apg.MaskToggleMode(config.Mode, apg.ModeNumeric)
	}
	if special {
		config.Mode = apg.MaskToggleMode(config.Mode, apg.ModeSpecial)
	}
	if complexPass {
		config.Mode = apg.MaskSetMode(config.Mode, apg.ModeLowerCase|apg.ModeNumeric|
			apg.ModeSpecial|apg.ModeUpperCase)
		config.Mode = apg.MaskClearMode(config.Mode, apg.ModeHumanReadable)
	}
}

func generate(config *apg.Config) {
	generator := apg.New(config)

	// In binary mode we only generate a single secret
	if config.Algorithm == apg.AlgoBinary {
		password, err := generator.Generate()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to generate password: %s\n", err)
			os.Exit(1)
		}
		if config.BinaryNewline {
			fmt.Println(password)
			return
		}
		fmt.Print(password)
		return
	}

	// For any other mode we cycle through the amount of passwords to be generated
	for i := int64(0); i < config.NumberPass; i++ {
		password, err := generator.Generate()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to generate password: %s\n", err)
			os.Exit(1)
		}
		if config.Algorithm == apg.AlgoRandom && config.SpellPassword {
			spellPass, err := apg.Spell(password)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "failed to spell password: %s\n", err)
			}
			fmt.Printf("%s (%s)\n", password, spellPass)
			continue
		}
		if config.Algorithm == apg.AlgoPronounceable && config.SpellPronounceable {
			pronouncePass, err := generator.Pronounce()
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "failed to pronounce password: %s\n", err)
			}
			fmt.Printf("%s (%s)\n", password, pronouncePass)
			continue
		}
		fmt.Println(password)

		if config.CheckHIBP {
			pwned, err := apg.HasBeenPwned(password)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "failed to check HIBP database: %s\n", err)
			}
			if pwned {
				fmt.Print("^-- !!WARNING: The previously generated password was found in " +
					"HIBP database. Do not use it!!\n")
			}
		}
	}
}

// usage is used by the flag package to display the CLI usage message
func usage() {
	// Usage text
	const ut = `apg-go v` +
		apg.VERSION + "\n" +
		`A OSS "Automated Password Generator"-clone -- https://github.com/wneessen/apg-go/
Created 2021-2024 by Winni Neessen (MIT licensed)

apg [-a <algo>] [-m <length>] [-x <length>] [-L] [-U] [-N] [-S] [-H] [-C]
    [-l] [-M mode] [-E char_string] [-n num_of_pass] [-mX number] [-t] [-p] [-v] [-h]

Flags:
    -a ALGORITH          Choose the password generation algorithm (Default: 1)
                          - 0: pronounceable password generation (koremutake syllables)
                          - 1: random password generation according to password modes/flags
                          - 2: coinflip (returns heads or tails)
                          - 3: full binary mode (generates simple 256 bit randomness)
    -bh                  When set, will print the generated secret in its hex representation (Default: off)
    -bn                  When set, will return a new line character after the generated secret (Default: off)
                          - Note: The -bX options only apply to binary mode (Algo: 3)
    -m LENGTH            Minimum length of the password to be generated (Default: 12)
    -x LENGTH            Maximum length of the password to be generated (Default: 20)
    -f LENGTH            Fixed length of the password to be generated (Ignores -m and -x)
                          - Note: Due to the way the pronounceable password algorithm works,
	                        this setting might not always apply
    -g                   When set, mobile-friendly character grouping will be enabled in Algo: 1
                          - Note: Grouping characters in random passwords makes them much
                            more predictable and lowers the entropy of the generated password.
    -n NUMBER            Amount of password to be generated (Default: 6)
                          - Note: Does not apply to binary mode (Algo: 3)
    -E CHARS             List of characters to be excluded in the generated password
    -M [LUNSHClunshc]    New style password flags
                          - Note: new-style flags have higher priority than any of the old-style flags
    -mL NUMBER           Minimum amount of lower-case characters (implies -L)
    -mN NUMBER           Minimum amount of numeric characters (implies -N)
    -mS NUMBER           Minimum amount of special characters (implies -S)
    -mU NUMBER           Minimum amount of upper-case characters (implies -U)
                          - Note: any of the "Minimum amount of" modes may result in
                            extraordinarily long calculation times
                          - Note 2: The "minimum amount of" modes do not apply in
                            pronounceable mode (-a 0)
    -C                   Enable complex password mode (implies -L -U -N -S and disables -H)
    -H                   Avoid ambiguous characters in passwords (i. e.: 1, l, I, O, 0) (Default: off)
    -L                   Toggle lower-case characters in passwords (Default: on)
    -N                   Toggle numeric characters in passwords (Default: on)
    -S                   Toggle special characters in passwords (Default: off)
    -U                   Toggle upper-case characters in passwords (Default: on)
                          - Note: this flag has higher priority than the other old-style flags
    -l                   Spell generated passwords in phonetic alphabet (Default: off)
    -t                   Spell generated pronounceable passwords with the corresponding 
                         syllables (Default: off)
    -p                   Check the HIBP database if the generated passwords was found in a leak before (Default: off)
                          - Note: this feature requires internet connectivity
    -h                   Show this help text
    -v                   Show version string`

	_, _ = os.Stderr.WriteString(ut + "\n\n")
}
