package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
	"regexp"
)

// Constants
const DefaultPwLenght int = 20
const VersionString string = "0.2.3"
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
		_, _ = os.Stderr.WriteString("Winni's Advanced Password Generator Clone (apg.go) v" + VersionString + "\n")
		os.Exit(0)
	}
}

// Main function that generated the passwords and returns them
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

	charRange := getCharRange()

	for i := 1; i <= config.numOfPass; i++ {
		pwString, err := getRandChar(&charRange, pwLength)
		if err != nil {
			fmt.Printf("getRandChar returned an error: %q\n", err.Error())
			os.Exit(1)
		}
		fmt.Println(pwString)
	}
}

// Provide the range of available characters based on provided parameters
func getCharRange() string {
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

	return charRange
}

// Generate random characters based on given character range
// and password length
func getRandChar(charRange *string, pwLength int) (string, error) {
	if pwLength <= 0 {
		err := fmt.Errorf("provided pwLength value is <= 0: %v", pwLength)
		return "", err
	}
	availCharsLength := len(*charRange)
	charSlice := []byte(*charRange)
	returnString := make([]byte, pwLength)
	for i := 0; i < pwLength; i++ {
		randNum, err := getRandNum(availCharsLength)
		if err != nil {
			return "", err
		}
		returnString[i] = charSlice[randNum]
	}
	return string(returnString), nil
}

// Generate a random number with given maximum value
func getRandNum(maxNum int) (int, error) {
	if maxNum <= 0 {
		err := fmt.Errorf("provided maxNum is <= 0: %v", maxNum)
		return 0, err
	}
	maxNumBigInt := big.NewInt(int64(maxNum))
	if !maxNumBigInt.IsUint64() {
		err := fmt.Errorf("big.NewInt() generation returned negative value: %v", maxNumBigInt)
		return 0, err
	}
	randNum64, err := rand.Int(rand.Reader, maxNumBigInt)
	if err != nil {
		return 0, err
	}
	randNum := int(randNum64.Int64())
	if randNum < 0 {
		err := fmt.Errorf("generated random number does not fit as int64: %v", randNum64)
		return 0, err
	}
	return randNum, nil
}
