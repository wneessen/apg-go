package apg

import (
	"bytes"
	"strings"
	"testing"
)

func TestGenerator_CoinFlip(t *testing.T) {
	g := New()
	cf := g.CoinFlip()
	if cf < 0 || cf > 1 {
		t.Errorf("CoinFlip failed(), expected 0 or 1, got: %d", cf)
	}
}

func TestGenerator_CoinFlipBool(t *testing.T) {
	g := New()
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

	g := New()
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

	g := New()
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
	g := New()
	l := 32
	tt := []struct {
		name string
		cr   string
		nr   string
		sf   bool
	}{
		{
			"CharRange:AlphaLower", CharRangeAlphaLower,
			CharRangeAlphaUpper + CharRangeNumber + CharRangeSpecial, false,
		},
		{
			"CharRange:AlphaUpper", CharRangeAlphaUpper,
			CharRangeAlphaLower + CharRangeNumber + CharRangeSpecial, false,
		},
		{
			"CharRange:Number", CharRangeNumber,
			CharRangeAlphaLower + CharRangeAlphaUpper + CharRangeSpecial, false,
		},
		{
			"CharRange:Special", CharRangeSpecial,
			CharRangeAlphaLower + CharRangeAlphaUpper + CharRangeNumber, false,
		},
		{
			"CharRange:Invalid", "",
			CharRangeAlphaLower + CharRangeAlphaUpper + CharRangeNumber + CharRangeSpecial, true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			rs, err := g.RandomString(l, tc.cr)
			if err != nil && !tc.sf {
				t.Errorf("RandomString failed: %s", err)
			}
			if len(rs) != l && !tc.sf {
				t.Errorf("RandomString failed. Expected length: %d, got: %d", l, len(rs))
			}
			if strings.ContainsAny(rs, tc.nr) {
				t.Errorf("RandomString failed. Unexpected character found in returned string: %s", rs)
			}
		})
	}
}

func BenchmarkGenerator_CoinFlip(b *testing.B) {
	b.ReportAllocs()
	g := New()
	for i := 0; i < b.N; i++ {
		_ = g.CoinFlip()
	}
}

func BenchmarkGenerator_RandomBytes(b *testing.B) {
	b.ReportAllocs()
	g := New()
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
	g := New()
	cr := CharRangeAlphaUpper + CharRangeAlphaLower + CharRangeNumber + CharRangeSpecial
	for i := 0; i < b.N; i++ {
		_, err := g.RandomString(32, cr)
		if err != nil {
			b.Errorf("RandomString() failed: %s", err)
		}
	}
}
