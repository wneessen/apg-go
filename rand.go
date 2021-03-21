package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

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
