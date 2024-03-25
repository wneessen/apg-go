// SPDX-FileCopyrightText: 2021-2024 Winni Neessen <wn@neessen.dev>
//
// SPDX-License-Identifier: MIT

package apg

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	c := NewConfig()
	if c == nil {
		t.Errorf("NewConfig() failed, expected config pointer but got nil")
		return
	}
	c = NewConfig(nil)
	if c == nil {
		t.Errorf("NewConfig() failed, expected config pointer but got nil")
		return
	}
	if c.MinLength != DefaultMinLength {
		t.Errorf("NewConfig() failed, expected min length: %d, got: %d", DefaultMinLength,
			c.MinLength)
	}
	if c.MaxLength != DefaultMaxLength {
		t.Errorf("NewConfig() failed, expected max length: %d, got: %d", DefaultMaxLength,
			c.MaxLength)
	}
	if c.NumberPass != DefaultNumberPass {
		t.Errorf("NewConfig() failed, expected number of passwords: %d, got: %d",
			DefaultNumberPass, c.NumberPass)
	}
	if c.Mode != DefaultMode {
		t.Errorf("NewConfig() failed, expected mode mask: %d, got: %d",
			DefaultMode, c.Mode)
	}
}

func TestWithAlgorithm(t *testing.T) {
	tests := []struct {
		name string
		algo Algorithm
		want int
	}{
		{"Pronounceable passwords", AlgoPronounceable, 0},
		{"Random passwords", AlgoRandom, 1},
		{"Coinflip", AlgoCoinFlip, 2},
		{"Binary", AlgoBinary, 3},
		{"Unsupported", AlgoUnsupported, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewConfig(WithAlgorithm(tt.algo))
			if c == nil {
				t.Errorf("NewConfig(WithAlgorithm()) failed, expected config pointer but got nil")
				return
			}
			if c.Algorithm != tt.algo {
				t.Errorf("NewConfig(WithAlgorithm()) failed, expected algo: %d, got: %d",
					tt.algo, c.Algorithm)
			}
			if IntToAlgo(tt.want) != c.Algorithm {
				t.Errorf("IntToAlgo() failed, expected algo: %d, got: %d",
					tt.want, c.Algorithm)
			}
		})
	}
}

func TestWithBinaryHexMode(t *testing.T) {
	c := NewConfig(WithBinaryHexMode())
	if c == nil {
		t.Errorf("NewConfig(WithBinaryHexMode()) failed, expected config pointer but got nil")
		return
	}
	if !c.BinaryHexMode {
		t.Errorf("NewConfig(WithBinaryHexMode()) failed, expected chars: %t, got: %t",
			true, c.BinaryHexMode)
	}
}

func TestWithExcludeChars(t *testing.T) {
	e := "abcdefg"
	c := NewConfig(WithExcludeChars(e))
	if c == nil {
		t.Errorf("NewConfig(WithExcludeChars()) failed, expected config pointer but got nil")
		return
	}
	if c.ExcludeChars != e {
		t.Errorf("NewConfig(WithExcludeChars()) failed, expected chars: %s, got: %s",
			e, c.ExcludeChars)
	}
}

func TestWithFixedLength(t *testing.T) {
	var e int64 = 10
	c := NewConfig(WithFixedLength(e))
	if c == nil {
		t.Errorf("NewConfig(WithFixedLength()) failed, expected config pointer but got nil")
		return
	}
	if c.FixedLength != e {
		t.Errorf("NewConfig(WithFixedLength()) failed, expected fixed length: %d, got: %d",
			e, c.FixedLength)
	}
}

func TestWithMaxLength(t *testing.T) {
	var e int64 = 123
	c := NewConfig(WithMaxLength(e))
	if c == nil {
		t.Errorf("NewConfig(WithMaxLength()) failed, expected config pointer but got nil")
		return
	}
	if c.MaxLength != e {
		t.Errorf("NewConfig(WithMaxLength()) failed, expected max length: %d, got: %d",
			e, c.MaxLength)
	}
}

func TestWithMinLength(t *testing.T) {
	var e int64 = 1
	c := NewConfig(WithMinLength(e))
	if c == nil {
		t.Errorf("NewConfig(WithMinLength()) failed, expected config pointer but got nil")
		return
	}
	if c.MinLength != e {
		t.Errorf("NewConfig(WithMinLength()) failed, expected min length: %d, got: %d",
			e, c.MinLength)
	}
}

func TestWithMinLowercase(t *testing.T) {
	var e int64 = 2
	c := NewConfig(WithMinLowercase(e))
	if c == nil {
		t.Errorf("NewConfig(WithMinLowercase()) failed, expected config pointer but got nil")
		return
	}
	if c.MinLowerCase != e {
		t.Errorf("NewConfig(WithMinLowercase()) failed, expected min amount: %d, got: %d",
			e, c.MinLowerCase)
	}
}

func TestWithMinNumeric(t *testing.T) {
	var e int64 = 3
	c := NewConfig(WithMinNumeric(e))
	if c == nil {
		t.Errorf("NewConfig(WithMinNumeric()) failed, expected config pointer but got nil")
		return
	}
	if c.MinNumeric != e {
		t.Errorf("NewConfig(WithMinNumeric()) failed, expected min amount: %d, got: %d",
			e, c.MinNumeric)
	}
}

func TestWithMinSpecial(t *testing.T) {
	var e int64 = 4
	c := NewConfig(WithMinSpecial(e))
	if c == nil {
		t.Errorf("NewConfig(WithMinSpecial()) failed, expected config pointer but got nil")
		return
	}
	if c.MinSpecial != e {
		t.Errorf("NewConfig(WithMinSpecial()) failed, expected min amount: %d, got: %d",
			e, c.MinSpecial)
	}
}

func TestWithMinUppercase(t *testing.T) {
	var e int64 = 5
	c := NewConfig(WithMinUppercase(e))
	if c == nil {
		t.Errorf("NewConfig(WithMinUppercase()) failed, expected config pointer but got nil")
		return
	}
	if c.MinUpperCase != e {
		t.Errorf("NewConfig(WithMinUppercase()) failed, expected min amount: %d, got: %d",
			e, c.MinUpperCase)
	}
}

func TestWithMobileGrouping(t *testing.T) {
	c := NewConfig(WithMobileGrouping())
	if c == nil {
		t.Errorf("NewConfig(WithMobileGrouping()) failed, expected config pointer but got nil")
		return
	}
	if c.MobileGrouping != true {
		t.Errorf("NewConfig(WithMobileGrouping()) failed, expected: %t, got: %t",
			true, c.MobileGrouping)
	}
}

func TestWithModeMask(t *testing.T) {
	e := DefaultMode
	c := NewConfig(WithModeMask(e))
	if c == nil {
		t.Errorf("NewConfig(WithModeMask()) failed, expected config pointer but got nil")
		return
	}
	if c.Mode != e {
		t.Errorf("NewConfig(WithModeMask()) failed, expected mask: %d, got: %d",
			e, c.Mode)
	}
}

func FuzzWithAlgorithm(f *testing.F) {
	f.Add(0)
	f.Add(1)
	f.Add(2)
	f.Add(3)
	f.Add(-1)
	f.Add(100)
	f.Fuzz(func(t *testing.T, algo int) {
		config := NewConfig(WithAlgorithm(Algorithm(algo)))
		if config.MaxLength < config.MinLength {
			t.Errorf("Invalid algorithm: %d", algo)
		}
	})
}
