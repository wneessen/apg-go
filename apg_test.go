package main

import (
	"math/big"
	"testing"
)

// Make sure the flags are initalized
var _ = func() bool {
	testing.Init()
	return true
}()

// Test getRandNum with max 1000
func TestGetRandNum(t *testing.T) {
	t.Run("maxNum_is_1000", func(t *testing.T) {
		randNum, err := getRandNum(1000)
		if err != nil {
			t.Fatalf("Random number generation failed: %v", err.Error())
		}
		if randNum > 1000 {
			t.Fatalf("Generated random number between 0 and 1000 is too big: %v", randNum)
		}
		if randNum < 0 {
			t.Fatalf("Generated random number between 0 and 1000 is too small: %v", randNum)
		}
	})
	t.Run("maxNum_is_1", func(t *testing.T) {
		randNum, err := getRandNum(1)
		if err != nil {
			t.Fatalf("Random number generation failed: %v", err.Error())
		}
		if randNum > 1 {
			t.Fatalf("Generated random number between 0 and 1000 is too big: %v", randNum)
		}
		if randNum < 0 {
			t.Fatalf("Generated random number between 0 and 1000 is too small: %v", randNum)
		}
	})
	t.Run("maxNum_is_0", func(t *testing.T) {
		randNum, err := getRandNum(0)
		if err == nil {
			t.Fatalf("Random number expected to fail, but provided a value instead: %v", randNum)
		}
	})
}

// Test getRandChar
func TestGetRandChar(t *testing.T) {
	t.Run("return_value_is_A_B_or_C", func(t *testing.T) {
		charRange := "ABC"
		randChar, err := getRandChar(&charRange, 1)
		if err != nil {
			t.Fatalf("Random character generation failed => %v", err.Error())
		}
		if randChar != "A" && randChar != "B" && randChar != "C" {
			t.Fatalf("Random character generation failed. Expected A, B or C but got: %v", randChar)
		}
	})

	t.Run("return_value_has_specific_length", func(t *testing.T) {
		charRange := "ABC"
		randChar, err := getRandChar(&charRange, 1000)
		if err != nil {
			t.Fatalf("Random character generation failed => %v", err.Error())
		}
		if len(randChar) != 1000 {
			t.Fatalf("Generated random characters with 1000 chars returned wrong amount of chars: %v",
				len(randChar))
		}
	})

	t.Run("fail", func(t *testing.T) {
		charRange := "ABC"
		randChar, err := getRandChar(&charRange, -2000)
		if err == nil {
			t.Fatalf("Generated random characters expected to fail, but returned a value => %v",
				randChar)
		}
	})
}

