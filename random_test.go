// SPDX-FileCopyrightText: 2021-2024 Winni Neessen <wn@neessen.dev>
//
// SPDX-License-Identifier: MIT

package apg

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestGenerator_CoinFlip(t *testing.T) {
	g := New(NewConfig())
	cf := g.CoinFlip()
	if cf < 0 || cf > 1 {
		t.Errorf("CoinFlip failed(), expected 0 or 1, got: %d", cf)
	}
}

func TestGenerator_CoinFlipBool(t *testing.T) {
	g := New(NewConfig())
	gt := false
	for i := 0; i < 500_000; i++ {
		cf := g.CoinFlipBool()
		if cf {
			gt = true
			break
		}
	}
	if !gt {
		t.Error("CoinFlipBool likely not working, expected at least one true value in 500k tries, got none")
	}
}

func TestGenerator_RandNum(t *testing.T) {
	tt := []struct {
		name string
		v    int64
		max  int64
		min  int64
		sf   bool
	}{
		{"RandNum up to 1000", 1000, 1000, 0, false},
		{"RandNum should be 1", 1, 1, 0, false},
		{"RandNum should fail on 1", 0, 0, 0, true},
		{"RandNum should fail on negative", -1, 0, 0, true},
	}

	g := New(NewConfig())
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			rn, err := g.RandNum(tc.v)
			if err == nil && tc.sf {
				t.Errorf("random number generation was supposed to fail, but didn't, got: %d", rn)
			}
			if err != nil && !tc.sf {
				t.Errorf("random number generation failed: %s", err)
			}
			if rn > tc.max {
				t.Errorf("random number generation returned too big number expected below: %d, got: %d",
					tc.max, rn)
			}
			if rn < tc.min {
				t.Errorf("random number generation returned too small number, expected min: %d, got: %d",
					tc.min, rn)
			}
		})
	}
}

func TestGenerator_RandomBytes(t *testing.T) {
	tt := []struct {
		name string
		l    int64
		sf   bool
	}{
		{"1 bytes of randomness", 1, false},
		{"100 bytes of randomness", 100, false},
		{"1024 bytes of randomness", 1024, false},
		{"4096 bytes of randomness", 4096, false},
		{"-1 bytes of randomness", -1, true},
	}

	g := New(NewConfig())
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			rb, err := g.RandomBytes(tc.l)
			if err != nil && !tc.sf {
				t.Errorf("random byte generation failed: %s", err)
				return
			}
			if tc.sf {
				return
			}
			bl := len(rb)
			if int64(bl) != tc.l {
				t.Errorf("lenght of provided bytes does not match requested length: got: %d, expected: %d",
					bl, tc.l)
			}
			if bytes.Equal(rb, make([]byte, tc.l)) {
				t.Errorf("random byte generation failed. returned slice is empty")
			}
		})
	}
}

func TestGenerator_RandomString(t *testing.T) {
	g := New(NewConfig())
	var l int64 = 32 * 1024
	tt := []struct {
		name string
		cr   string
		nr   string
		sf   bool
	}{
		{
			"CharRange:AlphaLower", CharRangeAlphaLower,
			CharRangeAlphaUpper + CharRangeNumeric + CharRangeSpecial, false,
		},
		{
			"CharRange:AlphaUpper", CharRangeAlphaUpper,
			CharRangeAlphaLower + CharRangeNumeric + CharRangeSpecial, false,
		},
		{
			"CharRange:Number", CharRangeNumeric,
			CharRangeAlphaLower + CharRangeAlphaUpper + CharRangeSpecial, false,
		},
		{
			"CharRange:Special", CharRangeSpecial,
			CharRangeAlphaLower + CharRangeAlphaUpper + CharRangeNumeric, false,
		},
		{
			"CharRange:Invalid", "",
			CharRangeAlphaLower + CharRangeAlphaUpper + CharRangeNumeric + CharRangeSpecial, true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			rs, err := g.RandomStringFromCharRange(l, tc.cr)
			if err != nil && !tc.sf {
				t.Errorf("RandomStringFromCharRange failed: %s", err)
			}
			if int64(len(rs)) != l && !tc.sf {
				t.Errorf("RandomStringFromCharRange failed. Expected length: %d, got: %d", l, len(rs))
			}
			if strings.ContainsAny(rs, tc.nr) {
				t.Errorf("RandomStringFromCharRange failed. Unexpected character found in returned string: %s", rs)
			}
		})
	}
}

