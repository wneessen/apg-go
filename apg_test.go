package main

import "testing"

var config Config

// Make sure the flags are initalized
var _ = func() bool {
	testing.Init()
	config = parseFlags()
	return true
}()

// Test getRandNum with max 1000
func TestGetRandNum(t *testing.T) {
	testTable := []struct {
		testName   string
		givenVal   int
		maxRet     int
		minRet     int
		shouldFail bool
	}{
		{"randNum up to 1000", 1000, 1000, 0, false},
		{"randNum should be 1", 1, 1, 0, false},
		{"randNum should fail on 0", 0, 0, 0, true},
		{"randNum should fail on negative", -1, 0, 0, true},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			randNum, err := getRandNum(testCase.givenVal)
			if testCase.shouldFail {
				if err == nil {
					t.Errorf("Random number generation succeeded but was expected to fail. Given: %v, returned: %v",
						testCase.givenVal, randNum)
				}
			} else {
				if err != nil {
					t.Errorf("Random number generation failed: %v", err.Error())
				}
				if randNum > testCase.maxRet {
					t.Errorf("Random number generation returned too big value. Given %v, expected max: %v, got: %v",
						testCase.givenVal, testCase.maxRet, randNum)
				}
				if randNum < testCase.minRet {
					t.Errorf("Random number generation returned too small value. Given %v, expected max: %v, got: %v",
						testCase.givenVal, testCase.minRet, randNum)
				}
			}
		})
	}
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
	lowerCaseBytes := []int{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r',
		's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	lowerCaseHumanBytes := []int{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'j', 'k', 'm', 'n', 'p', 'q', 'r',
		's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	upperCaseBytes := []int{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R',
		'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	upperCaseHumanBytes := []int{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K', 'M', 'N', 'P', 'Q',
		'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	numberBytes := []int{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	numberHumanBytes := []int{'2', '3', '4', '5', '6', '7', '8', '9'}
	specialBytes := []int{'!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/', ':',
		';', '<', '=', '>', '?', '@', '[', '\\', ']', '^', '_', '`', '{', '|', '}', '~'}
	specialHumanBytes := []int{'"', '#', '%', '*', '+', '-', '/', ':', ';', '=', '\\', '_', '|', '~'}
	testTable := []struct {
		testName      string
		allowedBytes  []int
		useLowerCase  bool
		useUpperCase  bool
		useNumber     bool
		useSpecial    bool
		humanReadable bool
	}{
		{"lowercase_only", lowerCaseBytes, true, false, false, false, false},
		{"lowercase_only_human", lowerCaseHumanBytes, true, false, false, false, true},
		{"uppercase_only", upperCaseBytes, false, true, false, false, false},
		{"uppercase_only_human", upperCaseHumanBytes, false, true, false, false, true},
		{"number_only", numberBytes, false, false, true, false, false},
		{"number_only_human", numberHumanBytes, false, false, true, false, true},
		{"special_only", specialBytes, false, false, false, true, false},
		{"special_only_human", specialHumanBytes, false, false, false, true, true},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			config.useLowerCase = testCase.useLowerCase
			config.useUpperCase = testCase.useUpperCase
			config.useNumber = testCase.useNumber
			config.useSpecial = testCase.useSpecial
			config.humanReadable = testCase.humanReadable
			charRange := getCharRange(&config)
			for _, curChar := range charRange {
				searchAllowedBytes := containsByte(testCase.allowedBytes, int(curChar), t)
				if !searchAllowedBytes {
					t.Errorf("Character range returned invalid value: %v", string(curChar))
				}
			}
		})
	}
}

// Test Conversions
func TestConvert(t *testing.T) {
	testTable := []struct {
		testName   string
		givenVal   byte
		expVal     string
		shouldFail bool
	}{
		{"convert_A_to_Alfa", 'A', "Alfa", false},
		{"convert_a_to_alfa", 'a', "alfa", false},
		{"convert_0_to_ZERO", '0', "ZERO", false},
		{"convert_/_to_SLASH", '/', "SLASH", false},
	}

	for _, testCase := range testTable {
		t.Run(testCase.testName, func(t *testing.T) {
			charToString, err := convertCharToName(testCase.givenVal)
			if testCase.shouldFail {
				if err == nil {
					t.Errorf("Character to string conversion succeeded but was expected to fail. Given: %v, returned: %v",
						testCase.givenVal, charToString)
				}
			} else {
				if err != nil {
					t.Errorf("Character to string conversion failed: %v", err.Error())
				}
				if charToString != testCase.expVal {
					t.Errorf("Character to String conversion fail. Given: %q, expected: %q, got: %q",
						testCase.givenVal, testCase.expVal, charToString)
				}
			}
		})
	}

	t.Run("all_chars_must_return_a_conversion_string", func(t *testing.T) {
		config.useUpperCase = true
		config.useLowerCase = true
		config.useNumber = true
		config.useSpecial = true
		config.humanReadable = false
		charRange := getCharRange(&config)
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
	charRange := getCharRange(&config)
	for i := 0; i < b.N; i++ {
		charToConv, _ := getRandChar(&charRange, 1)
		charBytes := []byte(charToConv)
		_, _ = convertCharToName(charBytes[0])
	}
}

// Contains function to search a given slice for values
func containsByte(allowedBytes []int, currentChar int, t *testing.T) bool {
	t.Helper()

	for _, charInt := range allowedBytes {
		if charInt == currentChar {
			return true
		}
	}
	return false
}
