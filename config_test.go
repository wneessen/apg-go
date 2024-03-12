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
}

func TestWithAlgorithm(t *testing.T) {
	tests := []struct {
		name string
		algo Algorithm
		want int
	}{
		{"Pronouncble passwords", AlgoPronounceable, 0},
		{"Random passwords", AlgoRandom, 1},
		{"Coinflip", AlgoCoinFlip, 2},
		{"Unsupported", AlgoUnsupported, 3},
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

func TestWithNumberPass(t *testing.T) {
	var e int64 = 123
	c := NewConfig(WithNumberPass(e))
	if c == nil {
		t.Errorf("NewConfig(WithNumberPass()) failed, expected config pointer but got nil")
		return
	}
	if c.NumberPass != e {
		t.Errorf("NewConfig(WithNumberPass()) failed, expected number of passwords: %d, got: %d",
			e, c.NumberPass)
	}
}