func TestGetCharRangeFromConfig(t *testing.T) {
	config := NewConfig()
	generator := New(config)
	testCases := []struct {
		Name          string
		ConfigMode    ModeMask
		ExpectedRange string
	}{
		{
			Name:          "LowerCaseHumanReadable",
			ConfigMode:    ModeLowerCase | ModeHumanReadable,
			ExpectedRange: CharRangeAlphaLowerHuman,
		},
		{
			Name:          "LowerCaseNonHumanReadable",
			ConfigMode:    ModeLowerCase,
			ExpectedRange: CharRangeAlphaLower,
		},
		{
			Name:          "NumericHumanReadable",
			ConfigMode:    ModeNumeric | ModeHumanReadable,
			ExpectedRange: CharRangeNumericHuman,
		},
		{
			Name:          "NumericNonHumanReadable",
			ConfigMode:    ModeNumeric,
			ExpectedRange: CharRangeNumeric,
		},
		{
			Name:          "SpecialHumanReadable",
			ConfigMode:    ModeSpecial | ModeHumanReadable,
			ExpectedRange: CharRangeSpecialHuman,
		},
		{
			Name:          "SpecialNonHumanReadable",
			ConfigMode:    ModeSpecial,
			ExpectedRange: CharRangeSpecial,
		},
		{
			Name:          "UpperCaseHumanReadable",
			ConfigMode:    ModeUpperCase | ModeHumanReadable,
			ExpectedRange: CharRangeAlphaUpperHuman,
		},
		{
			Name:          "UpperCaseNonHumanReadable",
			ConfigMode:    ModeUpperCase,
			ExpectedRange: CharRangeAlphaUpper,
		},
		{
			Name:          "MultipleModes",
			ConfigMode:    ModeLowerCase | ModeNumeric | ModeUpperCase | ModeHumanReadable,
			ExpectedRange: CharRangeAlphaLowerHuman + CharRangeNumericHuman + CharRangeAlphaUpperHuman,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			generator.config.Mode = tc.ConfigMode
			actualRange := generator.GetCharRangeFromConfig()
			if actualRange != tc.ExpectedRange {
				t.Errorf("Expected range %s, got %s", tc.ExpectedRange, actualRange)
			}
		})
	}
}

func TestGetPasswordLength(t *testing.T) {
	config := NewConfig()
	generator := New(config)
	testCases := []struct {
		Name              string
		ConfigFixedLength int64
		ConfigMinLength   int64
		ConfigMaxLength   int64
		ExpectedLength    int64
		ExpectedError     error
	}{
		{
			Name:              "FixedLength",
			ConfigFixedLength: 10,
			ConfigMinLength:   5,
			ConfigMaxLength:   15,
			ExpectedLength:    10,
			ExpectedError:     nil,
		},
		{
			Name:              "MinLengthEqualToMaxLength",
			ConfigFixedLength: 0,
			ConfigMinLength:   8,
			ConfigMaxLength:   8,
			ExpectedLength:    8,
			ExpectedError:     nil,
		},
		{
			Name:              "MinLengthGreaterThanMaxLength",
			ConfigFixedLength: 0,
			ConfigMinLength:   12,
			ConfigMaxLength:   5,
			ExpectedLength:    12,
			ExpectedError:     nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			generator.config.FixedLength = tc.ConfigFixedLength
			generator.config.MinLength = tc.ConfigMinLength
			generator.config.MaxLength = tc.ConfigMaxLength
			length, err := generator.GetPasswordLength()

			if err != nil && !errors.Is(err, tc.ExpectedError) {
				t.Errorf("Unexpected error: %v", err)
			} else if err == nil && tc.ExpectedError != nil {
				t.Errorf("Expected error %v, got nil", tc.ExpectedError)
			} else if err == nil && length != tc.ExpectedLength {
				t.Errorf("Expected length %d, got %d", tc.ExpectedLength, length)
			}
		})
	}
}

// TestGenerateCoinFlip tries to test the coinflip. Randomness is hard to
// test, since it's supposed to be not prodictable. We think that in 100k
// tries at least one of each two results should be returned, even though
// it is possible that not. Therefore we only throw a warning
func TestGenerateCoinFlip(t *testing.T) {
	config := NewConfig()
	generator := New(config)
	foundTails := false
	foundHeads := false
	for range 100_000 {
		res, err := generator.generateCoinFlip()
		if err != nil {
			t.Errorf("generateCoinFlip() failed: %s", err)
			return
		}
		switch res {
		case "Tails":
			foundTails = true
		case "Heads":
			foundHeads = true
		}
	}
	if !foundTails && !foundHeads {
		t.Logf("WARNING: generateCoinFlip() was supposed to find heads and tails "+
			"in 100_000 tries but didn't. Heads: %t, Tails: %t", foundHeads, foundTails)
	}
}

func TestGeneratePronounceable(t *testing.T) {
	config := NewConfig()
	generator := New(config)
	foundSylables := 0
	for range 100 {
		res, err := generator.generatePronounceable()
		if err != nil {
			t.Errorf("generatePronounceable() failed: %s", err)
			return
		}
		for _, syl := range KoremutakeSyllables {
			if strings.Contains(res, syl) {
				foundSylables++
			}
		}
	}
	if foundSylables < 100 {
		t.Errorf("generatePronounceable() failed, expected at least 1 sylable, got none")
	}
}