// Test getCharRange() with different config settings
func TestGetCharRange(t *testing.T) {

	t.Run("lower_case_only", func(t *testing.T) {
		// Lower case only
		allowedBytes := []int{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r',
			's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
		config.useLowerCase = true
		config.useUpperCase = false
		config.useNumber = false
		config.useSpecial = false
		config.humanReadable = false
		charRange := getCharRange()
		for _, curChar := range charRange {
			searchAllowedBytes := containsByte(allowedBytes, int(curChar))
			if !searchAllowedBytes {
				t.Fatalf("Character range for lower-case only returned invalid value: %v",
					string(curChar))
			}
		}
	})

	// Lower case only (human readable)
	t.Run("lower_case_only_human_readable", func(t *testing.T) {
		allowedBytes := []int{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'j', 'k', 'm', 'n', 'p', 'q', 'r',
			's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
		config.useLowerCase = true
		config.useUpperCase = false
		config.useNumber = false
		config.useSpecial = false
		config.humanReadable = true
		charRange := getCharRange()
		for _, curChar := range charRange {
			searchAllowedBytes := containsByte(allowedBytes, int(curChar))
			if !searchAllowedBytes {
				t.Fatalf("Character range for lower-case only (human readable) returned invalid value: %v",
					string(curChar))
			}
		}
	})

	// Upper case only
	t.Run("upper_case_only", func(t *testing.T) {
		allowedBytes := []int{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R',
			'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
		config.useLowerCase = false
		config.useUpperCase = true
		config.useNumber = false
		config.useSpecial = false
		config.humanReadable = false
		charRange := getCharRange()
		for _, curChar := range charRange {
			searchAllowedBytes := containsByte(allowedBytes, int(curChar))
			if !searchAllowedBytes {
				t.Fatalf("Character range for upper-case only returned invalid value: %v",
					string(curChar))
			}
		}
	})

	// Upper case only (human readable)
	t.Run("upper_case_only_human_readable", func(t *testing.T) {
		allowedBytes := []int{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K', 'M', 'N', 'P', 'Q',
			'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
		config.useLowerCase = false
		config.useUpperCase = true
		config.useNumber = false
		config.useSpecial = false
		config.humanReadable = true
		charRange := getCharRange()
		for _, curChar := range charRange {
			searchAllowedBytes := containsByte(allowedBytes, int(curChar))
			if !searchAllowedBytes {
				t.Fatalf("Character range for upper-case only (human readable) returned invalid value: %v",
					string(curChar))
			}
		}
	})

	// Numbers only
	t.Run("numbers_only", func(t *testing.T) {
		allowedBytes := []int{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
		config.useLowerCase = false
		config.useUpperCase = false
		config.useNumber = true
		config.useSpecial = false
		config.humanReadable = false
		charRange := getCharRange()
		for _, curChar := range charRange {
			searchAllowedBytes := containsByte(allowedBytes, int(curChar))
			if !searchAllowedBytes {
				t.Fatalf("Character range for numbers only returned invalid value: %v",
					string(curChar))
			}
		}
	})

	// Numbers only (human readable)
	t.Run("numbers_only_human_readable", func(t *testing.T) {
		allowedBytes := []int{'2', '3', '4', '5', '6', '7', '8', '9'}
		config.useLowerCase = false
		config.useUpperCase = false
		config.useNumber = true
		config.useSpecial = false
		config.humanReadable = true
		charRange := getCharRange()
		for _, curChar := range charRange {
			searchAllowedBytes := containsByte(allowedBytes, int(curChar))
			if !searchAllowedBytes {
				t.Fatalf("Character range for numbers (human readable) only returned invalid value: %v",
					string(curChar))
			}
		}
	})

	// Special characters only
	t.Run("special_chars_only", func(t *testing.T) {
		allowedBytes := []int{'!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/', ':',
			';', '<', '=', '>', '?', '@', '[', '\\', ']', '^', '_', '`', '{', '|', '}', '~'}
		config.useLowerCase = false
		config.useUpperCase = false
		config.useNumber = false
		config.useSpecial = true
		config.humanReadable = false
		charRange := getCharRange()
		for _, curChar := range charRange {
			searchAllowedBytes := containsByte(allowedBytes, int(curChar))
			if !searchAllowedBytes {
				t.Fatalf("Character range for special characters only returned invalid value: %v",
					string(curChar))
			}
		}
	})

	// Special characters only (human readable)
	t.Run("special_chars_only_human_readable", func(t *testing.T) {
		allowedBytes := []int{'"', '#', '%', '*', '+', '-', '/', ':', ';', '=', '\\', '_', '|', '~'}
		config.useLowerCase = false
		config.useUpperCase = false
		config.useNumber = false
		config.useSpecial = true
		config.humanReadable = true
		charRange := getCharRange()
		for _, curChar := range charRange {
			searchAllowedBytes := containsByte(allowedBytes, int(curChar))
			if !searchAllowedBytes {
				t.Fatalf("Character range for special characters only returned invalid value: %v",
					string(curChar))
			}
		}
	})
}

// Test Conversions
func TestConvert(t *testing.T) {
	t.Run("convert_A_to_Alfa", func(t *testing.T) {
		charToString, err := convertCharToName('A')
		if err != nil {
			t.Fatalf("Character to string conversion failed: %v", err.Error())
		}
		if charToString != "Alfa" {
			t.Fatalf("Converting 'A' to string did not return the correct value of 'Alfa': %q", charToString)
		}
	})
	t.Run("convert_a_to_alfa", func(t *testing.T) {
		charToString, err := convertCharToName('a')
		if err != nil {
			t.Fatalf("Character to string conversion failed: %v", err.Error())
		}
		if charToString != "alfa" {
			t.Fatalf("Converting 'a' to string did not return the correct value of 'alfa': %q", charToString)
		}
	})
	t.Run("convert_0_to_ZERO", func(t *testing.T) {
		charToString, err := convertCharToName('0')
		if err != nil {
			t.Fatalf("Character to string conversion failed: %v", err.Error())
		}
		if charToString != "ZERO" {
			t.Fatalf("Converting '0' to string did not return the correct value of 'ZERO': %q", charToString)
		}
	})
	t.Run("convert_/_to_SLASH", func(t *testing.T) {
		charToString, err := convertCharToName('/')
		if err != nil {
			t.Fatalf("Character to string conversion failed: %v", err.Error())
		}
		if charToString != "SLASH" {
			t.Fatalf("Converting '/' to string did not return the correct value of 'SLASH': %q", charToString)
		}
	})
	t.Run("all_chars_convert_to_string", func(t *testing.T) {
		config.useUpperCase = true
		config.useLowerCase = true
		config.useNumber = true
		config.useSpecial = true
		config.humanReadable = false
		charRange := getCharRange()
		for _, curChar := range charRange {
			_, err := convertCharToName(byte(curChar))
			if err != nil {
				t.Fatalf("Character to string conversion failed: %v", err.Error())
			}
		}
	})
	t.Run("spell_Ab!_to_strings", func(t *testing.T) {
		pwString := "Ab!"
		spelledString, err := spellPasswordString(pwString)
		if err != nil {
			t.Fatalf("password spelling failed: %v", err.Error())
		}
		if spelledString != "Alfa/bravo/EXCLAMATION_POINT" {
			t.Fatalf(
				"Spelling pwString 'Ab!' is expected to provide 'Alfa/bravo/EXCLAMATION_POINT', but returned: %q",
				spelledString)
		}
	})
}

// Forced failures
func TestForceFailures(t *testing.T) {
	t.Run("too_big_big.NewInt_value", func(t *testing.T) {
		maxNum := 9223372036854775807
		maxNumBigInt := big.NewInt(int64(maxNum) + 1)
		if maxNumBigInt.IsUint64() {
			t.Fatalf("Calling big.NewInt() with too large number expected to fail: %v", maxNumBigInt)
		}
	})

	t.Run("negative value for big.NewInt()", func(t *testing.T) {
		randNum, err := getRandNum(-20000)
		if err == nil {
			t.Fatalf("Calling getRandNum() with negative value is expected to fail, but returned value: %v",
				randNum)
		}
	})
}

// Benchmark: Random number generation
func BenchmarkGetRandNum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = getRandNum(100000)
	}
}

// Benchmark: Random char generation
func BenchmarkGetRandChar(b *testing.B) {
	charRange := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890\"#/!\\$%&+-*.,?=()[]{}:;~^|"
	for i := 0; i < b.N; i++ {
		_, _ = getRandChar(&charRange, 20)
	}
}

// Benchmark: Random char generation
func BenchmarkConvertChar(b *testing.B) {
	config.useUpperCase = true
	config.useLowerCase = true
	config.useNumber = true
	config.useSpecial = true
	config.humanReadable = false
	charRange := getCharRange()
	for i := 0; i < b.N; i++ {
		charToConv, _ := getRandChar(&charRange, 1)
		charBytes := []byte(charToConv)
		_, _ = convertCharToName(charBytes[0])
	}
}

// Contains function to search a given slice for values
func containsByte(allowedBytes []int, currentChar int) bool {
	for _, charInt := range allowedBytes {
		if charInt == currentChar {
			return true
		}
	}
	return false
}
