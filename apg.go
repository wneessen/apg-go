package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"regexp"
)

// Constants
const DefaultPwLenght int = 20
const VersionString string = "0.2.0"
const PwLowerCharsHuman string = "abcdefghjkmnpqrstuvwxyz"
const PwUpperCharsHuman string = "ABCDEFGHJKMNPQRSTUVWXYZ"
const PwLowerChars string = "abcdefghijklmnopqrstuvwxyz"
const PwUpperChars string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const PwSpecialCharsHuman string = "\"#/\\$%&+-*"
const PwSpecialChars string = "\"#/!\\$%&+-*.,?=()[]{}:;~^|"
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
	showHelp      bool
	showVersion   bool
}

var config cliOpts

// Read flags
func init() {
	// Bool flags
	flag.BoolVar(&config.useLowerCase, "L", true, "Use lower case characters in passwords")
	flag.BoolVar(&config.useUpperCase, "U", false, "Use upper case characters in passwords")
	flag.BoolVar(&config.useNumber, "N", false, "Use numbers in passwords")
	flag.BoolVar(&config.useSpecial, "S", false, "Use special characters in passwords")
	flag.BoolVar(&config.useComplex, "C", true, "Generate complex passwords (implies -L -U -N -S, disables -H)")
	flag.BoolVar(&config.humanReadable, "H", false, "Generate human-readable passwords")
	flag.BoolVar(&config.showVersion, "v", false, "Show version")

	// Int flags
	flag.IntVar(&config.minPassLen, "m", 10, "Minimum password length")
	flag.IntVar(&config.maxPassLen, "x", DefaultPwLenght, "Maxiumum password length")
	flag.IntVar(&config.numOfPass, "n", 1, "Number of passwords to generate")

	// String flags
	flag.StringVar(&config.excludeChars, "E", "", "Exclude list of characters from generated password")

	flag.Parse()
	if config.showVersion {
		_, _ = os.Stderr.WriteString("Advanced Password Generator v" + VersionString + "\n")
		os.Exit(0)
	}
}

func main() {
	pwLength := config.minPassLen
	if pwLength < config.minPassLen {
		pwLength = config.minPassLen
	}
	if pwLength > config.maxPassLen {
		pwLength = config.maxPassLen
	}
	if config.useComplex {
		config.useUpperCase = true
		config.useLowerCase = true
		config.useSpecial = true
		config.useNumber = true
		config.humanReadable = false
	}

	pwUpperChars := PwUpperChars
	pwLowerChars := PwLowerChars
	pwNumbers := PwNumbers
	pwSpecialChars := PwSpecialChars
	if config.humanReadable {
		pwUpperChars = PwUpperCharsHuman
		pwLowerChars = PwLowerCharsHuman
		pwNumbers = PwNumbersHuman
		pwSpecialChars = PwSpecialCharsHuman
	}

	var charRange string
	if config.useLowerCase {
		charRange = charRange + pwLowerChars
	}
	if config.useUpperCase {
		charRange = charRange + pwUpperChars
	}
	if config.useNumber {
		charRange = charRange + pwNumbers
	}
	if config.useSpecial {
		charRange = charRange + pwSpecialChars
	}
	if config.excludeChars != "" {
		regExp := regexp.MustCompile("[" + config.excludeChars + "]")
		charRange = regExp.ReplaceAllLiteralString(charRange, "")
	}

	for i := 1; i <= config.numOfPass; i++ {
		pwString := getRandChar(&charRange, pwLength)
		fmt.Println(pwString)
	}
}

func getRandChar(charRange *string, pwLength int) string {
	availCharsLength := len(*charRange)
	charSlice := []byte(*charRange)
	returnString := []byte{}
	for i := 0; i < pwLength; i++ {
		randNum := getRandNum(availCharsLength)
		returnString = append(returnString, charSlice[randNum])
	}
	return string(returnString)
}

func getRandNum(maxNum int) int {
	maxNumBigInt := big.NewInt(int64(maxNum))
	randNum64, err := rand.Int(rand.Reader, maxNumBigInt)
	if err != nil {
		log.Fatal("An error occured generating random number: %v", err)
	}
	randNum := int(randNum64.Int64())
	return randNum
}