func TestCheckMinimumRequirements(t *testing.T) {
	config := NewConfig()
	generator := New(config)
	testCases := []struct {
		Name               string
		Password           string
		ConfigMinLowerCase int64
		ConfigMinNumeric   int64
		ConfigMinSpecial   int64
		ConfigMinUpperCase int64
		ExpectedResult     bool
	}{
		{
			Name:               "Meets all requirements",
			Password:           "Th1sIsA$trongP@ssword",
			ConfigMinLowerCase: 2,
			ConfigMinNumeric:   1,
			ConfigMinSpecial:   1,
			ConfigMinUpperCase: 1,
			ExpectedResult:     true,
		},
		{
			Name:               "Missing lowercase",
			Password:           "THISIS@STRONGPASSWORD",
			ConfigMinLowerCase: 2,
			ConfigMinNumeric:   1,
			ConfigMinSpecial:   1,
			ConfigMinUpperCase: 1,
			ExpectedResult:     false,
		},
		{
			Name:               "Missing numeric",
			Password:           "ThisIsA$trongPassword",
			ConfigMinLowerCase: 2,
			ConfigMinNumeric:   1,
			ConfigMinSpecial:   1,
			ConfigMinUpperCase: 1,
			ExpectedResult:     false,
		},
		{
			Name:               "Missing special",
			Password:           "ThisIsALowercaseNumericPassword",
			ConfigMinLowerCase: 2,
			ConfigMinNumeric:   1,
			ConfigMinSpecial:   1,
			ConfigMinUpperCase: 1,
			ExpectedResult:     false,
		},
		{
			Name:               "Missing uppercase",
			Password:           "thisisanumericspecialpassword",
			ConfigMinLowerCase: 2,
			ConfigMinNumeric:   1,
			ConfigMinSpecial:   1,
			ConfigMinUpperCase: 1,
			ExpectedResult:     false,
		},
		{
			Name:               "Bare minimum",
			Password:           "a1!",
			ConfigMinLowerCase: 1,
			ConfigMinNumeric:   1,
			ConfigMinSpecial:   1,
			ConfigMinUpperCase: 0,
			ExpectedResult:     true,
		},
		{
			Name:               "Empty password",
			Password:           "",
			ConfigMinLowerCase: 1,
			ConfigMinNumeric:   1,
			ConfigMinSpecial:   1,
			ConfigMinUpperCase: 1,
			ExpectedResult:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			generator.config.MinLowerCase = tc.ConfigMinLowerCase
			generator.config.MinNumeric = tc.ConfigMinNumeric
			generator.config.MinSpecial = tc.ConfigMinSpecial
			generator.config.MinUpperCase = tc.ConfigMinUpperCase
			result := generator.checkMinimumRequirements(tc.Password)
			if result != tc.ExpectedResult {
				t.Errorf("Expected result %v, got %v", tc.ExpectedResult, result)
			}
		})
	}
}

func TestGenerateRandom(t *testing.T) {
	config := NewConfig(WithAlgorithm(AlgoRandom), WithMinLength(1),
		WithMaxLength(1))
	config.MinNumeric = 1
	generator := New(config)
	pw, err := generator.generateRandom()
	if err != nil {
		t.Errorf("generateRandom() failed: %s", err)
	}
	if len(pw) > 1 {
		t.Errorf("expected password with length 1 but got: %d", len(pw))
	}
	n, err := strconv.Atoi(pw)
	if err != nil {
		t.Errorf("expected password to be a number but got an error: %s", err)
	}
	if n < 0 || n > 9 {
		t.Errorf("expected password to be a number between 0 and 9, got: %d", n)
	}
}

func TestGenerate(t *testing.T) {
	tests := []struct {
		name        string
		algorithm   Algorithm
		expectedErr error
	}{
		{
			name:      "Pronounceable",
			algorithm: AlgoPronounceable,
		},
		{
			name:      "CoinFlip",
			algorithm: AlgoCoinFlip,
		},
		{
			name:      "Random",
			algorithm: AlgoRandom,
		},
		{
			name:        "Unsupported",
			algorithm:   AlgoUnsupported,
			expectedErr: fmt.Errorf("unsupported algorithm"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewConfig(WithAlgorithm(tt.algorithm))
			g := New(config)
			_, err := g.Generate()
			if tt.expectedErr != nil {
				if err == nil || err.Error() != tt.expectedErr.Error() {
					t.Errorf("Expected error: %s, got: %s", tt.expectedErr, err)
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %s", err)
				return
			}
		})
	}
}

func BenchmarkGenerator_CoinFlip(b *testing.B) {
	b.ReportAllocs()
	g := New(NewConfig())
	for i := 0; i < b.N; i++ {
		_ = g.CoinFlip()
	}
}

func BenchmarkGenerator_RandomBytes(b *testing.B) {
	b.ReportAllocs()
	g := New(NewConfig())
	var l int64 = 1024
	for i := 0; i < b.N; i++ {
		_, err := g.RandomBytes(l)
		if err != nil {
			b.Errorf("failed to generate random bytes: %s", err)
			return
		}
	}
}

func BenchmarkGenerator_RandomString(b *testing.B) {
	b.ReportAllocs()
	g := New(NewConfig())
	cr := CharRangeAlphaUpper + CharRangeAlphaLower + CharRangeNumeric + CharRangeSpecial
	for i := 0; i < b.N; i++ {
		_, err := g.RandomStringFromCharRange(32, cr)
		if err != nil {
			b.Errorf("RandomStringFromCharRange() failed: %s", err)
		}
	}
}
